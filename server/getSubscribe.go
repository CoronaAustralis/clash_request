package server

import (
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type Result struct {
	Data Data `json:"data"`
}

type Data struct {
	Subscribe_url string `json:"subscribe_url"`
	Token         string `json:"token"`
}

var SubsrciptionPath Result

func GetSubscribePath() bool {

	url := Config.RequestUrl
	method := "GET"

	client := resty.New()
	req := client.R()

	req.Header.Add("authority", "mojie.me")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("authorization", Config.RequestToken)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-language", "zh-CN")
	req.Header.Add("cookie", "dark_mode=0; crisp-client%2Fsession%2F733d8013-c930-4360-bb46-67934108eb32=session_36d7861c-bb0f-43b7-bd50-3218e39cfc74")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://mojie.me/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	res, err := req.Execute(method, url)
	if err != nil {
		log.Infoln("get subscribe path failed, retry! reason: ",err)
		return false
	}

	body := res.Body()

	err = jsoniter.Unmarshal(body, &SubsrciptionPath)

	if err != nil {
		log.Infoln("parse subscribe path failed, retry!")
		return false
	}

	log.Infof("subscribe url: %s", SubsrciptionPath.Data.Subscribe_url)
	return true
}
