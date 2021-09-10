package cli

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type CliInfo struct {
	Crd        *Credentials
	Endpoint   string
	HTTPClient *http.Client
}

func (cli *CliInfo) SetHTTPClient(client *http.Client) {
	if client != nil {
		cli.HTTPClient = client
		return
	}

	cli.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		}}
}

type Credentials struct {
	UserName     string
	Password     string
	CookieJar    *cookiejar.Jar
	ProviderName string
}

// Request doc
type Request struct {
	ClientInfo             *CliInfo
	Operation              *Operation
	Retryer                *Retry
	HTTPRequest            *http.Request
	HTTPResponse           *http.Response
	Actions                ReqAction
	ReqBody                []byte
	RespBody               []byte
	Error                  error
	Data                   interface{}
	DisableFollowRedirects bool
}

// Operation doc
type Operation struct {
	HTTPMethod  string
	HTTPPath    string
	ContentType string
	ContentLen  string

	Cookie         string
	ShouldRedirect bool
}

// SetContent doc
func (o *Operation) SetContent(r *http.Request) {
	if o.ContentType != "" {
		r.Header.Set("Content-Type", o.ContentType)
	}

	if o.ContentLen != "" {
		r.Header.Set("Content-Length", o.ContentLen)
	}
}

// SetCookie doc
func (o *Operation) SetCookie(r *http.Request) {
	if o.Cookie != "" {
		r.Header.Set("Cookie", o.Cookie)
	}
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

func New(cliInfo *CliInfo, reqActions ReqAction, reqRetry *Retry,
	operation *Operation, reqBody []byte, data interface{}) *Request {
	if cliInfo.HTTPClient == nil {
		cliInfo.HTTPClient = http.DefaultClient
	}

	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	httpReq, _ := http.NewRequest(method, "", bytes.NewReader(reqBody))
	operation.SetContent(httpReq)
	operation.SetCookie(httpReq)

	var err error
	httpReq.URL, err = url.Parse(cliInfo.Endpoint + operation.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		err = fmt.Errorf("InvalidEndpointURL, invalid endpoint uri : %v", err)
	}
	fmt.Println("r.url", httpReq.URL)

	r := &Request{
		ClientInfo:             cliInfo,
		Actions:                reqActions,
		Operation:              operation,
		HTTPRequest:            httpReq,
		ReqBody:                reqBody,
		Retryer:                reqRetry,
		Data:                   data,
		Error:                  err,
		DisableFollowRedirects: operation.ShouldRedirect,
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

	r.Actions.Unmarshal.Run(r)
	if r.Error != nil {
		// 日志
		return r.Error
	}

	r.Actions.UnmarshalError.Run(r)
	if r.Error != nil {
		// 日志
		return r.Error
	}

	return nil
}
