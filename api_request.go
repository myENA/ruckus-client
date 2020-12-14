package bigdog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	uriPathParameterSearchFormat  = "{%s}"
	uriQueryParameterPrefixFormat = "%s?%s"

	apiRequestURLFormat = "%s%s%s"

	headerKeyContentType         = "Content-Type"
	headerKeyContentDisposition  = "Content-Disposition"
	headerKeyAccept              = "Accept"
	headerValueApplicationJSON   = "application/json"
	headerValueMultipartFormData = "multipart/form-data"
)

type RequestMutator func(*http.Request)

func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int:
		return strconv.Itoa(value.(int))
	case bool:
		return strconv.FormatBool(value.(bool))
	case []string:
		return strings.Join(value.([]string), ",")
	case []int:
		l := len(value.([]int))
		if l == 0 {
			return ""
		}
		tmp := make([]string, l, l)
		for i, v := range value.([]int) {
			tmp[i] = strconv.Itoa(v)
		}
		return ToString(tmp)
	default:
		panic(fmt.Sprintf("unable to handle type %T", value))
	}
}

type PathValues map[string]string

func (pv PathValues) Set(key string, value interface{}) {
	pv[key] = ToString(value)
}

func (pv PathValues) Get(key string) string {
	v, _ := pv[key]
	return v
}

type QueryValues url.Values

func (qv QueryValues) Values() url.Values {
	return url.Values(qv)
}

func (qv QueryValues) Get(key string) string {
	return url.Values(qv).Get(key)
}

func (qv QueryValues) Set(key string, value interface{}) {
	url.Values(qv).Set(key, ToString(value))
}

func (qv QueryValues) Add(key string, value interface{}) {
	url.Values(qv).Add(key, ToString(value))
}

func (qv QueryValues) Del(key string) {
	url.Values(qv).Del(key)
}

func (qv QueryValues) Encode() string {
	return url.Values(qv).Encode()
}

var apiRequestID uint64

type APIRequest struct {
	Method        string
	URI           string
	Authenticated bool
	QueryParams   QueryValues
	PathParams    PathValues
	Header        http.Header

	id   uint64
	body io.Reader
	mpw  *multipart.Writer
}

var apiRequestPool = sync.Pool{New: func() interface{} { return new(APIRequest) }}

func bootstrapRequest(req *APIRequest, method, uri string, authenticated bool) {
	req.Method = method
	req.URI = uri
	req.Authenticated = authenticated
	req.QueryParams = make(QueryValues)
	req.PathParams = make(PathValues)
	req.Header = make(http.Header)

	req.id = atomic.AddUint64(&apiRequestID, 1)
}

func apiRequestFromPool(method, uri string, authenticated bool) *APIRequest {
	req := apiRequestPool.Get().(*APIRequest)
	bootstrapRequest(req, method, uri, authenticated)
	return req
}

func recycleAPIRequest(req *APIRequest) {
	req.Method = ""
	req.URI = ""
	req.Authenticated = false
	req.QueryParams = nil
	req.PathParams = nil
	req.Header = nil

	req.id = 0
	req.body = nil
	req.mpw = nil

	apiRequestPool.Put(req)
}

func NewAPIRequest(method, uri string, authenticated bool) *APIRequest {
	req := new(APIRequest)
	bootstrapRequest(req, method, uri, authenticated)
	return req
}

func (r *APIRequest) ID() uint64 {
	return r.id
}

func (r *APIRequest) SetBody(body interface{}) error {
	// test for reader
	if ar, ok := body.(io.Reader); ok {
		r.body = ar
		return nil
	}

	// test for raw bytes
	if b, ok := body.([]byte); ok {
		r.body = bytes.NewBuffer(b)
		return nil
	}

	// test for form data
	if v, ok := body.(url.Values); ok {
		r.body = bytes.NewBufferString(v.Encode())
		return nil
	}

	// finally, attempt json marshal
	if b, err := json.Marshal(body); err != nil {
		return err
	} else {
		r.body = bytes.NewBuffer(b)
		return nil
	}
}

func (r *APIRequest) Body() io.Reader {
	return r.body
}

func (r *APIRequest) MultipartForm() {
	r.body = new(bytes.Buffer)
	r.mpw = multipart.NewWriter(r.body.(*bytes.Buffer))
}

func (r *APIRequest) AddMultipartFile(key, filename string, f io.Reader) error {
	w, err := r.mpw.CreateFormFile(key, filename)
	if err != nil {
		return fmt.Errorf("error creating multipart form file part with key=%q and filename=%q: %w", key, filename, err)
	}
	if _, err = io.Copy(w, f); err != nil {
		return fmt.Errorf("error copying bytes from file %q to multipart writer: %w", filename, err)
	}
	addContentDispositionHeader(r, key, filename)
	return nil
}

func (r *APIRequest) AddMultipartField(key string, value interface{}) error {
	var (
		vr io.Reader
		ok bool
	)
	w, err := r.mpw.CreateFormField(key)
	if err != nil {
		return fmt.Errorf("error creating form field part with key=%q: %w", key, err)
	}
	vr, ok = value.(io.Reader)
	if !ok {
		if b, ok := value.([]byte); ok {
			vr = bytes.NewBuffer(b)
		} else if s, ok := value.(string); ok {
			vr = bytes.NewBufferString(s)
		} else if b, err := json.Marshal(value); err != nil {
			return fmt.Errorf("error marshalling form field part with key=%q and type=%T: %w", key, value, err)
		} else {
			vr = bytes.NewBuffer(b)
		}
	}
	if _, err = io.Copy(w, vr); err != nil {
		return fmt.Errorf("error copying bytes into multipart writer for key=%q with type=%T: %w", key, value, err)
	}
	return nil
}

func (r *APIRequest) AddMultipartFieldsFromValues(values url.Values) error {
	for k, vs := range values {
		for _, v := range vs {
			if err := r.AddMultipartField(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// CompiledURI will return to you the full request URI, not including scheme, hostname, and port.  This method is not
// thread safe, as you shouldn't be calling this asynchronously anyway.
func (r *APIRequest) CompiledURI() string {
	uri := r.URI
	if len(r.PathParams) > 0 {
		for k, v := range r.PathParams {
			uri = strings.Replace(uri, fmt.Sprintf(uriPathParameterSearchFormat, k), url.PathEscape(v), 1)
		}
	}
	if len(r.QueryParams) > 0 {
		uri = fmt.Sprintf(uriQueryParameterPrefixFormat, uri, r.QueryParams.Encode())
	}
	return uri
}

// ToHTTP will attempt to construct an executable http.request
func (r *APIRequest) ToHTTP(ctx context.Context, addr, pathPrefix, authParamName, authParamValue string) (*http.Request, error) {
	var (
		compiledURL string
		httpRequest *http.Request
		err         error
	)

	if r.Authenticated {
		if authParamName == "" || authParamValue == "" {
			return nil, errors.New("authParam cannot be empty with authenticated request")
		}
		r.QueryParams.Set(authParamName, authParamValue)
	}

	compiledURL = fmt.Sprintf(apiRequestURLFormat, addr, pathPrefix, r.CompiledURI())

	if r.mpw != nil {
		r.Header.Set(headerKeyContentType, r.mpw.FormDataContentType())
		if err = r.mpw.Close(); err != nil {
			return nil, fmt.Errorf("error closing multipart writer: %w", err)
		}
	}

	if httpRequest, err = http.NewRequest(r.Method, compiledURL, r.Body()); err != nil {
		return nil, err
	}

	for header, values := range r.Header {
		for _, value := range values {
			httpRequest.Header.Add(header, value)
		}
	}

	return httpRequest.WithContext(ctx), nil
}

func addContentDispositionHeader(req *APIRequest, key, filename string) {
	req.Header.Add(
		headerKeyContentDisposition,
		fmt.Sprintf(
			"form-data: name=%s; filename=%s;",
			strconv.Quote(key),
			strconv.Quote(filename),
		),
	)
}
