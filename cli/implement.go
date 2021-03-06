package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
)

// ========================================= Send ======================================

// SendAction doc
var SendAction = NamedAction{
	Name: SendBasic,
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

// ================================================== unmarshal  ===================================

// PfSenseAPIResp pfSense API 接口返出结构
type PfSenseAPIResp struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Return  int         `json:"return"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

var UnmarshalAPIBasicAction = NamedAction{
	Name: UnmarshalAPIBasic,
	Fn: func(r *Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}

		resData := PfSenseAPIResp{
			Data: r.Data,
		}

		err = json.Unmarshal(r.RespBody, &resData)
		if err != nil {
			r.Error = fmt.Errorf("unmarshal API  output(%s) info to data err : %v",
				string(r.RespBody), err)
			return
		}
	},
}

// UnmarshalBasicAction doc
var UnmarshalBasicAction = NamedAction{
	Name: UnmarshalBasic,
	Fn: func(r *Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}

		//fmt.Println("basic body : ", string(r.RespBody))
	},
}

// GetCsrfInfo doc
func GetCsrfInfo(html string) (string, error) {
	rel := regexp.MustCompile(`var csrfMagicToken = "(.*)";var csrfMagicName = "__csrf_magic"`)
	ress := rel.FindAllString(html, -1)
	if len(ress) <= 0 {
		return "", errors.New("未找到 index 页的 Crsf")
	}

	return rel.ReplaceAllString(ress[0], "$1"), nil
}

// unmarshalIndexRespAction dco
var unmarshalIndexRespAction = NamedAction{
	Name: UnmarshalIndex,
	Fn: func(r *Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		//fmt.Println("string resp :", string(bodyBuf))
		csrf, err := GetCsrfInfo(string(r.RespBody))
		if err != nil {
			r.Error = err
			return
		}

		setCookie := r.HTTPResponse.Header.Get("Set-Cookie")
		if setCookie == "" {
			r.Error = fmt.Errorf("未找到 index 页的 Set-Cookie")
			return
		}

		out := IndexResp{
			csrf,
			setCookie,
		}

		outRaw, err := json.Marshal(out)
		if err != nil {
			r.Error = fmt.Errorf("marshal index output info err : %v", err)
			return
		}

		err = json.Unmarshal(outRaw, &r.Data)
		if err != nil {
			r.Error = fmt.Errorf("unmarshal index output info to data err : %v", err)
			return
		}
	},
}

// unmarshalLoginRespAction dco
var unmarshalLoginRespAction = NamedAction{
	Name: UnmarshalLogin,
	Fn: func(r *Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}

		setCookie := r.HTTPResponse.Header.Get("Set-Cookie")
		if setCookie == "" {
			r.Error = fmt.Errorf("未找到登录成功页的 Set-Cookie")
			return
		}

		out := LoginResp{
			setCookie,
		}

		outRaw, err := json.Marshal(out)
		if err != nil {
			r.Error = fmt.Errorf("marshal login output info err : %v", err)
			return
		}

		err = json.Unmarshal(outRaw, &r.Data)
		if err != nil {
			r.Error = fmt.Errorf("unmarshal login output info to data err : %v", err)
			return
		}
	},
}
