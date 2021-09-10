package cert

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"pfsense/cli"
	"strings"
)

// EditCertResp doc
type EditCertResp struct {
	CSRF  string `json:"csrf"`
	Descr string `json:"descr"`
}

func getDescriptionName(html string) (string, error) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	if val, ok := reader.Find("input#descr").Attr("value"); ok {
		fmt.Println("ok : ", val, ok)
		return val, nil
	}

	return "", errors.New("获取证书名称失败！")
}

// unmarshalEditCertRespAction dco
var unmarshalEditCertRespAction = cli.NamedAction{
	Name: cli.UnmarshalEditCertPage,
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		var err error
		r.RespBody, err = ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		//fmt.Println("string resp :", string(r.RespBody))

		csrf, err := cli.GetCsrfInfo(string(r.RespBody))
		if err != nil {
			r.Error = err
			return
		}

		descr, err := getDescriptionName(string(r.RespBody))
		if err != nil {
			r.Error = fmt.Errorf("未找到证书编辑页的证书名称(您的选择的更新证书可能已经不存在于 Pfsense 请检查): %v", err)
			return
		}

		out := EditCertResp{
			CSRF:  csrf,
			Descr: descr,
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
