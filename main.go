package main

import (
	"codewarsKataScraper/http"
	"codewarsKataScraper/langs"
	"codewarsKataScraper/models"
	"codewarsKataScraper/parser"
	"fmt"
	"time"
)

const URL  = "https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&beta=false"
var urls = []string{
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&beta=false",
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&beta=false&order_by=published_at+asc",
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&xids=played&beta=false&order_by=published_at+asc",
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&xids=played&beta=false&order_by=total_completed+desc",
	"https://www.codewars.com/kata/search/go?q=num&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&xids=played&beta=false&order_by=total_completed+desc",
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&xids=played&beta=false&order_by=popularity+desc",
	"https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&r%5B%5D=-7&r%5B%5D=-6&r%5B%5D=-5&r%5B%5D=-4&xids=played&beta=false&order_by=published_at+desc",
}

func KatasToPool(pool *models.KataPool)  {
	var k *models.Kata
	for {
		k = <- parser.KataChan
		if exist := pool.GetKataByTitle(k.Title); exist == nil {
			pool.Katas = append(pool.Katas, k)
		} else {
			fmt.Println(k.Title, " already in pool")
		}
	}
}

func main() {
	for _, url := range urls {
		doc, err := http.FetchHtmlDoc(url)
		go parser.GetKatasFromHtmlDoc(doc)
		if err != nil {
			fmt.Println("well, fuck")
		}
	}
	kataPool := &models.KataPool{}
	go KatasToPool(kataPool)
	time.Sleep(50 * time.Second)
	fmt.Println("Kata pool size: ", len(kataPool.Katas))
	fmt.Println("5 kyue: ", kataPool.FindKatasByKyu(5))
	req := kataPool.FindKatasByLangList([]int{langs.Csharp, langs.Python, langs.Javascript, langs.Java, langs.Php, langs.GoLang})
	fmt.Println(len(req))
	for _, r := range req {
		fmt.Println(r.Title)
		fmt.Println(r.Url)
	}
}