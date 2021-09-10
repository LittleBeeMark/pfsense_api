package haproxy

import (
	"fmt"
	"testing"
)

func TestHaproxyList(t *testing.T) {
	ha := NewHaproxy(UserName, Password, EndPoint)
	cookie, err := ha.GetCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	op, err := ha.GetHaproxyList(cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(op)
}
