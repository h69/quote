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
	stockCloseTime := GetStockCloseTime()
	stockMarketOverview := GetStockMarketOverview()
	stockMarketForeign := GetStockMarketForeign()
	stockIndustry := GetStockIndustry()
	stockComment := GetStockComment()
	stockMarketIndex := GetStockMarketIndex()
	stockFollow := GetStockFollow()
	stockMao20 := GetStockMao20()
	stockPercent := GetStockPercent()
	stockCurrentYearPercent := GetStockCurrentYearPercent()
	stockCurrent := GetStockCurrent()
	stockMarketCapital := GetStockMarketCapital()
	stockVolume := GetStockVolume()
	stockAmount := GetStockAmount()
	stockForeign := GetStockForeign()
	stockEvent := GetStockEvent()
	stockChance := GetStockChance()

	// 封面
	var article Article
	article.Cover = cover

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
	article.Content += RenderHeader("👆点击关注，领取你的行情精灵")
	article.Content += RenderPlaceholder()
	if len(stockMarketOverview) > 0 {
		article.Content += RenderStockChart(stockMarketOverview)
		article.Content += RenderStockThermometer(stockMarketOverview)
	}
	if len(stockMarketForeign) > 0 {
		article.Content += RenderStockCard(stockMarketForeign)
	}
	if len(stockIndustry) > 0 {
		article.Content += RenderStockPlate(stockIndustry)
	}
	article.Content += RenderPlaceholder()
	if len(stockComment) > 0 {
		article.Content += RenderContent(stockComment)
		article.Content += RenderPlaceholder()
		article.Content += RenderPlaceholder()
	}
	if len(stockMarketIndex) > 0 {
		article.Content += RenderSubtitle("大盘指数")
		article.Content += RenderStockTable(stockMarketIndex)
		article.Content += RenderPlaceholder()
	}
	if len(stockFollow) > 0 {
		article.Content += RenderSubtitle("热股榜")
		article.Content += RenderStockTable(stockFollow)
		article.Content += RenderPlaceholder()
	}
	if len(stockMao20) > 0 {
		article.Content += RenderSubtitle("招财大牛猫茅 20 组合")
		article.Content += RenderStockTable(stockMao20)
		article.Content += RenderPlaceholder()
	}
	if len(stockPercent) > 0 {
		article.Content += RenderSubtitle("日涨幅榜")
		article.Content += RenderStockTable(stockPercent)
		article.Content += RenderPlaceholder()
	}
	if len(stockCurrentYearPercent) > 0 {
		article.Content += RenderSubtitle("年涨幅榜")
		article.Content += RenderStockTable(stockCurrentYearPercent)
		article.Content += RenderPlaceholder()
	}
	if len(stockCurrent) > 0 {
		article.Content += RenderSubtitle("价格榜")
		article.Content += RenderStockTable(stockCurrent)
		article.Content += RenderPlaceholder()
	}
	if len(stockMarketCapital) > 0 {
		article.Content += RenderSubtitle("市值榜/万亿")
		article.Content += RenderStockTable(stockMarketCapital)
		article.Content += RenderPlaceholder()
	}
	if len(stockVolume) > 0 {
		article.Content += RenderSubtitle("成交量榜/亿")
		article.Content += RenderStockTable(stockVolume)
		article.Content += RenderPlaceholder()
	}
	if len(stockAmount) > 0 {
		article.Content += RenderSubtitle("成交额榜/亿")
		article.Content += RenderStockTable(stockAmount)
		article.Content += RenderPlaceholder()
	}
	if len(stockForeign) > 0 {
		article.Content += RenderSubtitle("主力净流入榜/亿")
		article.Content += RenderStockTable(stockForeign)
		article.Content += RenderPlaceholder()
	}
	if len(stockEvent) > 0 {
		article.Content += RenderSubtitle("行情回顾")
		article.Content += RenderStockTimeline(stockEvent)
		article.Content += RenderPlaceholder()
	}
	if len(stockChance) > 0 {
		article.Content += RenderSubtitle("市场机会")
		article.Content += RenderContent(stockChance)
		article.Content += RenderPlaceholder()
	}
	author, motto := GetMotto()
	article.Content += RenderSubtitle(author)
	article.Content += RenderContent([]string{motto})
	article.Content += RenderPlaceholder()
	article.Content += RenderFooter("牛市三连「分享」「点赞」「在看」👇")

	log.Println(article.Title)
	log.Println(article.Digest)
	log.Println(article.Content)

	return article
}
