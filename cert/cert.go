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

	c.ReqActions.Unmarshal.PushBackNamed(cli.UnmarshalAPIBasic)
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
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
YTHRA4YS9VvyXGNBWbn/6x560Q==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIErjCCA5agAwIBAgIQBYAmfwbylVM0jhwYWl7uLjANBgkqhkiG9w0BAQsFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0xNzEyMDgxMjI4MjZaFw0yNzEyMDgxMjI4MjZaMHIxCzAJBgNVBAYTAkNO
MSUwIwYDVQQKExxUcnVzdEFzaWEgVGVjaG5vbG9naWVzLCBJbmMuMR0wGwYDVQQL
ExREb21haW4gVmFsaWRhdGVkIFNTTDEdMBsGA1UEAxMUVHJ1c3RBc2lhIFRMUyBS
U0EgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCgWa9X+ph+wAm8
Yh1Fk1MjKbQ5QwBOOKVaZR/OfCh+F6f93u7vZHGcUU/lvVGgUQnbzJhR1UV2epJa
e+m7cxnXIKdD0/VS9btAgwJszGFvwoqXeaCqFoP71wPmXjjUwLT70+qvX4hdyYfO
JcjeTz5QKtg8zQwxaK9x4JT9CoOmoVdVhEBAiD3DwR5fFgOHDwwGxdJWVBvktnoA
zjdTLXDdbSVC5jZ0u8oq9BiTDv7jAlsB5F8aZgvSZDOQeFrwaOTbKWSEInEhnchK
ZTD1dz6aBlk1xGEI5PZWAnVAba/ofH33ktymaTDsE6xRDnW97pDkimCRak6CEbfe
3dXw6OV5AgMBAAGjggFPMIIBSzAdBgNVHQ4EFgQUf9OZ86BHDjEAVlYijrfMnt3K
AYowHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUwDgYDVR0PAQH/BAQD
AgGGMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjASBgNVHRMBAf8ECDAG
AQH/AgEAMDQGCCsGAQUFBwEBBCgwJjAkBggrBgEFBQcwAYYYaHR0cDovL29jc3Au
ZGlnaWNlcnQuY29tMEIGA1UdHwQ7MDkwN6A1oDOGMWh0dHA6Ly9jcmwzLmRpZ2lj
ZXJ0LmNvbS9EaWdpQ2VydEdsb2JhbFJvb3RDQS5jcmwwTAYDVR0gBEUwQzA3Bglg
hkgBhv1sAQIwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQuY29t
L0NQUzAIBgZngQwBAgEwDQYJKoZIhvcNAQELBQADggEBAK3dVOj5dlv4MzK2i233
lDYvyJ3slFY2X2HKTYGte8nbK6i5/fsDImMYihAkp6VaNY/en8WZ5qcrQPVLuJrJ
DSXT04NnMeZOQDUoj/NHAmdfCBB/h1bZ5OGK6Sf1h5Yx/5wR4f3TUoPgGlnU7EuP
ISLNdMRiDrXntcImDAiRvkh5GJuH4YCVE6XEntqaNIgGkRwxKSgnU3Id3iuFbW9F
UQ9Qqtb1GX91AJ7i4153TikGgYCdwYkBURD8gSVe8OAco6IfZOYt/TEwii1Ivi1C
qnuUlWpsF1LdQNIdfbW3TSe0BhQa7ifbVIfvPWHYOu3rkg1ZeMo6XRU9B4n5VyJY
RmE=
-----END CERTIFICATE-----`

var testKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1Qemi451qOiC8PS2qJIuQKpuV4Ct2rRG9Cdq/RiMnbezLhC6
m1+CRMdUf8tgWiIvTr72As2jsEi8hYEb5vPJCFWVTleQZhQzxOQ5WPCVkXdoZEsX
nEaW3eDEuVfailcyNXMnpPrQ8WmEqnAZyqo31gbdXvg4lX4doAXZQYTW4UTMn4eU
yUITsdA9VqRKNkFprvpYH6zoYZho58/CJ5T0+UuVc7FbTP1suaHiHQyP4YlJG4Gd
dQiaB/D447SDti4CU5PlaEdSUvR4xFCKzyILExYG96bbOmptFuJDI1YjcB7NrARF
dG1F/2R8AsP7QWF9gVX162fu9WC3QKkU6+FBZwIDAQABAoIBAAw3/x58ElWY2IHX
l259y/dbjy6nq9Eii/tPE3zm3pHhSn626O0mgkNNp5QY5zLkSRBcNaIdh2kCEwYZ
TK9nhj1bI1A2B4tDV2DQslVen0WTssIl6HnHNroNPVtgJaHPtIqKS1jUJg3ivoBE
I8aTEtbM1/+tfxlb/XkDPN3cL7MF82iPxEE5qbGSbQ7XxsfVlVnDqVeo/JPAJInj
BgtzRFT16bCX8E+azk+MPDjGfynlRozEkTqE/qwg9pGjWEZU+YAWX5CNMU1q2cQ7
QxThcWFuusGYBeIyi6oC1Bey6f0BaVT5/8HXnW4F/zvts55vznhsHP14Lv0UjI2b
bH0FopkCgYEA6e3eNmvnE42hAnsUlSYRcY4Y1DVZs/+gvPHCdX44CmxOfmFePmRg
aJnA35YYJXF7wUBNeWLKTJalkXVNJyDXrV13x4OOu5ZuG9himGxJ9fyaneGe5xrI
F/Z8aFiQoF9to0Lit61KZMa6eaAf9y53F1JAuXNSgltKU8OHmSjvcIUCgYEA6SD/
mYwDBEeXC18Ud1HzRtUDWvahGTo22hCkW2hfCtDwYO3YvOoJGgeY1g4XOH3HoF7W
1vj2PHn37hk2ZsCztnsfv/v19bLpcM1ZZHMEVOi1jssQgImS0H3euxaR6Tm7J2NA
bDa7aZavghV77xazT1qPqF98kcSu9v1e5/J44/sCgYEAg1RosZLoewDDSPpCdu2V
U3QcGl8NSlGUMx9jNcUzvx5I/wi4+TwvJ+pR7vR7/+FzQe5Q0HoW6uKUd1Isi3AT
xZ+41EjWhBgTHwzfZYU+DJzRTRvSsxpFBbb65zX5lB+fFM5DLC1cm7E6FXhBGet7
Lpud/L9yrf6EfvxCD0l9C2UCgYAlcl58LIhDhLhXZENmUyyPoSGz18/SvI4ZAvlT
kXVKyHsEJtBcr/8cRkIfiA3kzhdlxDYgO8dZuYHAph4d7TAwzKAx92fvunhS3TlR
sEPu255mPGn/K5oAkWdYh+ySGOreDcxIVZZPBJxWedr5cZ0FzxcqRYQ96ejs1ZAM
E9+lRwKBgEgqDhq2ZP5IvErZtC1fw4qfKRgqHwvNikYvSUZ900HKNk4EmFbXrW7+
Ol0xAPMkMOxZRY6k4BOW9jdACt7GbDgVAScu8oyDLTAvMo+vpHOkHCrfb+zUBVKd
0lPEB4jSdJ9phmesKwWs3FzJZfu6Fwj+N9BooNfMi3gNnqZGycaI
-----END RSA PRIVATE KEY-----`

func (c *Cert) EditCertPage(certID, cookie string) (*EditCertResp, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   fmt.Sprintf("/system_certmanager.php?act=edit&id=%s", certID),
		Cookie:     cookie,
	}

	c.ReqActions.Unmarshal.PushBackNamed(unmarshalEditCertRespAction)
	c.CliInfo.SetHTTPClient(nil)

	resp := &EditCertResp{}
	err := c.NewRequest(op, nil, nil, resp).Send()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// EditCert doc
func (c *Cert) EditCert(csrf, dsc, cookie, certID string) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod: "POST",
		HTTPPath:   "/system_certmanager.php",
		Cookie:     cookie,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", csrf)
	writer.WriteField("method", "edit")
	writer.WriteField("descr", "zzzz.com")
	writer.WriteField("import_type", "x509")
	writer.WriteField("cert", testCert)
	writer.WriteField("key", testKey)
	writer.WriteField("id", certID)
	writer.WriteField("save", "Save")
	op.ContentType = writer.FormDataContentType()
	op.ContentLen = strconv.Itoa(payload.Len())

	c.ReqActions.Unmarshal.PushBackNamed(unmarshalEditRespAction)
	c.CliInfo.SetHTTPClient(nil)

	output := &map[string]interface{}{}

	req := c.NewRequest(op, nil, payload.Bytes(), output)
	return output, req.Send()
}
