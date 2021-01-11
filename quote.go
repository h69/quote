package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

// Bar 股票分布
type Bar struct {
	Name  string
	Value float64
	Ratio float64
	Flag  int
}

// Stock 股票
type Stock struct {
	Name    string
	Code    string
	Value   string
	Current string
	Percent string
}

// Event 事件
type Event struct {
	Content string
	Time    string
}

// GetStockCloseTime 获取股票休市时间
func GetStockCloseTime() string {
	body, err := Get("https://x-quote.cls.cn/quote/stock/tline?app=CailianpressWeb&fields=date&secu_code=sh000001")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			Date []int64 `json:"date"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	var date string
	if len(resp.Data.Date) > 0 {
		for i, s := range Int64ToString(resp.Data.Date[0]) {
			date += string(s)
			if i == 3 || i == 5 {
				date += "/"
			}
		}
	}

	log.Println(date)

	return date
}

// GetStockMarketOverview 获取股票市场总览
func GetStockMarketOverview() []Bar {
	var bars []Bar

	body, err := Get("https://x-quote.cls.cn/quote/index/home?app=CailianpressWeb")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			UpDownDis struct {
				AverageRise float64 `json:"average_rise"`
				Down2       float64 `json:"down_2"`
				Down4       float64 `json:"down_4"`
				Down6       float64 `json:"down_6"`
				Down8       float64 `json:"down_8"`
				Down10      float64 `json:"down_10"`
				DownNum     float64 `json:"down_num"`
				FallNum     float64 `json:"fall_num"`
				FlatNum     float64 `json:"flat_num"`
				RiseNum     float64 `json:"rise_num"`
				SuspendNum  float64 `json:"suspend_num"`
				Up2         float64 `json:"up_2"`
				Up4         float64 `json:"up_4"`
				Up6         float64 `json:"up_6"`
				Up8         float64 `json:"up_8"`
				Up10        float64 `json:"up_10"`
				UpNum       float64 `json:"up_num"`
			} `json:"up_down_dis"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	bars = append(bars, Bar{
		Name:  "跌停",
		Value: resp.Data.UpDownDis.DownNum,
		Ratio: resp.Data.UpDownDis.DownNum / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  -1,
	})
	bars = append(bars, Bar{
		Name:  ">6",
		Value: resp.Data.UpDownDis.Down10 + resp.Data.UpDownDis.Down8,
		Ratio: (resp.Data.UpDownDis.Down10 + resp.Data.UpDownDis.Down8) / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  -1,
	})
	bars = append(bars, Bar{
		Name:  "0-6",
		Value: resp.Data.UpDownDis.Down6 + resp.Data.UpDownDis.Down4 + resp.Data.UpDownDis.Down2,
		Ratio: (resp.Data.UpDownDis.Down6 + resp.Data.UpDownDis.Down4 + resp.Data.UpDownDis.Down2) / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  -1,
	})
	bars = append(bars, Bar{
		Name:  "0",
		Value: resp.Data.UpDownDis.FlatNum,
		Ratio: resp.Data.UpDownDis.FlatNum / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  0,
	})
	bars = append(bars, Bar{
		Name:  "0-6",
		Value: resp.Data.UpDownDis.Up2 + resp.Data.UpDownDis.Up4 + resp.Data.UpDownDis.Up6,
		Ratio: (resp.Data.UpDownDis.Up2 + resp.Data.UpDownDis.Up4 + resp.Data.UpDownDis.Up6) / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  1,
	})
	bars = append(bars, Bar{
		Name:  ">6",
		Value: resp.Data.UpDownDis.Up8 + resp.Data.UpDownDis.Up10,
		Ratio: (resp.Data.UpDownDis.Up8 + resp.Data.UpDownDis.Up10) / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  1,
	})
	bars = append(bars, Bar{
		Name:  "涨停",
		Value: resp.Data.UpDownDis.UpNum,
		Ratio: resp.Data.UpDownDis.UpNum / (resp.Data.UpDownDis.FallNum + resp.Data.UpDownDis.FlatNum + resp.Data.UpDownDis.RiseNum + resp.Data.UpDownDis.SuspendNum),
		Flag:  1,
	})

	log.Println(Sprintf(bars))

	return bars
}

// GetStockMarketIndex 获取股票大盘指数
func GetStockMarketIndex() []Stock {
	var stocks []Stock

	body, err := Get("https://x-quote.cls.cn/quote/index/home?app=CailianpressWeb")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			IndexQuote []struct {
				SecuName string  `json:"secu_name"`
				SecuCode string  `json:"secu_code"`
				LastPx   float64 `json:"last_px"`
				Change   float64 `json:"change"`
			} `json:"index_quote"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.IndexQuote {
		var flag string
		if stock.Change > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.SecuName,
			Code:    strings.ToUpper(stock.SecuCode),
			Value:   "",
			Current: fmt.Sprintf("%.2f", stock.LastPx),
			Percent: flag + fmt.Sprintf("%.2f", stock.Change*100) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockMarketCapital 获取股票市值
func GetStockMarketCapital() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=market_capital&order_by=market_capital&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name          string  `json:"name"`
				Symbol        string  `json:"symbol"`
				Current       float64 `json:"current"`
				Percent       float64 `json:"percent"`
				MarketCapital float64 `json:"market_capital"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   fmt.Sprintf("%.2f", stock.MarketCapital/100000000/10000) + " 万亿",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockCurrent 获取股票当前价
func GetStockCurrent() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=current&order_by=current&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name    string  `json:"name"`
				Symbol  string  `json:"symbol"`
				Percent float64 `json:"percent"`
				Current float64 `json:"current"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   "",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockVolume 获取股票成交量
func GetStockVolume() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=volume&order_by=volume&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name    string  `json:"name"`
				Symbol  string  `json:"symbol"`
				Current float64 `json:"current"`
				Percent float64 `json:"percent"`
				Volume  float64 `json:"volume"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   fmt.Sprintf("%.2f", stock.Volume/100000000) + " 亿",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockAmount 获取股票成交额
func GetStockAmount() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=amount&order_by=amount&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name    string  `json:"name"`
				Symbol  string  `json:"symbol"`
				Current float64 `json:"current"`
				Percent float64 `json:"percent"`
				Amount  float64 `json:"amount"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   fmt.Sprintf("%.2f", stock.Amount/100000000) + " 亿",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockCurrentYearPercent 获取股票年涨幅
func GetStockCurrentYearPercent() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=current_year_percent&order_by=current_year_percent&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name               string  `json:"name"`
				Symbol             string  `json:"symbol"`
				Current            float64 `json:"current"`
				Percent            float64 `json:"percent"`
				CurrentYearPercent float64 `json:"current_year_percent"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		var flag2 string
		if stock.CurrentYearPercent > 0 {
			flag2 = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   flag2 + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockPercent 获取股票日涨幅
func GetStockPercent() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=10&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name    string  `json:"name"`
				Symbol  string  `json:"symbol"`
				Current float64 `json:"current"`
				Percent float64 `json:"percent"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Percent > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   "",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Percent) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockFollow 获取股票热门
func GetStockFollow() []Stock {
	var stocks []Stock

	body, err := Get("https://xueqiu.com/service/v5/stock/screener/screen?category=CN&size=10&order=desc&order_by=follow7d&only_count=0&page=1")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			List []struct {
				Name    string  `json:"name"`
				Symbol  string  `json:"symbol"`
				Current float64 `json:"current"`
				Pct     float64 `json:"pct"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	for _, stock := range resp.Data.List {
		var flag string
		if stock.Pct > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    stock.Name,
			Code:    stock.Symbol,
			Value:   "",
			Current: fmt.Sprintf("%.2f", stock.Current),
			Percent: flag + fmt.Sprintf("%.2f", stock.Pct) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockMao20 获取股票茅20组合
func GetStockMao20() []Stock {
	var stocks []Stock

	codes := []string{
		"sh600036",
		"sz300015",
		"sz300760",
		"sz002714",
		"sh603288",
		"sz000858",
		"sz300059",
		"sh600519",
		"sh601318",
		"sh600900",
		"sh603259",
		"sz300750",
		"sh600031",
		"sh600309",
		"sh600276",
		"sz000333",
		"sh600887",
		"sz002352",
		"sz002475",
		"sh601888",
	}

	for _, code := range codes {
		body, err := Get("https://x-quote.cls.cn/quote/stock/basic?secu_code=" + code + "&fields=secu_name,secu_code,last_px,change&app=CailianpressWeb")
		if err != nil {
			panic(err)
		}

		log.Println(body)

		type Resp struct {
			Data struct {
				SecuName string  `json:"secu_name"`
				SecuCode string  `json:"secu_code"`
				LastPx   float64 `json:"last_px"`
				Change   float64 `json:"change"`
			} `json:"data"`
		}

		var resp Resp
		err = json.Unmarshal([]byte(body), &resp)
		if err != nil {
			panic(err)
		}

		var flag string
		if resp.Data.Change > 0 {
			flag = "+"
		}

		stocks = append(stocks, Stock{
			Name:    resp.Data.SecuName,
			Code:    strings.ToUpper(resp.Data.SecuCode),
			Value:   "",
			Current: fmt.Sprintf("%.2f", resp.Data.LastPx),
			Percent: flag + fmt.Sprintf("%.2f", resp.Data.Change*100) + "%",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockEvents 获取股票事件
func GetStockEvents() []Event {
	var events []Event

	body, err := Get("https://api.wallstcn.com/apiv1/search/live?channel=a-stock-channel&limit=100&score=2")
	if err != nil {
		panic(err)
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			Items []struct {
				ContentText string `json:"content_text"`
				DisplayTime int64  `json:"display_time"`
			} `json:"items"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		panic(err)
	}

	closeTime := GetStockCloseTime()

	for _, event := range resp.Data.Items {
		if event.DisplayTime >= GetTimestamp(closeTime+" 09:00:00") && event.DisplayTime <= GetTimestamp(closeTime+" 15:00:00") && utf8.RuneCountInString(event.ContentText) <= 100 {
			events = append(events, Event{
				Content: strings.ReplaceAll(event.ContentText, "\n", ""),
				Time:    GetHourMinute(event.DisplayTime),
			})
		}
	}

	log.Println(Sprintf(events))

	return events
}
