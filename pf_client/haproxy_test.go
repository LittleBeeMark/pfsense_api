package pf_client

import (
	"fmt"
	"testing"
)

func TestHaproxyIndex(t *testing.T) {
	HaproxyIndex(0)
}

func TestIndex(t *testing.T) {
	err := Index()
	if err != nil {
		fmt.Println("err :", err)
		return
	}
}

func TestGetHaproxyList(t *testing.T) {
	GetHaproxyList("PHPSESSID=08bb921a358242308f6c21a6d7d5b13c")

}

func TestLoginHaproxy(t *testing.T) {
	cookie, err := LoginHaproxy("", "")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println("cookie : ", cookie)
}
