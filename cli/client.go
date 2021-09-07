package cli

import (
	"bytes"
	"mime/multipart"
	"strconv"
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

// GetIndexReq doc
func (pf *PfSense) GetIndexReq() (*map[string]interface{}, error) {
	op := &Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	pf.ReqActions.Unmarshal.PushBackNamed(unmarshalIndexRespAction)
	pf.CliInfo.SetHTTPClient(nil)

	output := &map[string]interface{}{}
	req := pf.NewRequest(op, nil, nil, output)
	err := req.Send()
	return output, err
}

// Login doc
func (pf *PfSense) Login(csrf, cookie string) (*map[string]interface{}, error) {
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
	pf.CliInfo.SetHTTPClient(nil)

	output := &map[string]interface{}{}

	req := pf.NewRequest(op, nil, payload.Bytes(), output)
	return output, req.Send()
}
