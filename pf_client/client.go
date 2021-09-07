package pf_client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)

// Cl doc
type Cl struct {
	BaseUrl     string
	ClientID    string
	ClientToken string
	SetCookie   string
}

var GenCookieJar *cookiejar.Jar
var Cli *http.Client

func init() {
	GenCookieJar, _ = cookiejar.New(nil)
	Cli = &http.Client{
		Jar: GenCookieJar,
		Transport: &http.Transport{
			//Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		}}
}

const (
	DefaultTimeout = 60 * time.Second
	AuthHeader     = "X-CertCloud-Auth-Key"
)

// NewRequestWithContext doc
func NewRequestWithContext(ctx context.Context, method, url string,
	body io.Reader) (req *http.Request, cancel context.CancelFunc, err error) {

	// 如果没有超时，设置超时
	if t, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, DefaultTimeout)
	} else {
		ctx, cancel = context.WithDeadline(ctx, t)
	}

	req, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return
	}

	return
}

// cmcResp certCloud返出结构
type cmcResp struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Return  int         `json:"return"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// Request 适配请求
func (prov *Cl) Request(ctx context.Context,
	method, endpoint, cookie string, body io.Reader, resp interface{}) (err error) {
	if prov.BaseUrl == "" {
		return errors.New("未配置CertCloud BaseURL,请配置。")
	}

	if prov.ClientToken == "" {
		return errors.New("未配置CertCloud CMCTPLAccessKey,请配置。")
	}

	req, cancel, err := NewRequestWithContext(ctx, method, prov.BaseUrl+endpoint, body)
	if err != nil {
		return
	}
	defer cancel()

	// 构造头部
	if cookie != "" {
		req.Header.Add("Cookie", cookie)
	}

	//req.Header.Add(AuthHeader, prov.ClientToken)
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Content-Length", strconv.Itoa(int(req.ContentLength)))

	fmt.Println("======================= header :", req.Header)
	fmt.Println("======================= body :", req.Body)
	doResp, err := Cli.Do(req)
	if err != nil {
		return
	}

	fmt.Println("status code", doResp.StatusCode)
	setcookie := doResp.Header.Get("Set-Cookie")
	//fmt.Println("cookie :", setcookie)
	prov.SetCookie = setcookie

	var respBuf []byte
	w, ok := resp.(*bytes.Buffer)
	if ok {
		io.Copy(w, doResp.Body)
		respBuf = w.Bytes()
		return nil
	}

	// cmcResp 数据获取
	respBuf, err = ioutil.ReadAll(doResp.Body)
	if err != nil {
		return
	}

	// 打印诊断日志
	//	defer func() {
	//		var secKey string = "nil"
	//
	//		var bodyStr string = "nil"
	//		if req.GetBody != nil {
	//			if bodyReader, err := req.GetBody(); err == nil {
	//				if b, err := ioutil.ReadAll(bodyReader); err == nil {
	//					bodyStr = string(b)
	//				}
	//			}
	//		}
	//
	//		var errStr string = "nil"
	//		if err != nil {
	//			errStr = err.Error()
	//		}
	//
	//		str := `
	//New Req To CertManager
	//	URL: ` + prov.BaseUrl + endpoint + `
	//	Method: ` + method + `
	//	AccessKey: ` + secKey + `
	//	Body: ` + bodyStr + `
	//	Resp:` + string(respBuf) + `
	//	Error: ` + errStr
	//
	//		fmt.Println(prov, str)
	//	}()
	//
	result := &cmcResp{
		Data: resp,
	}

	fmt.Println(string(respBuf))
	//fmt.Println("respBuf ", string(respBuf))
	err = json.Unmarshal(respBuf, result)
	if err != nil {
		// 直接返出返回的数据
		return fmt.Errorf("pfsense 返回数据（%s）有误!, url: (%s), err: (%s)",
			string(respBuf), prov.BaseUrl+endpoint, err)
	}

	// 错误处理
	fmt.Println(doResp.StatusCode)
	if doResp.StatusCode/100 == 4 || result.Code != 200 {
		return fmt.Errorf("pfsense API错误: code:%d, message:%s", result.Code, result.Error)
	}

	// 网络错误
	if doResp.StatusCode/100 != 2 {
		return fmt.Errorf("网络连接错误: code:%d, message:%s", doResp.StatusCode, string(respBuf))
	}

	return
}
