package main

import (
	"fmt"
	"strings"
)

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

// RenderStockTable 股票表格
func RenderStockTable(stocks []Stock) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	for i := 0; i < len(stocks); i++ {
		var color string
		if stocks[i].Percent == "0.00%" {
			color = "color: rgb(153, 153, 153);"
		} else if strings.Contains(stocks[i].Percent, "+") {
			color = "color: rgb(246, 66, 69);"
		} else {
			color = "color: rgb(0, 171, 59);"
		}

		html += `<tr style="border: 0px;">`
		html += `<td style="width: 33.33%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em;">` + stocks[i].Name + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; color: #888;">` + stocks[i].Code + `</span></td>`
		html += `<td style="width: 33.33%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px; text-align: right;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.75em;">` + stocks[i].Value + `</span></td>`
		html += `<td style="width: 33.33%; border: 0px; border-bottom: 0px dashed #ccc; padding-top: 10px; padding-bottom: 10px; text-align: right;"><span style="font-size: 15px; letter-spacing: 0.5px; line-height: 1.25em; ` + color + `">` + stocks[i].Current + `</span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.25em; ` + color + `">` + stocks[i].Percent + `</span></td>`
		html += `</tr>`
	}
	html += `</tbody></table>`
	return html
}

// RenderStockChart 股票分布图
func RenderStockChart(bars []Bar) string {
	html := `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	html += `<tr style="border: 0px;">`
	var down, flat, up float64
	for i := 0; i < len(bars); i++ {
		var color, backgroundColor string
		if bars[i].Flag == 1 {
			color = "color: rgb(246, 66, 69);"
			backgroundColor = "background-color: rgb(246, 66, 69);"
			up += bars[i].Value
		} else if bars[i].Flag == 0 {
			color = "color: rgb(153, 153, 153);"
			backgroundColor = "background-color: rgb(153, 153, 153);"
			flat += bars[i].Value
		} else {
			color = "color: rgb(0, 171, 59);"
			backgroundColor = "background-color: rgb(0, 171, 59);"
			down += bars[i].Value
		}

		html += `<td valign="bottom" style="border: 0px; text-align: center;"><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; ` + color + `">` + Float64ToString(bars[i].Value) + `</span><br/><span style="display: inline-block; border-radius: 5px 5px 0 0; width: 100%; height: ` + fmt.Sprintf("%.2f", 200*bars[i].Ratio) + "px; " + backgroundColor + `"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em;">` + bars[i].Name + `</span></td>`
	}
	html += `</tr>`
	html += `</tbody></table>`

	html += `<table style="width: 100%; padding-right: 16px; padding-left: 16px; border: 0px; line-height: 0;"><tbody>`
	html += `<tr style="border: 0px;">`
	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", down/(down+flat+up)*100) + `%; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px;"><span style="display: inline-block; border-radius: 5px 0 0 5px; width: 100%; height: 5px; background-color: rgb(0, 171, 59);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(0, 171, 59);">` + Float64ToString(down) + `</span></td>`
	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", flat/(down+flat+up)*100) + `%; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px;"><span style="display: inline-block; width: 100%; height: 5px; background-color: rgb(153, 153, 153);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(153, 153, 153); display: none;">` + Float64ToString(flat) + `</span></td>`
	html += `<td valign="top" style="width: ` + fmt.Sprintf("%.2f", up/(down+flat+up)*100) + `%; border: 0px; text-align: center; padding-left: 0px; padding-right: 0px;"><span style="display: inline-block; border-radius: 0 5px 5px 0; width: 100%; height: 5px; background-color: rgb(246, 66, 69);"></span><br/><span style="font-size: 12px; letter-spacing: 0.5px; line-height: 1.75em; color: rgb(246, 66, 69);">` + Float64ToString(up) + `</span></td>`
	html += `</tr>`
	html += `</tbody></table>`

	return html
}
