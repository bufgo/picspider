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
	"github.com/picspider/models"
)

var flag = false

// TagBand is get page count
func TagBand(band string) {
	tmp := band
	var i = 1
	flag = false
	for {
		url := tmp
		url = band + "/page/" + fmt.Sprint(i)
		fmt.Println(url)
		TagInfoPageSpider(url)
		i++
		if flag {
			break
		}
	}
}

// TagInfoPageSpider tag info page spider
func TagInfoPageSpider(page string) {
	urlstr := page
	u, err := url.Parse(urlstr)
	if err != nil {
		go TagInfoPageSpider(page)
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

		//#post-list > div
		htmlDoc.Find(".empty-page").Each(func(i int, s *goquery.Selection) {
			flag = true
			return
		})

		// 找到抓取项 <div class="post-info"> 下所有的a解析
		htmlDoc.Find(".post-info h2 a").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Attr("href")
			title := s.Text()
			fmt.Printf("图集 %d: %s - %s\n", i+1, title, band)
			models.SaveSearchResult(models.PhotoAlbum{AlbumName: title, AlbumURL: band})
			// visit album info page
			PhotoAlbumPageSpider(title, band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})

	err = c.Visit(urlstr)
}
