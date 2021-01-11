package main

import "log"

// curl -F media=@cover.png "https://api.weixin.qq.com/cgi-bin/material/add_material?type=thumb&access_token="
const cover = "47AXncnDO_HCnVfHX6JDl07Rb3Z6LmXnu90dwYcuEag"

// Article 文章
type Article struct {
	Title   string
	Digest  string
	Content string
	Cover   string
}

// GenerateArticle 生成文章
func GenerateArticle() Article {
	var article Article
	article.Title = GetStockCloseTime()
	article.Digest = "A 股行情报告。"
	article.Cover = cover

	article.Content += RenderHeader()
	article.Content += RenderPlaceholder()
	article.Content += RenderStockChart(GetStockMarketOverview())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("大盘指数")
	article.Content += RenderStockTable(GetStockMarketIndex())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("热门榜")
	article.Content += RenderStockTable(GetStockFollow())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("市值榜")
	article.Content += RenderStockTable(GetStockMarketCapital())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("价格榜")
	article.Content += RenderStockTable(GetStockCurrent())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("成交量榜")
	article.Content += RenderStockTable(GetStockVolume())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("成交额榜")
	article.Content += RenderStockTable(GetStockAmount())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("年涨幅榜")
	article.Content += RenderStockTable(GetStockCurrentYearPercent())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("日涨幅榜")
	article.Content += RenderStockTable(GetStockPercent())
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("招财大牛猫茅 20 组合")
	article.Content += RenderStockTable(GetStockMao20())
	article.Content += RenderPlaceholder()
	article.Content += RenderFooter()

	log.Println(article.Content)

	return article
}
