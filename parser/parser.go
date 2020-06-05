package parser

import (
	"codewarsKataScraper/langs"
	"codewarsKataScraper/models"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)


var langsAll = []int{langs.C, langs.Cpp, langs.Csharp, langs.GoLang, langs.Java, langs.Javascript, langs.Php, langs.Python}

var langMap = map[int]string{
	langs.C: "i.icon-moon-c-lang",
	langs.Cpp: "i.icon-moon-cplusplus",
	langs.Csharp: "i.icon-moon-csharp",
	langs.Clojure: "i.icon-moon-clojure",
	langs.GoLang: "i.icon-moon-go",
	langs.Java: "i.icon-moon-java",
	langs.Javascript: "i.icon-moon-javascript",
	langs.Php: "i.icon-moon-php",
	langs.Python: "i.icon-moon-python",
}

var KataChan = make(chan *models.Kata, 1)

func GetKatasFromHtmlDoc(document *goquery.Document) *goquery.Selection {
	sels := document.Find("div.list-item.kata")
	sels.Each(ParseKata)
	return sels
}

func ParseKata(i int, sel *goquery.Selection)  {
	// TODO: precompile?
	r := regexp.MustCompile("(\\d) kyu(.*)")
	kata := &models.Kata{}
	for _, l := range langsAll {
		if len(sel.Find(langMap[l]).Nodes) == 1 {
			kata.Langs = append(kata.Langs, l)
		}
	}
	title := sel.Find("div.item-title")
	titleText := title.Text()
	reres := r.FindAllStringSubmatch(titleText, -1)
	kata.Title = reres[0][2]
	var err error
	kata.Kyu, err = strconv.Atoi(reres[0][1])
	if err != nil {
		fmt.Printf("kyu conversion error: %s", err)
	}
	a := title.Find("a")
	if a.Nodes[0].Attr[0].Key == "href" {
		kata.Url = "https://www.codewars.com" + a.Nodes[0].Attr[0].Val
	}
	KataChan <- kata
}

//func ParseKatas(sel *goquery.Selection)  {
//	for i, node := range sel.Each(ParseKata) {
//		ParseKata(node)
//	}
//}