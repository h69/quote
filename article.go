package main

import (
	"fmt"
	"log"
	"strings"
)

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
	article.Cover = cover

	stockCloseTime := GetStockCloseTime()
	stockMarketIndex := GetStockMarketIndex()
	stockMarketOverview := GetStockMarketOverview()

	// 摘要
	var down, flat, up float64
	for _, bar := range stockMarketOverview {
		if bar.Flag == 1 {
			up += bar.Value
		} else if bar.Flag == 0 {
			flat += bar.Value
		} else {
			down += bar.Value
		}
	}
	article.Digest += "上涨 " + Float64ToString(up) + " 家，下跌 " + Float64ToString(down) + " 家，平盘 " + Float64ToString(flat) + " 家。"

	// 标题
	article.Title += IntToString(StringToInt(strings.Split(stockCloseTime, "/")[1])) + "月" + IntToString(StringToInt(strings.Split(stockCloseTime, "/")[2])) + "日"
	for _, stock := range stockMarketIndex {
		if stock.Code == "SH000001" {
			article.Title += "：" + "沪深两市今日热度" + fmt.Sprintf("%.0f", (up/(up+flat+down))*100) + "℃，" + stock.Name
			if stock.Percent == "0.00%" {
				article.Title += "平盘" + "，维持在" + stock.Current + "点"
			} else if strings.Contains(stock.Percent, "+") {
				article.Title += "上涨" + strings.ReplaceAll(stock.Percent, "+", "") + "，收盘在" + stock.Current + "点"
			} else {
				article.Title += "下跌" + strings.ReplaceAll(stock.Percent, "-", "") + "，收盘在" + stock.Current + "点"
			}
			break
		}
	}

	// 内容
	article.Content += RenderHeader()
	article.Content += RenderPlaceholder()
	article.Content += RenderStockChart(stockMarketOverview)
	article.Content += RenderPlaceholder()
	article.Content += RenderSubtitle("大盘指数")
	article.Content += RenderStockTable(stockMarketIndex)
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
	article.Content += RenderSubtitle("行情回顾")
	article.Content += RenderStockEvent(GetStockEvents())
	article.Content += RenderPlaceholder()
	article.Content += RenderFooter()

	log.Println(article.Content)

	return article
}
