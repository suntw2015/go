package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "/html/gndy/dyzz/20170519/54026.html"
	p := "/html/zongyi2013/daluzongyi/20170515/53993.html"

	re := regexp.MustCompile("/html/gndy/(dyzz|jddy)+[0-9/]+(.html)$")
	fmt.Printf("%q\n", re.FindString(s))
	fmt.Printf("%q\n", re.FindString(p))
}
