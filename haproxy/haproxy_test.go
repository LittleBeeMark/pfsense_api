package haproxy

import (
	"fmt"
	"pfsense/cli"
	"testing"
)

func TestHaproxyList(t *testing.T) {
	cookie, err := cli.GetLoginCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	op, err := HaproxyList(cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(op)
}
