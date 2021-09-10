package haproxy

import (
	"bytes"
	"mime/multipart"
	"pfsense/cli"
	"strconv"
)

var (
	UserName string = "admin"
	Password string = "pfsense"
	EndPoint string = "https://192.168.252.183"
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

// ListResp doc
type ListResp struct {
	CSRF       string   `json:"csrf"`
	FrontNames []string `json:"front_names"`
}

// GetHaproxyList doc
func (hp *Haproxy) GetHaproxyList(cookie string) (*ListResp, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/haproxy/haproxy_listeners.php",
		Cookie:     cookie,
	}

	hp.ReqActions.Unmarshal.PushBackNamed(unmarshalHaproxyListNameRespAction)
	hp.ReqActions.UnmarshalError.PushBackNamed(cli.UnmarshalErrPageBasicAction)
	hp.CliInfo.SetHTTPClient(nil)

	output := &ListResp{}

	req := hp.NewRequest(op, nil, nil, output)
	return output, req.Send()
}

// RefreshHaproxy doc
func (hp *Haproxy) RefreshHaproxy(csrf, cookie string) error {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/status_services.php",
		Cookie:     cookie,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", csrf)
	writer.WriteField("ajax", "ajax")
	writer.WriteField("mode", "restartservice")
	writer.WriteField("service", "haproxy")
	op.ContentType = writer.FormDataContentType()
	op.ContentLen = strconv.Itoa(payload.Len())

	hp.ReqActions.Unmarshal.PushBackNamed(cli.UnmarshalBasicAction)
	hp.ReqActions.UnmarshalError.PushBackNamed(cli.UnmarshalErrPageBasicAction)
	hp.CliInfo.SetHTTPClient(nil)
	return hp.NewRequest(op, nil, payload.Bytes(), nil).Send()
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
