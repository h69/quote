package main

import (
	"fmt"
	"strings"
)

// RenderHeader 页头
func RenderHeader(header string) string {
	return `<p style="white-space: normal;"><span style="font-size: 15px; color: rgb(255, 255, 255); background-color: rgb(72, 91, 247); padding: 5px;"><strong>` + header + `</strong></span></p>`
}

// RenderFooter 页尾
func RenderFooter(footer string) string {
	return `<p style="white-space: normal; text-align: right;"><span style="font-size: 15px; color: rgb(255, 255, 255); background-color: rgb(72, 91, 247); padding: 5px;"><strong>` + footer + `</strong></span></p>`
}

// RenderTitle 主标题
func RenderTitle(title string) string {
	return `<p style="text-align: center; padding-right: 16px; padding-left: 16px; margin-bottom: 10px;"><span style="color: rgb(72, 91, 247); font-size: 24px; letter-spacing: 0.5px; line-height: 1.75em;"><strong>` + title + `</strong></span></p>`
}

// RenderSubtitle 子标题
func RenderSubtitle(subtitle string) string {
	return `<p style="text-align: center; padding-right: 16px; padding-left: 16px; margin-bottom: 10px;"><span style="color: rgb(72, 91, 247); font-size: 17px; letter-spacing: 0.5px; line-height: 1.75em;"><strong>` + subtitle + `</strong></span></p>`
}

// RenderPlaceholder 占位符
func RenderPlaceholder() string {
	return `<p style="text-align: center;"><br></p>`
}

// RenderContent 内容
func RenderContent(contents []string) string {
	html := `<p style="height: 10px; min-height: 0px;"></p>`
	for i, content := range contents {
		html += `<p style="padding-right: 10px; padding-left: 10px;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.75em; white-space: normal;">` + content + `</span></p>`
		if i < len(contents)-1 {
			html += `<p style="text-align: center;"><br></p>`
		}
	}
	return html
}

// RenderStockTable 股票表格
func RenderStockTable(stocks []Stock) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	for i := 0; i < len(stocks); i++ {
		var percentColor string
		if stocks[i].Percent == "0.00%" {
			percentColor = "color: rgb(153, 153, 153);"
		} else if strings.Contains(stocks[i].Percent, "+") {
			percentColor = "color: rgb(246, 66, 69);"
		} else {
			percentColor = "color: rgb(0, 171, 59);"
		}

		var currentYearPercentColor string
		if stocks[i].CurrentYearPercent == "0.00%" {
			currentYearPercentColor = "color: rgb(153, 153, 153);"
		} else if strings.Contains(stocks[i].CurrentYearPercent, "+") {
			currentYearPercentColor = "color: rgb(246, 66, 69);"
		} else {
			currentYearPercentColor = "color: rgb(0, 171, 59);"
		}

		html += `<tr style="border: 0px;">`
		html += `<td style="width: 30%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em;">` + stocks[i].Name + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; color: #888;">` + stocks[i].Code + `</span></td>`
		html += `<td style="width: 25%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px; text-align: right;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.75em;">` + stocks[i].Value + `</span></td>`
		html += `<td style="width: 45%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px; text-align: right;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em; ` + percentColor + `">` + stocks[i].Current + `</span><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em; ` + percentColor + `"> ` + stocks[i].Percent + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; ` + currentYearPercentColor + `">` + stocks[i].CurrentYearPercent + `</span></td>`
		html += `</tr>`
	}
	html += `</tbody></table>`
	return html
}

// RenderStockChart 股票柱状图
func RenderStockChart(bars []Bar) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	html += `<tr style="border: 0px;">`
	for i := 0; i < len(bars); i++ {
		var color, backgroundColor string
		if bars[i].Flag == 1 {
			color = "color: rgb(246, 66, 69);"
			backgroundColor = "background-color: rgb(246, 66, 69);"
		} else if bars[i].Flag == 0 {
			color = "color: rgb(153, 153, 153);"
			backgroundColor = "background-color: rgb(153, 153, 153);"
		} else {
			color = "color: rgb(0, 171, 59);"
			backgroundColor = "background-color: rgb(0, 171, 59);"
		}

		html += `<td valign="bottom" style="border: 0px; text-align: center;"><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; ` + color + `">` + Float64ToString(bars[i].Value) + `</span><br/><span style="display: inline-block; border-radius: 5px 5px 0 0; width: 100%; height: ` + fmt.Sprintf("%.2f", 200*bars[i].Ratio) + "px; " + backgroundColor + `"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em;">` + bars[i].Name + `</span></td>`
	}
	html += `</tr>`
	html += `</tbody></table>`
	return html
}

// RenderStockThermometer 股票温度计
func RenderStockThermometer(bars []Bar) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	html += `<tr style="border: 0px;">`
	var down, flat, up float64
	for i := 0; i < len(bars); i++ {
		if bars[i].Flag == 1 {
			up += bars[i].Value
		} else if bars[i].Flag == 0 {
			flat += bars[i].Value
		} else {
			down += bars[i].Value
		}
	}

	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", down/(down+flat+up)*100) + `%-25px; border: 0px; text-align: center; padding-left: 8px; padding-right: 0px; padding-top: 12px;"><span style="display: inline-block; border-radius: 5px 0 0 5px; width: 100%; height: 5px; background-color: rgb(0, 171, 59);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(0, 171, 59);">` + Float64ToString(down) + `</span></td>`
	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", flat/(down+flat+up)*100) + `%; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px; padding-top: 12px;"><span style="display: inline-block; width: 100%; height: 5px; background-color: rgb(153, 153, 153);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(153, 153, 153); display: none;">` + Float64ToString(flat) + `</span></td>`
	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", up/(down+flat+up)*100) + `%-25px; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px; padding-top: 12px;"><span style="display: inline-block; border-radius: 0 5px 5px 0; width: 100%; height: 5px; background-color: rgb(246, 66, 69);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(246, 66, 69);">` + Float64ToString(up) + `</span></td>`
	html += `<td valign="top" style="width: 50px; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px;"><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(246, 66, 69);">` + fmt.Sprintf("%.0f", (up/(up+flat+down))*100) + "℃" + `</span></td>`
	html += `</tr>`
	html += `</tbody></table>`
	return html
}

// RenderStockTimeline 股票时间线
func RenderStockTimeline(events []Event) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	for i := 0; i < len(events); i++ {
		html += `<tr style="border: 0px;">`
		html += `<td valign="top" style="width: 70px; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.75em;"><strong>` + events[i].Time + `</strong></span></td>`
		html += `<td valign="top" style="border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.75em;">` + events[i].Content + `</span></td>`
		html += `</tr>`
	}
	html += `</tbody></table>`
	return html
}

// RenderStockPlate 股票板块图
func RenderStockPlate(stocks []Stock) string {
	html := `<table style="width: 100%; padding-right: 2px; padding-left: 2px; border: 0px; line-height: 0; border-collapse: separate; border-spacing: 5px 5px;"><tbody>`
	for i := 0; i < len(stocks); i++ {
		var backgroundColor string
		if stocks[i].Percent == "0.00%" {
			backgroundColor = "background-color: rgb(153, 153, 153);"
		} else if strings.Contains(stocks[i].Percent, "+") {
			backgroundColor = "background-color: rgb(246, 66, 69);"
		} else {
			backgroundColor = "background-color: rgb(0, 171, 59);"
		}

		if i%3 == 0 {
			html += `<tr style="border: 0px;">`
		}

		var borderRadius string
		if i == 0 {
			borderRadius = "border-radius: 10px 0 0 0;"
		} else if i == 2 {
			borderRadius = "border-radius: 0 10px 0 0;"
		} else if i == len(stocks)-1 {
			borderRadius = "border-radius: 0 0 10px 0;"
		} else if i == len(stocks)-3 {
			borderRadius = "border-radius: 0 0 0 10px;"
		}

		html += `<td style="width: 33.33%; height: 50px; border: 0px; border-bottom: 0px dashed #ccc; text-align: center; ` + backgroundColor + " " + borderRadius + `"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em; color: #fff;">` + stocks[i].Name + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; color: #fff;">` + stocks[i].Percent + `</span></td>`

		if (i+1)%3 == 0 {
			html += `</tr>`
		}
	}
	html += `</tbody></table>`
	return html
}

// RenderStockCard 股票卡片图
func RenderStockCard(stocks []Stock) string {
	html := `<table style="width: 100%; padding-right: 2px; padding-left: 2px; border: 0px; line-height: 0; border-collapse: separate; border-spacing: 5px 5px;"><tbody>`
	for i := 0; i < len(stocks); i++ {
		var backgroundColor string
		if strings.Index(stocks[i].Value, "0.00") == 0 {
			backgroundColor = "background-color: rgb(153, 153, 153);"
		} else if strings.Contains(stocks[i].Value, "+") {
			backgroundColor = "background-color: rgb(246, 66, 69);"
		} else {
			backgroundColor = "background-color: rgb(0, 171, 59);"
		}

		if i%2 == 0 {
			html += `<tr style="border: 0px;">`
		}

		html += `<td style="width: 50%; height: 50px; border: 0px; border-bottom: 0px dashed #ccc; text-align: center; border-radius: 10px;` + backgroundColor + `"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em; color: #fff;">` + stocks[i].Name + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; color: #fff;">` + stocks[i].Value + `</span></td>`

		if (i+1)%2 == 0 {
			html += `</tr>`
		}
	}
	html += `</tbody></table>`
	return html
}
