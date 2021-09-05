package haproxy

import "pfsense/cli"

type Haproxy struct {
	CliInfo    *cli.CliInfo
	ReqActions *cli.ReqAction
}

// NewHaproxy doc
func NewHaproxy(userName, password, endPoint string) *Haproxy {
	haproxy := &Haproxy{
		CliInfo: &cli.CliInfo{
			Crd: &cli.Credentials{
				UserName: userName,
				Password: password,
			},
			Endpoint: endPoint,
		},
	}

	haproxy.ReqActions.Send.PushBackNamed(cli.SendAction)
	return haproxy
}

// newRequest doc
func (hp *Haproxy) newRequest(op *cli.Operation, retryer *cli.Retry, params, data interface{}) *cli.Request {
	return cli.New(hp.CliInfo, hp.ReqActions, retryer, op, params, data)
}

// GetIndexReq doc
func (hp *Haproxy) GetIndexReq(input *map[string]interface{}) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/index.html",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output := &map[string]interface{}{}

	req := hp.newRequest(op, nil, input, output)
	return output, req.Send()
}

// GetDomains doc
func (hp *Haproxy) Login(input *map[string]interface{}) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output := &map[string]interface{}{}

	req := hp.newRequest(op, nil, input, output)
	return output, req.Send()
}

// GetCertificatesRequest doc
func (hp *Haproxy) GetHaproxyList(input *map[string]interface{}) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output := &map[string]interface{}{}

	req := hp.newRequest(op, nil, input, output)
	return output, req.Send()
}

// GetCertificates doc
func (hp *Haproxy) GetHaproxyInfo(input *map[string]interface{}) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output := &map[string]interface{}{}

	req := hp.newRequest(op, nil, input, output)
	return output, req.Send()
}

// ModifyDomainCertRequest doc
func (hp *Haproxy) ModifyProxy(input *map[string]interface{}) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output := &map[string]interface{}{}

	req := hp.newRequest(op, nil, input, output)
	return output, req.Send()
}
