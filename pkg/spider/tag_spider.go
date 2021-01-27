package spider

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// TagSpider search page spider
func TagSpider() {
	urlstr := "https:///tags"
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Println("Tag request error")
	}
	c := colly.NewCollector()
	// timeout
	c.SetRequestTimeout(1000 * time.Second)
	// Agent
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		// Request header
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", urlstr)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("title:", e.Text)
	})

	c.OnResponse(func(resp *colly.Response) {
		fmt.Println("response received", resp.StatusCode)

		// goquery read resp.Body
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))

		// url send goquery，get the url respon，not recommended here
		// htmlDoc, err := goquery.NewDocument(resp.Request.URL.String())

		if err != nil {
			log.Fatal(err)
		}

		// 找到抓取项 <div class="post-info"> 下所有的a解析
		// #main > ul > li:nth-child(1) > a > h2
		htmlDoc.Find("main ul li a").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Attr("href")
			tag := s.Find("h2").Text()
			fmt.Printf("标签 %d: %s - %s\n", i+1, tag, band)
			// visit tag info page
			TagBand(band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})

	err = c.Visit(urlstr)
}
