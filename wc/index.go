package main

import (
	"fmt"
	"os"
	"strconv"

	ckp "./chukongchanpin"
)

func main() {

	len_args := len(os.Args)

	if len_args != 4 {
		return
	}

	args_url := os.Args[1]
	args_startNumber := os.Args[2]
	args_endNumber := os.Args[3]

	int_startNumber, _ := strconv.Atoi(args_startNumber)
	int_endNumber, _ := strconv.Atoi(args_endNumber)

	var url string
	for i := int_startNumber; i <= int_endNumber; i++ {
		url = args_url + "?p=" + strconv.Itoa(i)
		ckp.Handler(url)
	}

}
