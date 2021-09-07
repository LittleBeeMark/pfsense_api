package haproxy

import (
	"fmt"
)

var (
	UserName string = "admin"
	Password string = "pfsense"
	EndPoint string = "https://192.168.252.183"
)

func HaproxyList(cookie string) (map[string]interface{}, error) {
	ha := NewHaproxy(UserName, Password, EndPoint)

	fmt.Println("login cookie ===================", cookie)
	op, err := ha.GetHaproxyList(cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return nil, err
	}

	return *op, nil
}

func HaproxyInfo(cookie, name string) (map[string]interface{}, error) {
	ha := NewHaproxy(UserName, Password, EndPoint)
	ha.GetHaproxyInfo(cookie, name)
	return nil, nil
}
