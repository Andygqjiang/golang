package netDetail

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/axgle/mahonia"
)

const (
	phone        = "<div class=\"personal_bottom\"><span>([0-9]*?)</span></div>"
	name         = "<div class=\"personal_top\"><div class=\"t\"><span>(.*?)</span></div> </div>"
	companyName  = "<span class=\"corpname\">(.*?)</span>"
	prePhone     = "<div class=\"personal_bottom\"><span>"
	preName      = "<div class=\"personal_top\"><div class=\"t\"><span>"
	preCompany   = "<span class=\"corpname\">"
	stuffPhone   = "</span></div>"
	stuffName    = "</span></div> </div>"
	stuffCompany = "</span>"
	filePath     = "d:\\company.txt"
)

type Param struct {
	Url string
}

var phoneRegexp, nameRegexp, companyRegexp *regexp.Regexp

func Handler(param *Param) {
	resp, err := http.Get(param.Url)
	if err != nil {
		fmt.Println(err)
	}
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer resp.Body.Close()
	string_body := string(body)

	dec := mahonia.NewDecoder("gbk")

	utfBody := dec.ConvertString(string_body)
	phoneDiv := phoneRegexp.FindString(utfBody)
	fmt.Println("----------------------")
	if phoneDiv != "" {
		str_phone := phoneDiv[len(prePhone) : len(phoneDiv)-len(stuffPhone)]
		fmt.Println(str_phone)

		nameDiv := nameRegexp.FindString(utfBody)
		str_name := nameDiv[len(preName) : len(nameDiv)-len(stuffName)]
		fmt.Println(str_name)

		companyDiv := companyRegexp.FindString(utfBody)
		str_company := companyDiv[len(preCompany) : len(companyDiv)-len(stuffCompany)]
		fmt.Println(str_company)

		address := getNum()
		writeToFile(str_company + "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t" + str_phone + "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t" + str_name + "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t" + strconv.Itoa(address) + "\n")
	}

}

func writeToFile(content string) {

	var tmp_file *os.File
	if exist() {
		tmp_file, _ = os.Create(filePath)
		tmp_file.Chmod(os.ModeAppend)
	}
	tmp_file, _ = os.OpenFile(filePath, os.O_APPEND, os.ModeAppend)
	defer tmp_file.Close()
	tmp_file.WriteString(content)
}

func exist() bool {
	_, err := os.Stat(filePath)
	return os.IsNotExist(err)
}

func getNum() int {
	return rand.Intn(100)
}

func init() {

	phoneRegexp = regexp.MustCompile(phone)

	nameRegexp = regexp.MustCompile(name)

	companyRegexp = regexp.MustCompile(companyName)

}
