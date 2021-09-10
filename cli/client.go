package cli

import (
	"bytes"
	"mime/multipart"
	"strconv"
)

var (
	UserName string = "admin"
	Password string = "pfsense"
	EndPoint string = "https://192.168.252.183"
)

type PfSense struct {
	CliInfo    *CliInfo
	ReqActions ReqAction
}

// NewPfsense  基础信息配置
func NewPfsense(userName, password, endPoint string) *PfSense {
	haproxy := &PfSense{
		CliInfo: &CliInfo{
			Crd: &Credentials{
				UserName: userName,
				Password: password,
			},
			Endpoint: endPoint,
		},
	}

	haproxy.ReqActions.Send.PushBackNamed(SendAction)
	return haproxy
}

// NewRequest doc
func (pf *PfSense) NewRequest(op *Operation, retryer *Retry, reqBody []byte, data interface{}) *Request {
	return New(pf.CliInfo, pf.ReqActions, retryer, op, reqBody, data)
}

// IndexResp doc
type IndexResp struct {
	CSRF   string `json:"csrf"`
	Cookie string `json:"cookie"`
}

// GetIndexReq doc
func (pf *PfSense) GetIndexReq() (*IndexResp, error) {
	op := &Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	pf.ReqActions.Unmarshal.PushBackNamed(unmarshalIndexRespAction)
	pf.ReqActions.UnmarshalError.PushBackNamed(UnmarshalErrPageBasicAction)
	pf.CliInfo.SetHTTPClient(nil)

	output := &IndexResp{}
	req := pf.NewRequest(op, nil, nil, output)
	err := req.Send()
	return output, err
}

// LoginResp doc
type LoginResp struct {
	SetCookie string `json:"set_cookie"`
}

// Login doc
func (pf *PfSense) Login(csrf, cookie string) (*LoginResp, error) {
	op := &Operation{
		HTTPMethod:     "POST",
		HTTPPath:       "/",
		ShouldRedirect: true,
		Cookie:         cookie,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", csrf)
	writer.WriteField("usernamefld", pf.CliInfo.Crd.UserName)
	writer.WriteField("passwordfld", pf.CliInfo.Crd.Password)
	writer.WriteField("login", "Sign In")
	op.ContentType = writer.FormDataContentType()
	op.ContentLen = strconv.Itoa(payload.Len())

	pf.ReqActions.Unmarshal.PushBackNamed(unmarshalLoginRespAction)
	pf.ReqActions.UnmarshalError.PushBackNamed(UnmarshalErrPageBasicAction)
	pf.CliInfo.SetHTTPClient(nil)

	output := &LoginResp{}
	req := pf.NewRequest(op, nil, payload.Bytes(), output)
	return output, req.Send()
}

// GetCookie doc
func (pf *PfSense) GetCookie() (string, error) {
	iresp, err := pf.GetIndexReq()
	if err != nil {
		return "", err
	}
	pf.ReqActions.Clear()

	lresp, err := pf.Login(iresp.CSRF, iresp.Cookie)
	if err != nil {
		return "", err
	}

	return lresp.SetCookie, err
}
