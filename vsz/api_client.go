package vsz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

const (
	DefaultWSGPathPrefix     = "/wsg/api/public"
	DefaultSwitchMPathPrefix = "/switchm/api"

	logDebugAPIRequestPrepFormat     = "Preparing api request #%d \"%s %s\""
	logDebugAPIRequestNoBodyFormat   = "%s without body"
	logDebugAPIRequestWithBodyFormat = "%s with body: %s"
)

var (
	pkgValidator *validator.Validate
)

func init() {
	pkgValidator = validator.New()
}

// PackageValidator returns the global validator instance used by this client
func PackageValidator() *validator.Validate {
	return pkgValidator
}

type APIConfig struct {
	// Address [required]
	//
	// Full address of VSZ, including scheme and port
	Address string `json:"address"`

	// WSGPathPrefix [optional]
	//
	// Path prefix to access Wireless Security Gateway API endpoints
	WSGPathPrefix string `json:"wsgPathPrefix"`

	// SwitchMPathPrefix [optional]
	//
	// Path prefix to access Switch Management API endpoints
	SwitchMPathPrefix string `json:"switchMPathPrefix"`

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool `json:"debug"`

	// Authenticator [required]
	//
	// Authenticator to use to handle request auth session
	Authenticator Authenticator `json:"-"`

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger `json:"-"`

	// HTTPClient [optional]
	HTTPClient *http.Client `json:"-"`
}

type APIClient struct {
	log   *log.Logger
	debug bool

	addr        string
	wsgPath     string
	switchMPath string
	auth        Authenticator

	client *http.Client
}

func NewAPIClient(conf *APIConfig) (*APIClient, error) {
	if conf.Address == "" {
		return nil, errors.New("address must be defined")
	}
	if conf.Authenticator == nil {
		return nil, errors.New("authenticator must be defined")
	}

	if _, err := url.Parse(conf.Address); err != nil {
		return nil, fmt.Errorf("invalid address specified: %w", err)
	}

	c := new(APIClient)
	c.addr = conf.Address
	c.auth = conf.Authenticator
	c.debug = conf.Debug

	if conf.Logger != nil {
		c.log = conf.Logger
	} else {
		c.log = log.New(ioutil.Discard, "", 0)
	}

	if conf.HTTPClient != nil {
		c.client = conf.HTTPClient
	} else {
		// pooled transport config shamelessly borrowed from https://github.com/hashicorp/go-cleanhttp/blob/v0.5.1/cleanhttp.go
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			},
		}
	}

	if conf.WSGPathPrefix != "" {
		c.wsgPath = conf.WSGPathPrefix
	} else {
		c.wsgPath = DefaultWSGPathPrefix
	}
	if conf.SwitchMPathPrefix != "" {
		c.switchMPath = conf.SwitchMPathPrefix
	} else {
		c.switchMPath = DefaultSwitchMPathPrefix
	}

	return c, nil
}

// Address returns the address of the VSZ node currently being connected to
func (c *APIClient) Address() string {
	return c.addr
}

func (c *APIClient) Debug() bool {
	return c.debug
}

func (c *APIClient) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *APIClient) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}

func (c *APIClient) Do(ctx context.Context, request *APIRequest) (*http.Response, error) {
	_, httpResponse, err := c.do(ctx, request)
	return httpResponse, err
}

func (c *APIClient) do(ctx context.Context, request *APIRequest) (AuthCAS, *http.Response, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		cas          AuthCAS
		err          error
	)
	if request.authenticated {
		if cas, err = c.auth.Decorate(ctx, request); err != nil {
			if cas, err = c.auth.Refresh(ctx, c, cas); err != nil {
				return cas, nil, err
			} else if cas, err = c.auth.Decorate(ctx, request); err != nil {
				return cas, nil, err
			}
		}
	}
	if c.debug {
		// if debug mode is enabled, prepare a big'ol log statement.
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, request.ID(), request.Method(), request.CompiledURI())

		body := request.Body()

		if len(body) == 0 {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg, string(body))
		}

		c.log.Print(logMsg)
	}
	if httpRequest, err = request.toHTTP(ctx, c.addr); err != nil {
		return cas, nil, err
	}
	httpResponse, err = c.client.Do(httpRequest)
	return cas, httpResponse, err
}

type APIResponseError struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	ErrorType string `json:"errorType"`

	Err error `json:"err"`
}

func (e *APIResponseError) Error() string {
	return fmt.Sprintf("success=%t; errorCode=%d; errorType=%s; message=%s; err=%v", e.Success, e.ErrorCode, e.ErrorType, e.Message, e.Err)
}

func (e *APIResponseError) Unwrap() error {
	return e.Err
}

type APIResponseMeta struct {
	RequestMethod string `json:"requestMethod"`
	RequestURI    string `json:"requestURI"`

	SuccessCode int `json:"successCode"`

	ResponseCode   int    `json:"responseCode"`
	ResponseStatus string `json:"responseStatus"`
}

func newAPIResponseMeta(req *APIRequest, successCode, responseCode int, responseStatus string) *APIResponseMeta {
	rm := new(APIResponseMeta)
	rm.RequestMethod = req.Method()
	rm.RequestURI = req.CompiledURI()
	rm.SuccessCode = successCode
	rm.ResponseCode = responseCode
	rm.ResponseStatus = responseStatus
	return rm
}

func (rm *APIResponseMeta) String() string {
	var msg string
	if rm.SuccessCode == rm.ResponseCode {
		msg = "Successful"
	} else {
		msg = "Failed"
	}
	return fmt.Sprintf("%s response from request %s %s", msg, rm.RequestMethod, rm.RequestURI)
}

func wrapErr(err, prev error) error {
	if prev != nil {
		// i want to retain the most recent error context as unwrap-able.
		return fmt.Errorf("%v; %w", prev, err)
	}
	return err
}

// handleResponse will attempt to:
// - wrap errors nicely
// - if a response model is provided, unmarshal response into it
// -- in this case, the first value of the response tuple will be nil
// - if no response model is provided, simply read out all body bytes and return them
// - construct response meta type
func handleResponse(req *APIRequest, successCode int, httpResp *http.Response, modelPtr interface{}, sourceErr error) (*APIResponseMeta, *APIResponseError) {
	// todo: do better.
	var (
		responseCode   int
		responseStatus string
		apiErr         *APIResponseError
		finalErr       error
	)

	if sourceErr != nil {
		finalErr = sourceErr
	}

	// if we saw any kind of upstream error, immediately wrap in APIError type
	// todo: flesh out with more error types

	// if a response was received store the code, status, and response bytes for later use constructing the meta type
	if httpResp != nil {

		responseCode = httpResp.StatusCode
		responseStatus = http.StatusText(httpResp.StatusCode)

		// if we get here, we have some kind of response.  queue up body close...
		defer func() { _ = httpResp.Body.Close() }()

		if responseCode == successCode {
			// if the response code matches the expected "success" code...
			if modelPtr != nil {
				// ... and this query has a modeled response, attempt to unmarshal into that type
				if err := json.NewDecoder(httpResp.Body).Decode(modelPtr); err != nil {
					finalErr = wrapErr(fmt.Errorf("error unmarshalling resoonse body into %T: %w", modelPtr, err), finalErr)
				}
			} else {
				// todo: record bytes from here?
				_, _ = io.Copy(ioutil.Discard, httpResp.Body)
			}
		} else {
			apiErr = new(APIResponseError)
			if err := json.NewDecoder(httpResp.Body).Decode(apiErr); err != nil {
				finalErr = wrapErr(fmt.Errorf("error unmarshalling error body into %T: %w", apiErr, err), finalErr)
			}
			finalErr = wrapErr(fmt.Errorf("expected response code %d, saw %d (%s)", successCode, responseCode, responseStatus), finalErr)
			// todo: record bytes from here?
			_, _ = io.Copy(ioutil.Discard, httpResp.Body)
		}
	}

	if finalErr != nil {
		if apiErr == nil {
			apiErr = new(APIResponseError)
		}
		apiErr.Err = finalErr
	}

	// build response.
	return newAPIResponseMeta(req, successCode, responseCode, responseStatus), apiErr
}
