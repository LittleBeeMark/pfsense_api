package cert

import (
	"pfsense/haproxy"
)

// GetPfSenseCertList doc
func GetPfSenseCertList(userName, password, endPoint string) ([]*PfCert, error) {
	c := NewCert(userName, password, endPoint)
	return c.GetCertList()
}

// EditCertParam doc
type EditCertParam struct {
	UserName string
	Password string
	EndPoint string
	CertID   string
	Cert     string
	Key      string
}

// EditPfSenseCert doc
func EditPfSenseCert(param *EditCertParam) error {
	c := NewCert(param.UserName, param.Password, param.EndPoint)
	cookie, err := c.GetCookie()
	if err != nil {
		return err
	}

	c.ReqActions.Clear()
	op, err := c.EditCertPage(param.CertID, cookie)
	if err != nil {
		return err
	}

	c.ReqActions.Clear()
	err = c.EditCert(op.CSRF, op.Descr, cookie, param.CertID, param.Cert, param.Key)
	if err != nil {
		return err
	}

	// 重新启动 haproxy 服务
	h := haproxy.NewHaproxy(param.UserName, param.Password, param.EndPoint)
	hlResp, err := h.GetHaproxyList(cookie)
	if err != nil {
		return err
	}

	h.ReqActions.Clear()
	return h.RefreshHaproxy(hlResp.CSRF, cookie)
}
