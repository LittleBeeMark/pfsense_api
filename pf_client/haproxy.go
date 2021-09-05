package pf_client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

func HaproxyIndex(i int) {

	//cl := Cl{
	//	BaseUrl:     "https://192.168.252.183",
	//	ClientToken: "pfsense",
	//	ClientID:    "admin",
	//}
	//var resp bytes.Buffer
	//
	//err := cl.Request(context.Background(), http.MethodGet, "", "", nil, &resp)
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return
	//}
	//
	//fmt.Println("============================== index  ========================================", resp.String())

	//fmt.Println("resp", resp.String())

	//var test = `var csrfMagicToken = "sid:dc9433ef8f2b7ead2f118ac1a125e4571528f1ab,1630574645;ip:rerasoera";var csrfMagicName = "__csrf_magic";`
	//result2 := rel.ReplaceAllString(test, "$1")
	//fmt.Println("result2 : ", result2)

	//rel := regexp.MustCompile(`var csrfMagicToken = "(.*)";var csrfMagicName = "__csrf_magic"`)
	//ress := rel.FindAllString(resp.String(), -1)
	//result := rel.ReplaceAllString(ress[0], "$1")
	//fmt.Println("ressult : ", result)
	//
	//relCookie := strings.Split(cl.SetCookie, ";")
	//fmt.Println("index cookie", relCookie[0])

	//for i := 0; i < 2; i++ {
	//	HaproxyIndex(i)
	//}
	//coo, err := LoginHaproxy(result, relCookie[0])
	//if err != nil {
	//	fmt.Println("err ", err)
	//	return
	//}

	coo, err := LoginHaproxy("sid:02750c770c4aa60730b0793d367c36fce3603318,1630725723;ip:5a6d0e8dcc6cad6ea23818a0facb76f2681639da,1630725723", "PHPSESSID=0b19060f2979be8f80cfa03aec608163")
	if err != nil {
		fmt.Println("err ", err)
		return
	}

	relloCookie := strings.Split(coo, ";")
	fmt.Println("login cookie", relloCookie[0])
	GetHaproxyList(relloCookie[0])

	//doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Bytes()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//doc.Find(".login").Each(func(i int, s *goquery.Selection) {
	//	content, err := s.Html()
	//	if err != nil {
	//		return
	//	}
	//	fmt.Printf("%d: %s\n", i, content)
	//})
}
func Index() error {
	cl := Cl{
		BaseUrl:     "http://192.168.252.183",
		ClientToken: "pfsense",
		ClientID:    "admin",
	}

	req, err := http.NewRequest(http.MethodGet, cl.BaseUrl, nil)
	if err != nil {
		return err
	}

	resp, err := Cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("respbuf : ", string(respBuf))

	return nil

}

func LoginHaproxy(csrf, cookie string) (string, error) {

	cl := Cl{
		BaseUrl:     "http://192.168.252.183",
		ClientToken: "pfsense",
		ClientID:    "admin",
	}

	//data := url.Values{
	//	"__csrf_magic": []string{csrf},
	//	"usernamefld":  []string{"admin"},
	//	"passwordfld":  []string{"pfsense"},
	//	"login":        []string{"Sign In"},
	//}

	//req.PostForm = url.Values{
	//	"__csrf_magic": []string{csrf},
	//	"usernamefld":  []string{"admin"},
	//	"passwordfld":  []string{"pfsense"},
	//	"login":        []string{"Sign In"},
	//}
	//

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("__csrf_magic", "sid:d6bef3607ccc9f7fadc8b1c62532682f60023c65,1630751389;ip:662905e4b2f984817d2dd0d30b54ca039ce09ba0,1630751389")
	writer.WriteField("usernamefld", cl.ClientID)
	writer.WriteField("passwordfld", cl.ClientToken)
	writer.WriteField("login", "Sign In")

	req, err := http.NewRequest(http.MethodPost, cl.BaseUrl, payload)
	if err != nil {
		return "", err
	}
	fmt.Println(req.PostForm)
	//s := fmt.Sprintf("multipart/form-data;boundary=%s", strconv.Itoa(int(req.ContentLength)))
	//req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Length", strconv.Itoa(int(req.ContentLength)))
	cook := http.Cookie{Name: "PHPSESSID", Value: "308e337bb77786ec087aea9243a16bec"}
	req.AddCookie(&cook)
	//req.Header.Add("PHPSESSID", "5af1e9f665bb25f442f3fb01ec03a503")
	fmt.Println(req.Header)
	resp, err := Cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	cs := GenCookieJar.Cookies(req.URL)
	for _, c := range cs {
		fmt.Println(c.Name, c.Value)
	}
	//payload := &bytes.Buffer{}
	//writer := multipart.NewWriter(payload)
	//writer.WriteField("__csrf_magic", csrf)
	//writer.WriteField("usernamefld", "admin")
	//writer.WriteField("passwordfld", "pfsense")
	//writer.WriteField("login", "Sign In")
	//
	//var resp bytes.Buffer
	//err := cl.Request(context.Background(), http.MethodPost, "", cookie, payload, &resp)
	//if err != nil {
	//	return "", err
	//}

	respBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Header.Get("SetCookie"))

	fmt.Println("err :", err)
	fmt.Println("status code : ", resp.StatusCode)
	fmt.Println("====================== login  ================== ", string(respBuf))

	return resp.Header.Get("Set-Cookie"), nil
}
func GetHaproxyList(cookie string) {
	cl := Cl{
		BaseUrl:     "http://192.168.252.183/haproxy/haproxy_listeners.php",
		ClientToken: "pfsense",
		ClientID:    "admin",
	}
	var resp bytes.Buffer

	err := cl.Request(context.Background(), http.MethodGet, "", cookie, nil, &resp)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println("============================== get haproxy list ======================== : ", resp.String())

}
