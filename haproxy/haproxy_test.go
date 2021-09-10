package haproxy

import "testing"

func TestHaproxy_GetHaproxyList(t *testing.T) {
	GetHaproxyFrontNames(UserName, Password, EndPoint)
}
