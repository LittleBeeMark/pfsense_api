package haproxy

import (
	"bytes"
	"crypto/tls"
	"mime/multipart"
	"net/http"
	"pfsense/cli"
)

// Haproxy doc
type Haproxy struct {
	*cli.PfSense
}

// NewHaproxy doc
func NewHaproxy(userName, password, endPoint string) *Haproxy {
	return &Haproxy{
		cli.NewPfsense(userName, password, endPoint),
	}
}

// GetHaproxyList doc
func (hp *Haproxy) GetHaproxyList(cookie string) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/haproxy/haproxy_listeners.php",
		Cookie:     cookie,
	}

	hp.ReqActions.Unmarshal.PushBackNamed(unmarshalHaproxyListRespAction)
	hp.CliInfo.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		}}

	output := &map[string]interface{}{}

	req := hp.NewRequest(op, nil, nil, output)
	return output, req.Send()
}

// GetHaproxyInfo doc
func (hp *Haproxy) GetHaproxyInfo(id, cookie string) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/haproxy/haproxy_listeners_edit.php?id=mark",
		Cookie:     cookie,
	}

	output := &map[string]interface{}{}

	req := hp.NewRequest(op, nil, nil, output)
	return output, req.Send()
}

// ModifyParamReq doc
type ModifyParamReq struct {
	Csrf string
}

// ModifyProxy doc
func (hp *Haproxy) ModifyProxy(param *ModifyParamReq) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", param.Csrf)
	output := &map[string]interface{}{}

	req := hp.NewRequest(op, nil, payload.Bytes(), output)
	return output, req.Send()
}
