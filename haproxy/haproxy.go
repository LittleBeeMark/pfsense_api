package haproxy

import (
	"fmt"
)

func GetHaproxyFrontNames(userName, password, endPoint string) {
	ha := NewHaproxy(userName, password, endPoint)
	cookie, err := ha.GetCookie()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	ha.ReqActions.Clear()
	resp, err := ha.GetHaproxyList(cookie)
	if err != nil {
		fmt.Println("get haproxylist err :", err)
		return
	}

	fmt.Println(resp.FrontNames)
	return
}
