package cli

import (
	"errors"
	"fmt"
)

var (
	UserName string = "admin"
	Password string = "pfsense"
	EndPoint string = "https://192.168.252.183"
)

// Index doc
func Index() (map[string]interface{}, error) {
	ha := NewPfsense(UserName, Password, EndPoint)
	output, err := ha.GetIndexReq()
	if err != nil {
		return nil, err
	}

	fmt.Println("output:", output)
	return *output, nil
}

// Login doc
func Login(crsf, cookie string) (map[string]interface{}, error) {
	pf := NewPfsense(UserName, Password, EndPoint)
	output, err := pf.Login(crsf, cookie)
	if err != nil {
		return nil, err
	}

	return *output, nil
}

// GetLoginCookie doc
func GetLoginCookie() (string, error) {
	output, err := Index()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	var (
		csrf   string
		cookie string
	)
	if v, ok := output[RespInfoCSRF]; ok {
		csrf = v.(string)
	}

	if v, ok := output[RespInfoSetCookie]; ok {
		cookie = v.(string)
	}

	op, err := Login(csrf, cookie)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	if v, ok := op[RespInfoSetCookie]; ok {
		return v.(string), nil
	}

	return "", errors.New("未获取到 login cookie")
}
