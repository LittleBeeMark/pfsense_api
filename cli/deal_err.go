package cli

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	EditCertCsrfMissOrExpired = "Missing or expired CSRF token"
	FollowingError            = "The following input errors were detected"
	UserPasswordError         = "Username or Password incorrect"
)

func getInputErr(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	return doc.Find(".input-errors").Find("p~ul").Text(), nil
}

// UnmarshalErrPageBasicAction doc
var UnmarshalErrPageBasicAction = NamedAction{
	Name: UnmarshalErrPageBasic,
	Fn: func(r *Request) {
		if strings.Contains(string(r.RespBody), UserPasswordError) {
			r.Error = errors.New("登录用户或密码错误")
			return
		}

		if strings.Contains(string(r.RespBody), FollowingError) {
			errStr, err := getInputErr(string(r.RespBody))
			if err != nil {
				r.Error = fmt.Errorf("get input err: %v", err)
				return
			}

			if errStr != "" {
				r.Error = errors.New(errStr)
				return
			}

			r.Error = errors.New("input params error please check your param ")
			return
		}

		if strings.Contains(string(r.RespBody), EditCertCsrfMissOrExpired) {
			r.Error = errors.New("登录过期，请重新尝试：Missing or expired CSRF token ")
			return
		}
	},
}
