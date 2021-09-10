package haproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"pfsense/cli"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// unmarshalHaproxyListRespAction dco
var unmarshalHaproxyListRespAction = cli.NamedAction{
	Name: cli.UnmarshalHaproxyList,
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}

		// names := QueryHaproxyNamesOfHtml(string(bodyBuf))
		csrf, err := cli.GetCsrfInfo(string(r.RespBody))
		out := ListResp{
			CSRF: csrf,
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

// unmarshalHaproxyListRespAction dco
var unmarshalHaproxyListNameRespAction = cli.NamedAction{
	Name: cli.UnmarshalHaproxyNameList,
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}

		names := QueryHaproxyNamesOfHtml(string(r.RespBody))
		csrf, err := cli.GetCsrfInfo(string(r.RespBody))
		out := ListResp{
			CSRF:       csrf,
			FrontNames: names,
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

// QueryHaproxyNamesOfHtml doc
func QueryHaproxyNamesOfHtml(html string) []string {
	var haproxyNames []string
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	nameCount := 4
	dom.Find("tbody[class=user-entries]").Children().Children().Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" {
			if i == nameCount {
				haproxyNames = append(haproxyNames, selection.Text())
				nameCount += 10
			}
		}
	})
	return haproxyNames
}
