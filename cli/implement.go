package cli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// SendAction dco
var SendAction = NamedAction{
	Name: "pfsense.SendAction",
	Fn: func(r *Request) {
		sender := sendFollowRedirects
		if r.DisableFollowRedirects {
			sender = sendWithoutFollowRedirects
		}
		var err error
		r.HTTPResponse, err = sender(r)
		if err != nil {
			handleSendError(r, err)
		}
	},
}

func sendFollowRedirects(r *Request) (*http.Response, error) {
	return r.ClientInfo.HTTPClient.Do(r.HTTPRequest)
}

func sendWithoutFollowRedirects(r *Request) (*http.Response, error) {
	r.ClientInfo.Crd.CookieJar, _ = cookiejar.New(nil)
	r.ClientInfo.HTTPClient.Jar = r.ClientInfo.Crd.CookieJar
	return r.ClientInfo.HTTPClient.Do(r.HTTPRequest)
}

func handleSendError(r *Request, err error) {
	if r.HTTPResponse != nil {
		defer r.HTTPResponse.Body.Close()
	}

	// cmcResp 数据获取
	respBuf, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = fmt.Errorf("读取 HTTPResponse.Body 数据失败: %v", err)
		return
	}

	// 错误处理
	fmt.Println(r.HTTPResponse.StatusCode)
	if r.HTTPResponse.StatusCode/100 == 4 {
		r.Error = fmt.Errorf("pfsense API错误: code:%d, message:%s", r.HTTPResponse.StatusCode, err)
	}

	// 网络错误
	if r.HTTPResponse.StatusCode/100 != 2 {
		r.Error = fmt.Errorf("网络连接错误: code:%d, message:%s", r.HTTPResponse.StatusCode, string(respBuf))
	}

}

var windowBuilAction = NamedAction{
	Name: "pfsense.windowBuilAction",
	Fn: func(r *Request) {
		sender := sendFollowRedirects
		if r.DisableFollowRedirects {
			sender = sendWithoutFollowRedirects
		}
		var err error
		r.HTTPResponse, err = sender(r)
		if err != nil {
			handleSendError(r, err)
		}
	},
}

var apiBuilAction = NamedAction{
	Name: "pfsense.windowBuilAction",
	Fn: func(r *Request) {
		sender := sendFollowRedirects
		if r.DisableFollowRedirects {
			sender = sendWithoutFollowRedirects
		}
		var err error
		r.HTTPResponse, err = sender(r)
		if err != nil {
			handleSendError(r, err)
		}
	},
}
