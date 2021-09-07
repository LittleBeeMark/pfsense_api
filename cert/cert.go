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
MIIDojCCAoqgAwIBAgIIMqvJIdvMN/swDQYJKoZIhvcNAQELBQAwejELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MR8wHQYDVQQDExZLZXlN
YW5hZ2VyIFRlc3QgUlNBIENBMB4XDTIxMDgyMDA2NTY0NFoXDTIyMDgyMDA2NTY0
NFowHzELMAkGA1UEBhMCQ04xEDAOBgNVBAMTB21tbS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDZ2nyclm06xcg1uKqW2JMpAP8hjU9Zl59MQfzF
LBXJ3zEJn5x21kiRU4gRJiSSEocYfhAccf3L3mF7kro8aJjz6xem4OXvvP5CVHtU
9S86ZFzhnP/X+MT+3jkQhB8XhS5b+iw7B1FVhUsw5xRu1ZW+Bgc4uhFAFlbIYn5r
baLjMf778qsYtNrKHMBPSfORpoZqkSVZyOmEEbyy3dc/4NmIHTQQCSC1Y7bHb+LE
CaNllICmKwUfNdEexXB8MTJ0jumkUvArHD3FzB0kdtkBkfMIQcHrDpQQUHf8ePLH
4edm1nL//xhP3fUkXcwBEd237xtloLD1MT3H75jOAW4Dss4fAgMBAAGjgYYwgYMw
DgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAd
BgNVHQ4EFgQUmiN70WHD2SXhRQsh+p1nuiaVc4UwHwYDVR0jBBgwFoAUECM+bgg+
elE+HXbkMaIApFqFmZIwEgYDVR0RBAswCYIHbW1tLmNvbTANBgkqhkiG9w0BAQsF
AAOCAQEAQiv38MYOj0XQ1Qst8I6zNLaLGc7Vi3wOEjyTdXFP6F+cUQLNQp5JgZf9
3sa6L73Fh71HXTxrMyX2q3xZHWyb9uKqMJy/5z3NCZO8uVtHgUhu21JhkxV1z3VT
WV9fKSn5WyR6P/RUk0sGVrQ0F+E6rJTC9689E1Dbe8ocdp8b9HopeV2OcmMKQIZN
gYQlNi8lSCN0DWyB98ktTbMrbnomPutU6kNzFSoYkg2kmHMaOK84nVPMgX32+Vb8
IdkWsJSNOR2wHFUEos5hGLsJgg+m/+4RHL6jzfOiW8PWISnyuv+eityKrGI5Tjxy
a5KLqnLjWk1Rvu/LeJUsa7/laaH2Sw==
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
MIIEogIBAAKCAQEA2dp8nJZtOsXINbiqltiTKQD/IY1PWZefTEH8xSwVyd8xCZ+c
dtZIkVOIESYkkhKHGH4QHHH9y95he5K6PGiY8+sXpuDl77z+QlR7VPUvOmRc4Zz/
1/jE/t45EIQfF4UuW/osOwdRVYVLMOcUbtWVvgYHOLoRQBZWyGJ+a22i4zH++/Kr
GLTayhzAT0nzkaaGapElWcjphBG8st3XP+DZiB00EAkgtWO2x2/ixAmjZZSApisF
HzXRHsVwfDEydI7ppFLwKxw9xcwdJHbZAZHzCEHB6w6UEFB3/Hjyx+HnZtZy//8Y
T931JF3MARHdt+8bZaCw9TE9x++YzgFuA7LOHwIDAQABAoIBAD/MzbuqDjktHRIm
j8b3jDlw8kboNHnffqZ9mMJBw+vH8nuIA/GFleEBnpKvIfypcmkI2j0KYTJoYRIo
iWQRmeGtUtLrgEtyhN/2D9x0pa0rIUxthzu/vimJ+RpOJzDjLw1+uZ9b6ETscXXT
5tcCtATfjRPe4hhrsmSi+7UIebChNCKbQ8pv3fDDbDjnRBZiNAAlfTRJkBnba8W7
c6r61Jpe/bDKxpSphOPm+gGIVN88oNXv1xJwA7Qx0YoNPzFvm2kIjZ9LLDs9p10h
P/TolNGlKv6ac/q7qKEPRJw4UqAXvh2Y6jon62jHcdIwIcutxmlH4i18LECPiKa9
BOj9GXECgYEA+rRgR05eScB+xBevbm5msH/MFXIAD5lq0v9AkL12WHrfxk50XZnF
NS/E8vDOUFvVw/rzTbYyqV49jqdHb4ZuD4lCvBayCmmEMe56aaaTNri/ajEr4E5x
wwZyzRxXlDS+2ZOkrlHM0HLtsNahhmH1aOCyD1IrnZ864caLsGuE3CkCgYEA3nR5
6K7MoZhe+cxzi8V9xegVD1kiu27oFpiPvsLHjpLvIfiwNQlzgoU/2+eHKPLPXmZ2
rDBnY9MTxtY/8xlA3Z5TMUR2NhpFkC46eENkgnkf3yAM19WrLTVokvhXsfwUtymz
qYjgKG6MfsDvMAfb+ogMot9RvKO9g+1IDHMjoQcCgYAvoEiSAz9CP4FVezJmhi6X
5Q8+G7QLQpfakYcQeA2dbWpJX+oXRfkCy5pclIZ9GZUYb/n8j1o8dpy3FuwpMZ6C
8Q5ucNlNxRHJ8oXqwCxDPwGOCN1O9VgDNpxkfrfcfdCrwLKOMxf3mX2yFHQG9WEL
lXP+GRwUC4XCElfDIgnRUQKBgDK4vh821AO4eVddra7d7eqVG1Avk8LG6/ZS/NuT
D+tLR2koigzdxc+p0EC0ztWgX3X3yPFD7B8Pvr+klFo6lNazRebC5G07mkbgs4Y+
X4l8Uq8OYL9JwckCF4EDTQORJawJvyRVyD6PzksMdL0v3ZGHOdJdNwbbEtgk3zuv
eR07AoGANmcoQNUTkgczg9gKfysjCt2tCCR0Mk+rMiYE3q9LFUZT8uKue3EIXHPW
ATXb2bGle5s2ArSVtwGAjh3Cr8nlSpP+SlCxlBTKP/hrWsW5U+iIYRZiyAVF3KhW
3oLrXY39wqXrdry8TwbfX8PoYMrQVNtUM9zz5in+6pe5Myv7C7E=
-----END RSA PRIVATE KEY-----`

func (c *Cert) EditCertPage(certID string) (string, error) {
	op := &cli.Operation{
		HTTPMethod: "GET",
		HTTPPath:   "/system_certmanager.php?act=edit&id=%s",
	}
	op.HTTPPath = fmt.Sprintf(op.HTTPPath, certID)

	c.ReqActions.Unmarshal.PushBackNamed(unmarshalEditCertRespAction)
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
		return "", err
	}

	resp := &EditCertResp{}
	err = c.NewRequest(op, nil, reqB, resp).Send()
	if err != nil {
		return "", err
	}

	return resp.CSRF, nil
}

// EditCert doc
func (c *Cert) EditCert(csrf, cookie string) (*map[string]interface{}, error) {
	op := &cli.Operation{
		HTTPMethod:     "POST",
		HTTPPath:       "/system_certmanager.php",
		ShouldRedirect: true,
		Cookie:         cookie,
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", csrf)
	writer.WriteField("method", "edit")
	writer.WriteField("descr", "")
	writer.WriteField("import_type", "x509")
	writer.WriteField("cert", testCert)
	writer.WriteField("cert", testKey)
	op.ContentType = writer.FormDataContentType()
	op.ContentLen = strconv.Itoa(payload.Len())

	c.ReqActions.Unmarshal.PushBackNamed(unmarshalEditCertRespAction)
	c.CliInfo.SetHTTPClient(nil)

	output := &map[string]interface{}{}

	req := c.NewRequest(op, nil, payload.Bytes(), output)
	return output, req.Send()
}
