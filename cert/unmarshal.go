package cert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pfsense/cli"
	"regexp"
)

// EditCertResp doc
type EditCertResp struct {
	CSRF string `json:"csrf"`
}

// unmarshalEditCertRespAction dco
var unmarshalEditCertRespAction = cli.NamedAction{
	Name: "pfsense.unmarshalEditCertResp",
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		bodyBuf, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		//fmt.Println("string resp :", string(bodyBuf))

		rel := regexp.MustCompile(`var csrfMagicToken = "(.*)";var csrfMagicName = "__csrf_magic"`)
		ress := rel.FindAllString(string(bodyBuf), -1)
		if len(ress) <= 0 {
			r.Error = fmt.Errorf("未找到 index 页的 Crsf")
			return
		}

		out := EditCertResp{
			CSRF: rel.ReplaceAllString(ress[0], "$1"),
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
