package main

import (
	"github.com/picspider/models"
	"github.com/picspider/pkg/setting"
	"github.com/picspider/pkg/spider"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	spider.Spider()
}
