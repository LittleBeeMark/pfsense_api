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
	GetHaproxyList("PHPSESSID=5fa1244328b0d2acb5023b95866720d4")

}

func TestLoginHaproxy(t *testing.T) {
	cookie, err := LoginHaproxy("", "")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println("cookie : ", cookie)
}
