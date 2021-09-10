package cert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"pfsense/cli"
	"strconv"
)

// Cert doc
type Cert struct {
	*cli.PfSense
}

// NewCert doc
func NewCert(userName, password, endPoint string) *Cert {
	return &Cert{
		cli.NewPfsense(userName, password, endPoint),
	}
}

// PfCert doc
type PfCert struct {
	Refid string `json:"refid"`
	Descr string `json:"descr"`
	Type  string `json:"type,omitempty"`
	Crt   string `json:"crt"`
	Prv   string `json:"prv"`
}

// GetCertList doc
func (c *Cert) GetCertList() ([]*PfCert, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/system/certificate",
	}

	c.ReqActions.Unmarshal.PushBackNamed(cli.UnmarshalAPIBasicAction)
	c.CliInfo.SetHTTPClient(nil)

	reqBody := struct {
		ClientID    string `json:"client-id"`
		ClientToken string `json:"client-token"`
	}{
		ClientID:    c.CliInfo.Crd.UserName,
		ClientToken: c.CliInfo.Crd.Password,
	}

	reqB, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp := &struct {
		Cert []interface{} `json:"cert"`
	}{}

	var pfCerts []*PfCert
	err = c.NewRequest(op, nil, reqB, resp).Send()
	if err != nil {
		return nil, err
	}

	// 因为类型不一样所以要多做一层 很坑
	for _, c := range resp.Cert {
		if _, ok := c.(map[string]interface{}); ok {
			var cert PfCert
			vb, err := json.Marshal(c)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(vb, &cert)
			if err != nil {
				return nil, err
			}
			pfCerts = append(pfCerts, &cert)
		}
	}
	return pfCerts, nil
}

var testCert = `-----BEGIN CERTIFICATE-----
MIIDqjCCApKgAwIBAgIIJCXlJCsFgQ0wDQYJKoZIhvcNAQELBQAwejELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MR8wHQYDVQQDExZLZXlN
YW5hZ2VyIFRlc3QgUlNBIENBMB4XDTIxMDkwODA4MTIyOFoXDTIyMDkwODA4MTIy
OFowIzELMAkGA1UEBhMCQ04xFDASBgNVBAMTC3Bmc2Vuc2UuY29tMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9su46UIH2yh2bSw5rnIDnGmst9lfx+j2
+e/nWNC1WOv3CBz484j9Xh1XelAK3Zdg1H1t1Y/Eh/N0+DTUsKYwerm6dj+efEpV
4uzy0Y6Apwl8Vvz58O/uOG6EEglfR0J1+vRYbXB5bT6U+mexZ3fDd9M7wqC/aQqX
Rjjsh/2aGMNzUNrULHddUYjMgB5oVB08RXOvNvZwyQh2NcMSuNBOkMWwy91WxwLd
UU8iWeQqVvDY0F6Q5nyCbyj4ZN/7EjoHrpRQ7nqPvyEVASlMqFV3ezJl3XbwNV1w
sTplNB4eBh5u6hg3c0areevQBOY7LmjSfRFq94mq1pGQNxx9e5CGWwIDAQABo4GK
MIGHMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUH
AwIwHQYDVR0OBBYEFBvm90vNS2L6iGmLIQ3mn0cCLZwCMB8GA1UdIwQYMBaAFBAj
Pm4IPnpRPh125DGiAKRahZmSMBYGA1UdEQQPMA2CC3Bmc2Vuc2UuY29tMA0GCSqG
SIb3DQEBCwUAA4IBAQA+O4en3KnjyAmq01qOLBDr5yNMrI+6AwxfloC8NH7bMZJT
YNUPRYKei4OaLE8sj3x6LF40zeHilcEein6IQaU291PkUeWgh7N/KLvARfHaUpCX
AdINzD7+W7/dOkzwc4uS0oIJjLYFdQmFYGYwryfHzZ/r8G3iR90qyVVhozuFVnBh
3Dgwlgn1mU3pqkToC61mmgSKgzO+fmC8W9o52LhPsJ3PjpEPUA2lmppgvW008OXM
cD6z+bPHYxQz47WZJOFLlB3bSnNXFlwBd2AE0Hn6JLLQyreT9e88/MgtsLsnAb03
76Z9ljNiASWDTw2kxRiL9T64n5iSb4Y7PAJKZHN/
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIID2jCCAsKgAwIBAgIIeSI3M4rCCEswDQYJKoZIhvcNAQELBQAwezELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MSAwHgYDVQQDExdLZXlN
YW5hZ2VyIFRlc3QgUm9vdCBDQTAeFw0yMTA3MTQwNTEyMDBaFw0zMTA3MTQwNTEy
MDBaMHoxCzAJBgNVBAYTAkNOMRcwFQYDVQQKEw5LZXlNYW5hZ2VyLm9yZzExMC8G
A1UECxMoS2V5TWFuYWdlciBUZXN0IFJvb3QgLSBGb3IgVGVzdCBVc2UgT25seTEf
MB0GA1UEAxMWS2V5TWFuYWdlciBUZXN0IFJTQSBDQTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBANgs0N+IrziLtph3gDHpapH7Wn4moycKDi9ymb5FHBpj
Gs2TRIn7uJMhFbAJklcdN9usbXjgWkmP2oFdfTJsQQyvF8KcNpSL8OMSS8zy79sU
jV+VvW0w0Uv43lBrPVkXF2c2AhkguWFb5DjunFxBCY5PnEVUP3wYBOsJW6HCvUx8
tuICkfBudOvf612YtEixA2GAig6kviTDrxBbN2QsrrnZEyaAnR6+rDAbLzmVK+Px
a6JxbTQlC+qrjDV/o09XnbLWtVE8U67D4IhYfT0HIsugTRGUIkff49CQrAEZ0/pi
RWln7KSoM+rDeDp08ORWZsQ7Mx5G//Zep6evhQuCnWcCAwEAAaNjMGEwDgYDVR0P
AQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFBAjPm4IPnpRPh12
5DGiAKRahZmSMB8GA1UdIwQYMBaAFPKcEGuzWGA4xJynrNCqlL6+wFpQMA0GCSqG
SIb3DQEBCwUAA4IBAQA7r6WhiJerQweBsL+sUaPti1O5cR30lpw7YxhGERQsik7h
pRhc9tE+PWGzepBhx1tN9pq8N9lS+Mbcx2oOJOc8e1RPd3q3meU76868OuTSKtD+
mV3pcJ+rnbWr2pD1FWu5GDn1/5cmpNXotha+pWIZpGJ8lVtrhJCwmH9hUFFkxKr/
dXsis03TzIcbH5fyiJJOIXf77IunzFIgmwA4wBLOD+aRPP1wMa2Q/dagbvzRbx5g
sFw0xuJyZInrpC7czwz8AK2NLCbzeMLQL5HvQu71AhqysXQDBFZjdAQ2IEC8ZMsq
ew6BXkqiZDA5jg05jQSGydiIedO1ScXfyd8nMZls
-----END CERTIFICATE-----`

var testKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA9su46UIH2yh2bSw5rnIDnGmst9lfx+j2+e/nWNC1WOv3CBz4
84j9Xh1XelAK3Zdg1H1t1Y/Eh/N0+DTUsKYwerm6dj+efEpV4uzy0Y6Apwl8Vvz5
8O/uOG6EEglfR0J1+vRYbXB5bT6U+mexZ3fDd9M7wqC/aQqXRjjsh/2aGMNzUNrU
LHddUYjMgB5oVB08RXOvNvZwyQh2NcMSuNBOkMWwy91WxwLdUU8iWeQqVvDY0F6Q
5nyCbyj4ZN/7EjoHrpRQ7nqPvyEVASlMqFV3ezJl3XbwNV1wsTplNB4eBh5u6hg3
c0areevQBOY7LmjSfRFq94mq1pGQNxx9e5CGWwIDAQABAoIBAHLSR56mrPcG9UpH
yDJkM0/jTote5BQjTDzuo+kLeUP8zLsGl8DenFrcNeXfNZ7xHgjermos8Ff8KhCM
q+Pf/jArFxiK8taK0zi32yUqMqsEW8gw/KxfRKsWp/WoTJ7fyfpPEpEgQi5UboWu
Fri33MZ0DZq4pTVeyxDNzSZcFASuiKJaOgBjlayZ8sWa8LDNiIy06OUs6FOodE//
srYdQ79mv6yzJlRljHKky7TIqI9+gtN+6LBb9AT89+lkuc9vo6QtM/k9f93aDxYu
REtx0CLj69FEiaTb8hr5ioTOrDA8O1JA8DCQcfeMaBL2LriZIVnnFf4zLi/Mf35S
9i326MECgYEA/fRhz9YYswVm6iU6zBfozZ5yGh5P9TD7bMklLzF2keQDok6kCy6G
jfBNqmtQ14ny3voCbS64QTwojXUlyAbvzLch36PKiS73bepSgq/A3hlznwr62f8z
Pyc6/ojqdUrHA2WAhGtzHtpaDCEsGGg8YGxTuCznm6QlquMLCh3WgaECgYEA+MiU
axK15fcz0JODe6Co/WXVTBn4WweicUUFgVkwoboLa2UHOuzXJZc/7mKWVa3km/Ia
LfvaJeeSoi21hWpE7GhSto3rDE2RH+EaJ2DIBO+gnNDaWTdRiX7ROCcF8YABBXP7
t1mqGkM2DSQQGI7hHAHGm3toXiTbNwFhnnrTfnsCgYAdaEAvVgZ0LIr3UCpX2eju
boud9KajqPFkboJszZiCWloFYH/kf5W7N/l2Y4jv/JxwL5k6VW+mtjWn88nVOOBp
30b/47KfYh3qz3iQg5Oc5GucHiRgPAoOJRfSw74KqQcUoJtaOisRho1o3CqEBQYA
0Gp9aE3FmyN1f2cW28+sIQKBgC2eEeDkSGhHgL/BHL3hgrscHhVbObfvWXDtmAnU
wd8Vzxw9JSs/3F9vMXDTsP866I1TwksmQCTtJm0Idp1lFAhJRMlmkm/qFS9ERlhs
HaESE2BNx4vOYewVTeW++g8DSqymTMSc81zncBTOxQjLwikxhipiYYvJtEyMu0qk
+PzbAoGAXZVUSXiFLvQbPuS9h/Rmq/k7o6OsJql1ldXL/W8dsNaKF9VcZjlpqOEI
Z06qJEEw9Jt/CUTsCuvrghVDkQYx/YwyeR5p1wcLnzMc9J0OwEWq86aNuyVNP3n9
0LgfhUGruFQL//XnRkJ7sKTJwyDHXzR2xCNL6n2guKSX+bx2bKo=
-----END RSA PRIVATE KEY-----`

func (c *Cert) EditCertPage(certID, cookie string) (*EditCertResp, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   fmt.Sprintf("/system_certmanager.php?act=edit&id=%s", certID),
		Cookie:     cookie,
	}
	c.ReqActions.Unmarshal.PushBackNamed(unmarshalEditCertRespAction)
	c.ReqActions.UnmarshalError.PushBackNamed(cli.UnmarshalErrPageBasicAction)
	c.CliInfo.SetHTTPClient(nil)

	resp := &EditCertResp{}
	err := c.NewRequest(op, nil, nil, resp).Send()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// EditCert doc
func (c *Cert) EditCert(csrf, desc, cookie, certID, cert, key string) error {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/system_certmanager.php",
		Cookie:     cookie,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", csrf)
	writer.WriteField("method", "edit")
	writer.WriteField("descr", desc)
	writer.WriteField("import_type", "x509")
	writer.WriteField("cert", cert)
	writer.WriteField("key", key)
	writer.WriteField("id", certID)
	writer.WriteField("save", "Save")
	op.ContentType = writer.FormDataContentType()
	op.ContentLen = strconv.Itoa(payload.Len())

	c.ReqActions.Unmarshal.PushBackNamed(cli.UnmarshalBasicAction)
	c.ReqActions.UnmarshalError.PushBackNamed(cli.UnmarshalErrPageBasicAction)
	c.CliInfo.SetHTTPClient(nil)

	return c.NewRequest(op, nil, payload.Bytes(), nil).Send()
}
