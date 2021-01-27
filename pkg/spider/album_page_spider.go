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
	"github.com/picspider/pkg/util"
)

// PhotoAlbumPageSpider is photo album page spider
func PhotoAlbumPageSpider(dirname, band string) {
	photoAlbumID := models.GetPhotoAlbumID(band)

	urlstr := band
	u, err := url.Parse(urlstr)
	if err != nil {
		go PhotoAlbumPageSpider(dirname, band)
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

		count := htmlDoc.Find(".content-hidden").Length()
		if count > 0 {
			log.Println("需要支付")
			return
		}

		// 找到抓取项 <div class="entry-content"> 下所有的a解析
		htmlDoc.Find(".entry-content p img").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Attr("src")
			title, _ := s.Attr("title")
			fmt.Printf("图片 %d: %s - %s\n", i+1, title, band)
			models.SavePhotoResult(models.Photo{PhotoName: title, PhotoURL: band, PhotoAlbumID: photoAlbumID})
			// save photo to local
			util.DownloadPhoto(dirname, title, band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})

	err = c.Visit(urlstr)
}
