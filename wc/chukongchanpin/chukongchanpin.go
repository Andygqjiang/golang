package chukongchanpin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	d "../netDetail"

	"github.com/axgle/mahonia"
)

const (
	divLink = "<a href=\"(.*?)\" target=\"_blank\" data-scode=\"60440\" title=\"(.*?)\">"
	yinhao  = "\""
)

var divLinkRegexp, yinhaoRegexp *regexp.Regexp

func Handler(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
	}

	string_body := string(body)

	dec := mahonia.NewDecoder("gbk")
	utf_body := dec.ConvertString(string_body)

	aLink := divLinkRegexp.FindAllString(utf_body, -1)

	var param d.Param
	for i := 0; i < len(aLink); i++ {
		alinks := yinhaoRegexp.FindAllStringIndex(aLink[i], -1)
		detaiLink := aLink[i][alinks[0][1]:alinks[1][0]]
		param.Url = detaiLink
		d.Handler(&param)
		time.Sleep(time.Second * 6)
	}

}

func init() {
	divLinkRegexp = regexp.MustCompile(divLink)
	yinhaoRegexp = regexp.MustCompile(yinhao)
}
