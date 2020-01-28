package vsz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	DefaultAPIPort           = 8443
	DefaultWSGPathPrefix     = "wsg/api"
	DefaultSwitchMPathPrefix = "switchm/api"
)

type APIConfig struct {
	// Address [required]
	//
	// Address or hostname of your VSZ instance
	Address string `json:"address"`

	// WSGPathPrefix [optional]
	//
	// Path prefix to access Wireless Security Gateway API endpoints
	WSGPathPrefix string `json:"wsgPathPrefix"`

	// SwitchMPathPrefix [optional]
	//
	// Path prefix to access Switch Management API endpoints
	SwitchMPathPrefix string `json:"switchMPathPrefix"`

	// APIPort [optional]
	//
	// APIPort to use when executing API calls.
	APIPort int `json:"port"`

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
	port        int
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

	if conf.APIPort != 0 {
		c.port = conf.APIPort
	} else {
		c.port = DefaultAPIPort
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
	if httpRequest, err = request.toHTTP(ctx, c.addr, c.port, c.debug); err != nil {
		return cas, nil, err
	}
	httpResponse, err = c.client.Do(httpRequest)
	return cas, httpResponse, err
}

type APIError struct {
	SourceError error `json:"sourceError"`

	RequestMethod string `json:"requestMethod"`
	RequestURI    string `json:"requestURI"`

	SuccessCode int `json:"successCode"`

	ResponseCode   int    `json:"responseCode"`
	ResponseStatus string `json:"responseStatus"`
	ResponseBody   []byte `json:"responseBody"`
}

func newAPIError(req *APIRequest, successCode int, httpResp *http.Response, respBytes []byte, err error) error {
	ae := new(APIError)
	ae.SourceError = err
	ae.RequestMethod = req.Method()
	ae.RequestURI = req.CompiledURI()
	ae.SuccessCode = successCode
	if httpResp != nil {
		ae.ResponseCode = httpResp.StatusCode
		ae.ResponseStatus = httpResp.Status
		ae.ResponseBody = respBytes
	}
	return ae
}

func (e *APIError) Error() string {
	// if there was a client error
	if e.SourceError != nil {
		return e.SourceError.Error()
	}
	// if there was an api response error
	if e.ResponseCode != 0 && e.ResponseCode != e.SuccessCode {
		return fmt.Sprintf("expected code %d from %s %s, saw %d %s", e.SuccessCode, e.RequestMethod, e.RequestURI, e.ResponseCode, e.ResponseStatus)
	}
	// if no error
	return ""
}

func (e *APIError) Unwrap() error {
	return e.SourceError
}

func handleResponse(req *APIRequest, successCode int, httpResp *http.Response, modelPtr interface{}, respErr error) error {
	// todo: do better.
	var (
		b   []byte
		err error
	)

	// if response is nil and we saw an error
	if httpResp == nil && respErr != nil {
		return newAPIError(req, successCode, nil, nil, respErr)
	}

	// if we get here, we have some kind of response.  queue up body close...
	defer func() { _ = httpResp.Body.Close() }()

	// attempt to read all bytes from body
	if b, err = ioutil.ReadAll(httpResp.Body); err != nil {
		return newAPIError(req, successCode, httpResp, b, err)
	}

	// if returned code is not what was expected, build error
	if httpResp.StatusCode != successCode {
		return newAPIError(req, successCode, httpResp, b, err)
	}

	// finally, if necessary, attempt to unmarshal into response model
	if modelPtr != nil {
		if err = json.Unmarshal(b, modelPtr); err != nil {
			return newAPIError(req, successCode, httpResp, b, err)
		}
	}

	return nil
}
