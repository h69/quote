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
	Name               string
	Code               string
	Current            string
	Percent            string
	CurrentYearPercent string
	Value              string
}

// Event 事件
type Event struct {
	Content string
	Time    string
}

// Content 内容
type Content struct {
	Content string
	Strong  bool
}

// GetStockCloseTime 获取股票休市时间
func GetStockCloseTime() string {
	var date string

	body, err := Get("https://x-quote.cls.cn/quote/stock/tline?app=CailianpressWeb&fields=date&secu_code=sh000001")
	if err != nil {
		log.Println(err)
		return date
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
		log.Println(err)
		return date
	}

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
		log.Println(err)
		return bars
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
		log.Println(err)
		return bars
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
		log.Println(err)
		return stocks
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
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.IndexQuote {
		stocks = append(stocks, Stock{
			Name:               stock.SecuName,
			Code:               strings.ToUpper(stock.SecuCode),
			Current:            fmt.Sprintf("%.2f", stock.LastPx),
			Percent:            GetPercentSign(stock.Change) + fmt.Sprintf("%.2f", stock.Change*100) + "%",
			CurrentYearPercent: GetStockCurrentYearPercentByCode(stock.SecuCode),
			Value:              "",
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
		log.Println(err)
		return stocks
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
				MarketCapital      float64 `json:"market_capital"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              fmt.Sprintf("%.2f", stock.MarketCapital/100000000/10000),
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
		log.Println(err)
		return stocks
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
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              "",
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
		log.Println(err)
		return stocks
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
				Volume             float64 `json:"volume"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              fmt.Sprintf("%.2f", stock.Volume/100000000),
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
		log.Println(err)
		return stocks
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
				Amount             float64 `json:"amount"`
			} `json:"list"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              fmt.Sprintf("%.2f", stock.Amount/100000000),
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
		log.Println(err)
		return stocks
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
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              "",
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
		log.Println(err)
		return stocks
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
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Percent) + fmt.Sprintf("%.2f", stock.Percent) + "%",
			CurrentYearPercent: GetPercentSign(stock.CurrentYearPercent) + fmt.Sprintf("%.2f", stock.CurrentYearPercent) + "%",
			Value:              "",
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
		log.Println(err)
		return stocks
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
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.List {
		stocks = append(stocks, Stock{
			Name:               stock.Name,
			Code:               stock.Symbol,
			Current:            fmt.Sprintf("%.2f", stock.Current),
			Percent:            GetPercentSign(stock.Pct) + fmt.Sprintf("%.2f", stock.Pct) + "%",
			CurrentYearPercent: GetStockCurrentYearPercentByCode(stock.Symbol),
			Value:              "",
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
			log.Println(err)
			return []Stock{}
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
			log.Println(err)
			return []Stock{}
		}

		stocks = append(stocks, Stock{
			Name:               resp.Data.SecuName,
			Code:               strings.ToUpper(resp.Data.SecuCode),
			Current:            fmt.Sprintf("%.2f", resp.Data.LastPx),
			Percent:            GetPercentSign(resp.Data.Change) + fmt.Sprintf("%.2f", resp.Data.Change*100) + "%",
			CurrentYearPercent: GetStockCurrentYearPercentByCode(resp.Data.SecuCode),
			Value:              "",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockEvent 获取股票事件
func GetStockEvent() []Event {
	var events []Event

	body, err := Get("https://api.wallstcn.com/apiv1/search/live?channel=a-stock-channel&limit=100&score=2")
	if err != nil {
		log.Println(err)
		return events
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
		log.Println(err)
		return events
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

// GetStockCurrentYearPercentByCode 获取股票年涨幅
func GetStockCurrentYearPercentByCode(code string) string {
	var currentYearPercent string

	body, err := Get("https://x-quote.cls.cn/quote/stock/kline?app=CailianpressWeb&limit=1&offset=0&type=fy&secu_code=" + strings.ToLower(code))
	if err != nil {
		log.Println(err)
		return currentYearPercent
	}

	log.Println(body)

	type Resp struct {
		Data []struct {
			Change float64 `json:"change"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return currentYearPercent
	}

	if len(resp.Data) > 0 {
		currentYearPercent = GetPercentSign(resp.Data[0].Change) + fmt.Sprintf("%.2f", resp.Data[0].Change*100) + "%"
	}

	log.Println(currentYearPercent)

	return currentYearPercent
}

// GetStockIndustry 获取股票热门行业
func GetStockIndustry() []Stock {
	var stocks []Stock

	body, err := Get("https://x-quote.cls.cn/web_quote/plate/plate_list?app=CailianpressWeb&type=industry&page=1&rever=1")
	if err != nil {
		log.Println(err)
		return stocks
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			PlateData []struct {
				SecuName string  `json:"secu_name"`
				Change   float64 `json:"change"`
			} `json:"plate_data"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	if len(resp.Data.PlateData) >= 9 {
		resp.Data.PlateData = resp.Data.PlateData[:9]
	} else if len(resp.Data.PlateData) >= 6 {
		resp.Data.PlateData = resp.Data.PlateData[:6]
	} else {
		return stocks
	}

	for _, stock := range resp.Data.PlateData {
		stocks = append(stocks, Stock{
			Name:               stock.SecuName,
			Code:               "",
			Current:            "",
			Percent:            GetPercentSign(stock.Change) + fmt.Sprintf("%.2f", stock.Change*100) + "%",
			CurrentYearPercent: "",
			Value:              "",
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockChance 获取股票机会
func GetStockChance() []string {
	var contents []string

	body, err := Get("https://www.cls.cn/api/subject/recommend/article?app=CailianpressWeb&os=web&sv=7.5.5&sign=091f85dd6463852f8445039d60b4633d", "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	if err != nil {
		log.Println(err)
		return contents
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			ProspectArticle struct {
				Title string `json:"title"`
			} `json:"prospect_article"`
			TodayChances []struct {
				SubjectName string `json:"subject_name"`
				ArticleName string `json:"article_name"`
				StockList   []struct {
					Name string `json:"name"`
				} `json:"stock_list"`
			} `json:"today_chances"`
			TodayTuyeres []struct {
				SubjectName string `json:"subject_name"`
				Driver      string `json:"driver"`
				Stocks      []struct {
					Name string `json:"name"`
				} `json:"stocks"`
			} `json:"today_tuyeres"`
			ShortLatents []struct {
				SubjectName        string `json:"subject_name"`
				SubjectDescription string `json:"subject_description"`
				CashTime           string `json:"cash_time"`
			} `json:"short_latents"`
			LongChances []struct {
				SubjectName string `json:"subject_name"`
				ArticleName string `json:"article_name"`
				Stocks      []struct {
					Name string `json:"name"`
				} `json:"stocks"`
			} `json:"long_chances"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return contents
	}

	for _, chance := range resp.Data.TodayChances {
		var stocks []string
		for _, stock := range chance.StockList {
			stocks = append(stocks, stock.Name)
		}
		var comment string
		if len(stocks) > 0 {
			comment = "（" + strings.Join(stocks, "、") + "）"
		}
		contents = append(contents, chance.SubjectName+"："+strings.TrimRight(chance.ArticleName, "。")+comment)
	}
	for _, tuyere := range resp.Data.TodayTuyeres {
		var stocks []string
		for _, stock := range tuyere.Stocks {
			stocks = append(stocks, stock.Name)
		}
		var comment string
		if len(stocks) > 0 {
			comment = "（" + strings.Join(stocks, "、") + "）"
		}
		contents = append(contents, tuyere.SubjectName+"："+strings.TrimRight(tuyere.Driver, "。")+comment)
	}
	for _, chance := range resp.Data.LongChances {
		var stocks []string
		for _, stock := range chance.Stocks {
			stocks = append(stocks, stock.Name)
		}
		var comment string
		if len(stocks) > 0 {
			comment = "（" + strings.Join(stocks, "、") + "）"
		}
		contents = append(contents, chance.SubjectName+"："+strings.TrimRight(chance.ArticleName, "。")+comment)
	}
	for _, latent := range resp.Data.ShortLatents {
		var comment string
		if latent.CashTime != "" {
			comment = "（" + latent.CashTime + "）"
		}
		contents = append(contents, latent.SubjectName+"："+strings.TrimRight(latent.SubjectDescription, "。")+comment)
	}

	log.Println(Sprintf(contents))

	return contents
}

// GetStockMarketForeign 获取股票北向资金
func GetStockMarketForeign() []Stock {
	var stocks []Stock

	body, err := Get("http://push2.eastmoney.com/api/qt/kamt/get?fields1=f1,f2,f3,f4&fields2=f51,f52,f53,f54,f63")
	if err != nil {
		log.Println(err)
		return stocks
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			Hk2sh struct {
				DayNetAmtIn float64 `json:"dayNetAmtIn"`
			} `json:"hk2sh"`
			Hk2sz struct {
				DayNetAmtIn float64 `json:"dayNetAmtIn"`
			} `json:"hk2sz"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	stocks = append(stocks, Stock{
		Name:               "沪股通",
		Code:               "",
		Current:            "",
		Percent:            "",
		CurrentYearPercent: "",
		Value:              GetValueSign(resp.Data.Hk2sh.DayNetAmtIn) + fmt.Sprintf("%.2f", resp.Data.Hk2sh.DayNetAmtIn/10000) + "亿",
	})
	stocks = append(stocks, Stock{
		Name:               "深股通",
		Code:               "",
		Current:            "",
		Percent:            "",
		CurrentYearPercent: "",
		Value:              GetValueSign(resp.Data.Hk2sz.DayNetAmtIn) + fmt.Sprintf("%.2f", resp.Data.Hk2sz.DayNetAmtIn/10000) + "亿",
	})

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockForeign 获取股票净流入
func GetStockForeign() []Stock {
	var stocks []Stock

	body, err := Get("https://x-quote.cls.cn/web_quote/web_stock/stock_list?app=CailianpressWeb&market=all&os=web&page=1&rever=1&sv=7.5.5&types=last_px,change,tr,main_fund_diff,cmc&way=main_fund_diff&sign=a1017682e6f5e28779457f61d7bee9c3")
	if err != nil {
		log.Println(err)
		return stocks
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			Data []struct {
				SecuName     string  `json:"secu_name"`
				SecuCode     string  `json:"secu_code"`
				LastPx       float64 `json:"last_px"`
				Change       float64 `json:"change"`
				MainFundDiff float64 `json:"main_fund_diff"`
			} `json:"data"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return stocks
	}

	for _, stock := range resp.Data.Data {
		if len(stocks) >= 10 {
			break
		}

		stocks = append(stocks, Stock{
			Name:               stock.SecuName,
			Code:               strings.ToUpper(stock.SecuCode),
			Current:            fmt.Sprintf("%.2f", stock.LastPx),
			Percent:            GetPercentSign(stock.Change) + fmt.Sprintf("%.2f", stock.Change*100) + "%",
			CurrentYearPercent: GetStockCurrentYearPercentByCode(stock.SecuCode),
			Value:              fmt.Sprintf("%.2f", stock.MainFundDiff/100000000),
		})
	}

	log.Println(Sprintf(stocks))

	return stocks
}

// GetStockComment 获取股票评论
func GetStockComment() []string {
	var contents []string

	body, err := Get("https://www.cls.cn/nodeapi/telegraphList?app=CailianpressWeb&refresh_type=1&rn=50", "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	if err != nil {
		log.Println(err)
		return contents
	}

	log.Println(body)

	type Resp struct {
		Data struct {
			RollData []struct {
				Title   string `json:"title"`
				Content string `json:"content"`
				Ctime   int64  `json:"ctime"`
			} `json:"roll_data"`
		} `json:"data"`
	}

	var resp Resp
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Println(err)
		return contents
	}

	closeTime := GetStockCloseTime()

	for _, data := range resp.Data.RollData {
		if data.Ctime >= GetTimestamp(closeTime+" 15:00:00") || strings.Contains(data.Title, "收评：") {
			title := strings.TrimLeft(data.Title, "收评：")
			title = "<strong>" + title + "</strong>"
			content := strings.Replace(data.Content, "【"+data.Title+"】", "", 1)
			content = strings.TrimLeft(content, "财联社")
			content = string([]rune(content)[Find(content, "，")+1:])
			contents = append(contents, title)
			contents = append(contents, content)
			break
		}
	}

	log.Println(Sprintf(contents))

	return contents
}
