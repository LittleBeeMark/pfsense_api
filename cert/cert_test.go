package cert

import (
	"fmt"
	"pfsense/cli"
	"testing"
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
	c := NewCert(cli.UserName, cli.Password, cli.EndPoint)

	op, err := c.EditCertPage("613356ef7ac58")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(op)

}
