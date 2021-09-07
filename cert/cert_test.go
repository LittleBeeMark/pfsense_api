package cert

import (
	"fmt"
	"pfsense/cli"
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

	fmt.Println(op)
}

func TestCert_EditCertPage(t *testing.T) {
	cookie, err := cli.GetLoginCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)
	op, err := c.EditCertPage("613356ef7ac58", cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(op)

}

func Test_GETDescr(t *testing.T) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	reader.Find("input#descr").Attr("value")
	//if b {
	//	fmt.Println(val)
	//}

}

const (
	FollowingError = "The following input errors were detected"
)

func TestCert_EditCert(t *testing.T) {
	cookie, err := cli.GetLoginCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)
	op, err := c.EditCertPage("613356ef7ac58", cookie)
	if err != nil {
		fmt.Println("err edit: ", err)
		return
	}

	c = NewCert(cli.UserName, cli.Password, cli.EndPoint)
	resp, err := c.EditCert(op.CSRF, op.Descr, cookie, "613356ef7ac58")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(resp)
}
