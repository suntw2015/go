package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

type movie struct {
	name        string
	coverImg    string
	url         string
	year        string
	coutry      string
	category    string
	language    string
	subtitle    string
	displayDate string
	format      string
	mesure      string
	time        int
	director    string
	actor       string
	desc        string
}

func f(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(n.Attr)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c)
	}
}

func parseHtml(n *html.Node) []string {
	var result []string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				res := checkUrl(a.Val)
				if len(res) > 1 {
					result = append(result, a.Val)
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, parseHtml(c)...)
	}

	return result
}

func main() {
	rootUrl := "http://www.dytt8.net//html/gndy/dyzz/20170519/54026.html"

	error, body := getHtmlContent(rootUrl)
	if error != nil {
		fmt.Println(error)
		return
	}

	parseDetailHtml(body)

	// parseContent := parseHtml(content)
	// fmt.Println(parseContent)
}

func getHtmlContent(url string) (error, []byte) {
	res, error := http.Get(url)
	if error != nil {
		return error, nil
	}

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return error, nil
	}

	r := bytes.NewReader(body)
	d, _ := charset.NewReader(r, "gb2312")
	content, _ := ioutil.ReadAll(d)

	return nil, content
}

func checkUrl(s string) string {
	re := regexp.MustCompile("/html/gndy/(dyzz|jddy)+[0-9/]+(.html)$")
	return re.FindString(s)
}

func parseDetailHtml(html []byte) {

	re := regexp.MustCompile("(<br\\s/>)(\u25CE)?[\\p{Han}\u3000a-zA-Z0-9\\s\\/+_x&;\uff0c\u3002\uff08\uff09]+")
	s := re.FindAll(html, -1)

	for _, tmp := range s {

	}

	// fmt.Println(s)
}
