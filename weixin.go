// https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
package main

import (
	"encoding/json"
	"log"
)

// AppID 应用标识
var AppID = ""

// AppSecret 应用密钥
var AppSecret = ""

// GetAccessToken 获取 access_token
func GetAccessToken() string {
	body, _ := Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + AppID + "&secret=" + AppSecret)

	log.Println(body)

	type Resp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var resp Resp
	json.Unmarshal([]byte(body), &resp)

	return resp.AccessToken
}

// Add 添加草稿
func Add(title string, digest string, content string, thumbMediaID string) string {
	type Article struct {
		ThumbMediaID       string `json:"thumb_media_id"`
		Author             string `json:"author"`
		Title              string `json:"title"`
		ContentSourceURL   string `json:"content_source_url"`
		Content            string `json:"content"`
		Digest             string `json:"digest"`
		ShowCoverPic       int    `json:"show_cover_pic"`
		NeedOpenComment    int    `json:"need_open_comment"`
		OnlyFansCanComment int    `json:"only_fans_can_comment"`
	}

	type Req struct {
		Articles []Article `json:"articles"`
	}

	// 处理标题长度（64）
	if len([]rune(title)) > 64 {
		title = string([]rune(title)[:64])
	}
	title = string([]rune(title)[:FindInvert(title, "；")])

	// 处理概要长度（120）
	if len([]rune(digest)) > 110 {
		digest = string([]rune(digest)[:110])
		digest += "..."
	}

	body, _ := PostJSON("https://api.weixin.qq.com/cgi-bin/draft/add?access_token="+GetAccessToken(), Req{
		Articles: []Article{
			{
				ThumbMediaID:       thumbMediaID,
				Author:             "",
				Title:              title,
				ContentSourceURL:   "",
				Content:            content,
				Digest:             digest,
				ShowCoverPic:       0,
				NeedOpenComment:    0,
				OnlyFansCanComment: 0,
			},
		},
	})

	log.Println(body)

	type Resp struct {
		MediaID string `json:"media_id"`
	}

	var resp Resp
	json.Unmarshal([]byte(body), &resp)

	return resp.MediaID
}

// AddNews 添加图文素材
func AddNews(title string, digest string, content string, thumbMediaID string) string {
	type Article struct {
		ThumbMediaID       string `json:"thumb_media_id"`
		Author             string `json:"author"`
		Title              string `json:"title"`
		ContentSourceURL   string `json:"content_source_url"`
		Content            string `json:"content"`
		Digest             string `json:"digest"`
		ShowCoverPic       int    `json:"show_cover_pic"`
		NeedOpenComment    int    `json:"need_open_comment"`
		OnlyFansCanComment int    `json:"only_fans_can_comment"`
	}

	type Req struct {
		Articles []Article `json:"articles"`
	}

	// 处理标题长度（64）
	if len([]rune(title)) > 64 {
		title = string([]rune(title)[:64])
	}
	title = string([]rune(title)[:FindInvert(title, "；")])

	// 处理概要长度（120）
	if len([]rune(digest)) > 110 {
		digest = string([]rune(digest)[:110])
		digest += "..."
	}

	body, _ := PostJSON("https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="+GetAccessToken(), Req{
		Articles: []Article{
			{
				ThumbMediaID:       thumbMediaID,
				Author:             "",
				Title:              title,
				ContentSourceURL:   "",
				Content:            content,
				Digest:             digest,
				ShowCoverPic:       0,
				NeedOpenComment:    0,
				OnlyFansCanComment: 0,
			},
		},
	})

	log.Println(body)

	type Resp struct {
		MediaID string `json:"media_id"`
	}

	var resp Resp
	json.Unmarshal([]byte(body), &resp)

	return resp.MediaID
}

// SendAll 群发图文消息
func SendAll(mediaID string) {
	type Filter struct {
		IsToAll bool `json:"is_to_all"`
	}

	type MPNews struct {
		MediaID string `json:"media_id"`
	}

	type Req struct {
		Filter            Filter `json:"filter"`
		MPNews            MPNews `json:"mpnews"`
		MsgType           string `json:"msgtype"`
		SendIgnoreReprint int    `json:"send_ignore_reprint"`
	}

	body, _ := PostJSON("https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token="+GetAccessToken(), Req{
		Filter: Filter{
			IsToAll: true,
		},
		MPNews: MPNews{
			MediaID: mediaID,
		},
		MsgType:           "mpnews",
		SendIgnoreReprint: 0,
	})

	log.Println(body)
}
