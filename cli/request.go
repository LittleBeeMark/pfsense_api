package cli

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var NoBody = http.NoBody

type CliInfo struct {
	Crd        *Credentials
	Endpoint   string
	HTTPClient *http.Client
}

type Credentials struct {
	UserName     string
	Password     string
	Cookie       string
	Csrf         string
	CookieJar    *cookiejar.Jar
	ProviderName string
}

//  Request doc
type Request struct {
	ClientInfo             *CliInfo
	Actions                *ReqAction
	Operation              *Operation
	Retryer                *Retry
	HTTPRequest            *http.Request
	HTTPResponse           *http.Response
	Params                 interface{}
	Error                  error
	Data                   interface{}
	DisableFollowRedirects bool
}

// Operation doc
type Operation struct {
	HTTPMethod string
	HTTPPath   string
}

// Retry doc
type Retry struct {
	Time  time.Duration
	Count int
}

// ShouldRetry doc
func (rt *Retry) ShouldRetry() bool {
	return rt != nil && rt.Count > 0 && rt.Time > 0
}

func New(cliInfo *CliInfo, reqActions *ReqAction, retryer *Retry,
	operation *Operation, params interface{}, data interface{}) *Request {

	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	httpReq, _ := http.NewRequest(method, "", nil)

	var err error
	httpReq.URL, err = url.Parse(cliInfo.Endpoint + operation.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		err = fmt.Errorf("InvalidEndpointURL, invalid endpoint uri : %v", err)
	}

	r := &Request{
		ClientInfo:  cliInfo,
		Actions:     reqActions,
		Operation:   operation,
		HTTPRequest: httpReq,
		Params:      params,
		Error:       err,
		Retryer:     retryer,
		Data:        data,
	}

	return r
}

// A Option is a functional option that can augment or modify a request when
// using a WithContext API operation method.
type Option func(*Request)

// ApplyOptions will apply each option to the request calling them in the order
// the were provided.
func (r *Request) ApplyOptions(opts ...Option) {
	for _, opt := range opts {
		opt(r)
	}
}

// Send doc
func (r *Request) Send() error {
	if err := r.Error; err != nil {
		return err
	}

	for {
		r.Error = nil
		if err := r.sendRequest(); err == nil {
			return nil
		}

		// 加入重试机制
		if r.Retryer.ShouldRetry() && r.Retryer.Count > 0 {
			r.Retryer.Count--
			continue
		}
		return r.Error
	}
}

func (r *Request) sendRequest() (sendErr error) {
	r.Actions.Send.Run(r)
	if r.Error != nil {
		// 日志
		return r.Error
	}

	//r.Actions.UnmarshalMeta.Run(r)
	//r.Actions.ValidateResponse.Run(r)
	if r.Error != nil {
		r.Actions.UnmarshalError.Run(r)
		// 日志

		return r.Error
	}

	r.Actions.Unmarshal.Run(r)
	if r.Error != nil {

		// 日志
		return r.Error
	}

	return nil
}
