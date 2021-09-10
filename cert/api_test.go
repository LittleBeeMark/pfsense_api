package cert

import (
	"fmt"
	"pfsense/cli"
	"pfsense/haproxy"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestCert_GetCertList(t *testing.T) {
	//cookie, err := cli.GetLoginCookie()
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return
	//}
	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)

	op, err := c.GetCertList()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	for _, c := range op {
		fmt.Println(*c)
	}
}

func TestCert_EditCertPage(t *testing.T) {
	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)
	cookie, err := c.GetCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	op, err := c.EditCertPage("613356ef7ac58", cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(op)

}

func Test_GETDescr(t *testing.T) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(""))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	reader.Find("input#descr").Attr("value")
	//if b {
	//	fmt.Println(val)
	//}

}

func TestCert_EditCert2(t *testing.T) {
	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)
	cookie, err := c.GetCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	c.ReqActions.Clear()
	op, err := c.EditCertPage("613356ef7ac5", cookie)
	if err != nil {
		fmt.Println("err edit: ", err)
		return
	}

	c.ReqActions.Clear()
	err = c.EditCert(op.CSRF, op.Descr, cookie, "", "", "613356ef7ac58")
	if err != nil {
		fmt.Println("err edit cert: ", err)
		return
	}

}

func TestCert_EditCert(t *testing.T) {
	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)
	cookie, err := c.GetCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	c.ReqActions.Clear()
	op, err := c.EditCertPage("613356ef7ac58", cookie)
	if err != nil {
		fmt.Println("err edit: ", err)
		return
	}

	c.ReqActions.Clear()
	err = c.EditCert(op.CSRF, op.Descr, cookie, "", "", "613356ef7ac58")
	if err != nil {
		fmt.Println("err edit cert: ", err)
		return
	}

	// 重新启动 haproxy 服务
	h := haproxy.NewHaproxy(cli.UserName, cli.Password, cli.EndPoint)
	hlResp, err := h.GetHaproxyList(cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	h.ReqActions.Clear()
	err = h.RefreshHaproxy(hlResp.CSRF, cookie)
	if err != nil {
		fmt.Println("refresh err", err)
		return
	}

	fmt.Println("hlResp : ", hlResp)
}
