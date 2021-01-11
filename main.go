package main

import (
	"flag"
	"log"

	"github.com/robfig/cron"
)

func main() {
	var n int
	flag.StringVar(&AppID, "id", "", "")
	flag.StringVar(&AppSecret, "secret", "", "")
	flag.IntVar(&n, "n", 0, "")
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if n == 1 {
		article := GenerateArticle()
		AddNews(article.Title, article.Digest, article.Content, article.Cover)
	}

	c := cron.New()
	c.AddFunc("0 10 15 * * ?", func() {
		if GetStockCloseTime() == GetDate(0) {
			article := GenerateArticle()
			AddNews(article.Title, article.Digest, article.Content, article.Cover)
		}
	})
	c.Start()

	select {}
}
