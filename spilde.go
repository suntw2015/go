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
	name      string
	othername string
	coverImg  string
	url       string
	year      string
	country   string
	category  string
	language  string
	subtitle  string
	showDate  string
	fileType  string
	mesure    string
	duration  string
	director  string
	actor     string
	desc      string
}

func (m *movie) setKeyValue(key string, value string) {
	switch key {
	case "name":
		m.name = value
	case "othername":
		m.othername = value
	case "year":
		m.year = value
	case "country":
		m.country = value
	case "category":
		m.category = value
	case "language":
		m.language = value
	case "showDate":
		m.showDate = value
	case "fileType":
		m.fileType = value
	case "mesure":
		m.mesure = value
	case "duration":
		m.duration = value
	case "director":
		m.director += value
	case "actor":
		m.actor += value
	case "desc":
		m.desc += value
	default:
	}
}

func (m *movie) printMovie() {
	fmt.Printf("%s:%s\n", "name", m.name)
	fmt.Printf("%s:%s\n", "othername", m.othername)
	fmt.Printf("%s:%s\n", "year", m.year)
	fmt.Printf("%s:%s\n", "country", m.country)
	fmt.Printf("%s:%s\n", "category", m.category)
	fmt.Printf("%s:%s\n", "language", m.language)
	fmt.Printf("%s:%s\n", "showDate", m.showDate)
	fmt.Printf("%s:%s\n", "fileType", m.fileType)
	fmt.Printf("%s:%s\n", "mesure", m.mesure)
	fmt.Printf("%s:%s\n", "duration", m.duration)
	fmt.Printf("%s:%s\n", "director", m.director)
	fmt.Printf("%s:%s\n", "actor", m.actor)
	fmt.Printf("%s:%s\n", "desc", m.desc)
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
	rootUrl := "http://www.dytt8.net/html/gndy/dyzz/20170514/53986.html"

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

	re := regexp.MustCompile("(<br\\s/>)(\u25CE)?[\\p{Han}\u3000a-zA-Z0-9\\s\\/+_x()&;\uff0c\u3002\uff08\uff09\uff1a\uff01]+")
	s := re.FindAll(html, -1)

	var fieldStart = "\u25CE"
	var movieMapKey = map[string]string{
		"name":      "\u25CE\u7247[\u3000\\s]*\u540d",
		"othername": "\u25CE\u8BD1[\u3000\\s]*\u540D",
		"year":      "\u25CE\u5e74[\u3000\\s]*\u4ee3",
		"country":   "\u25CE(\u56fd|\u4ea7)[\u3000\\s]*(\u5730|\u5bb6)",
		"category":  "\u25CE\u7c7b[\u3000\\s]*\u522b",
		"language":  "\u25CE\u8bed[\u3000\\s]*\u8a00",
		"showDate":  "\u25CE\u4e0a\u6620\u65e5\u671f",
		"fileType":  "\u25CE\u6587\u4EF6\u683C\u5F0F",
		"mesure":    "\u25CE\u89C6\u9891\u5C3A\u5BF8",
		"duration":  "\u25CE\u7247[\u3000\\s]*\u957F",
		"director":  "\u25CE\u5BFC[\u3000\\s]*\u6F14",
		"actor":     "\u25CE\u4E3B[\u3000\\s]*\u6F14",
		"desc":      "\u25CE\u7B80[\u3000\\s]*\u4ECB",
	}

	var m movie
	var lastMapKey string

	re = regexp.MustCompile("(<br\\s/>)")
	fieldStartRe := regexp.MustCompile(fieldStart)
	for _, text := range s {
		tmp := re.ReplaceAll(text, []byte(""))
		if fieldStartRe.Match(tmp) {
			lastMapKey = ""
			for key, value := range movieMapKey {
				tmpRe := regexp.MustCompile(value)
				if tmpRe.Match(tmp) {
					lastMapKey = key
					tmp = tmpRe.ReplaceAll(tmp, []byte(""))
					m.setKeyValue(key, string(tmp))
					break
				}
			}
		} else {
			m.setKeyValue(lastMapKey, string(tmp))
		}
	}

	m.printMovie()
}
