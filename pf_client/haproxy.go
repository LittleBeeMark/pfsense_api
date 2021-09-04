package pf_client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

	coo, err := LoginHaproxy("sid:7c1df9a116c37f4235fae322edf31851f8a30938,1630676016;ip:d875c020b7ec055f979eaee60729c364ee9f8215,1630676016", "PHPSESSID=31683419d65650ddf6fc217a4c2a8687")
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

func LoginHaproxy(csrf, cookie string) (string, error) {

	cl := Cl{
		BaseUrl:     "https://192.168.252.183",
		ClientToken: "pfsense",
		ClientID:    "admin",
	}

	//data := url.Values{
	//	"__csrf_magic": []string{csrf},
	//	"usernamefld":  []string{"admin"},
	//	"passwordfld":  []string{"pfsense"},
	//	"login":        []string{"Sign In"},
	//}

	req, err := http.NewRequest("POST", cl.BaseUrl, nil)
	if err != nil {
		return "", err
	}
	req.PostForm = url.Values{
		"__csrf_magic": []string{csrf},
		"usernamefld":  []string{"admin"},
		"passwordfld":  []string{"pfsense"},
		"login":        []string{"Sign In"},
	}

	fmt.Println(req.PostForm)
	req.Header.Set("Cookie", cookie)
	//s := fmt.Sprintf("multipart/form-data;boundary=%s", strconv.Itoa(int(req.ContentLength)))
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Length", strconv.Itoa(len(req.PostForm.Encode())))
	fmt.Println(req.Header)
	resp, err := Cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

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

	fmt.Println("err :", err)
	fmt.Println("status code : ", resp.StatusCode)
	fmt.Println("====================== login  ================== ", string(respBuf))

	return resp.Header.Get("Set-Cookie"), nil
}
func GetHaproxyList(cookie string) {
	cl := Cl{
		BaseUrl:     "https://192.168.252.183/haproxy/haproxy_listeners.php",
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
