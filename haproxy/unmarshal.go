package haproxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"pfsense/cli"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// unmarshalHaproxyListRespAction dco
var unmarshalHaproxyListRespAction = cli.NamedAction{
	Name: "pfsense.unmarshalHaproxyListResp",
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()
		bodyBuf, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		//fmt.Println("haproxy string resp ==================================:", string(bodyBuf))
		names := QueryHaproxyNamesOfHtml(string(bodyBuf))
		fmt.Println("names ::::;", names)
		//outRaw, err := json.Marshal(out)
		//if err != nil {
		//	r.Error = fmt.Errorf("marshal login output info err : %v", err)
		//	return
		//}
		//
		//err = json.Unmarshal(outRaw, &r.Data)
		//if err != nil {
		//	r.Error = fmt.Errorf("unmarshal login output info to data err : %v", err)
		//	return
		//}
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
	//addressCount := 6
	//currentName := ""
	dom.Find("tbody[class=user-entries]").Children().Children().Each(func(i int, selection *goquery.Selection) {
		if selection.Text() != "" {
			if i == nameCount {
				haproxyNames = append(haproxyNames, selection.Text())
				//currentName = selection.Text()
				nameCount += 10
			}

			//if i == addressCount {
			//	n := fmt.Sprintf("%s(%s)", currentName, selection.Text())
			//	haproxyNames = append(haproxyNames, n)
			//	addressCount += 10
			//}

			//fmt.Printf("i : %d  :  text : %+v", i, selection.Text())
		}

	})
	return haproxyNames
}
