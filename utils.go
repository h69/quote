package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

// Convert 编码转换
func Convert(src string, srcEncoder string, desEncoder string) string {
	_, resBytes, _ := mahonia.NewDecoder(desEncoder).Translate([]byte(mahonia.NewDecoder(srcEncoder).ConvertString(src)), true)
	return string(resBytes)
}

// Get 请求
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// PostJSON 请求
func PostJSON(url string, v interface{}) (string, error) {
	data, _ := json.Marshal(v)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// PostForm 请求
func PostForm(url string, data url.Values) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// FindInvert 从后往前查找下标
func FindInvert(str string, char string) int {
	for i := len([]rune(str)) - 1; i >= 0; i-- {
		if string([]rune(str)[i]) == char {
			return i
		}
	}
	return len([]rune(str))
}

// Count 计算字符个数
func Count(str string, char string) int {
	count := 0
	for i := 0; i < len([]rune(str)); i++ {
		if string([]rune(str)[i]) == char {
			count++
		}
	}
	return count
}

// Int64ToString int64 -> string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Int32ToString int32 -> string
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// IntToString int -> string
func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// Float64ToString float64 -> string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// ByteToString byte -> string
func ByteToString(b []byte) string {
	return string(b)
}

// StringToByte string -> byte
func StringToByte(s string) []byte {
	return []byte(s)
}

// StringToInt64 string -> int64
func StringToInt64(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	return 0
}

// StringToInt32 string -> int32
func StringToInt32(s string) int32 {
	if i, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int32(i)
	}
	return 0
}

// StringToInt string -> int
func StringToInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

// GetDate 获取日期
func GetDate(days int) string {
	return time.Now().AddDate(0, 0, days).Format("2006/01/02")
}

// GetTimestamp 获取时间戳
func GetTimestamp(datetime string) int64 {
	t, _ := time.ParseInLocation("2006/01/02 15:04:05", datetime, time.Local)
	return t.Unix()
}

// GetHourMinute 获取小时分钟
func GetHourMinute(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("15:04")
}

// GetPercentSign 获取百分比符号
func GetPercentSign(percent float64) string {
	var sign string
	if percent > 0 {
		sign = "+"
	}
	return sign
}

// Sprintf 格式化对象结构
func Sprintf(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	return string(b)
}
