package main

import (
	"flag"
	"log"
	"time"

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
		Add(article.Title, article.Digest, article.Content, article.Cover)
	}

	c := cron.New()
	c.AddFunc("0 10 15 * * ?", func() {
		for i := 0; i < 3; i++ {
			if func() bool {
				defer func() {
					if err := recover(); err != nil {
						log.Println(err)
					}
				}()

				stockCloseTime := GetStockCloseTime()

				if stockCloseTime != "" {
					if stockCloseTime == GetDate(0) {
						article := GenerateArticle()
						Add(article.Title, article.Digest, article.Content, article.Cover)
					}
					return true
				}

				return false

			}() {
				break

			} else {
				time.Sleep(3 * time.Second)
			}
		}
	})
	c.Start()

	select {}
}
