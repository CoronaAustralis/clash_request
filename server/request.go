package server

import (
	"clash_request/utils"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Header http.Header
	Body   []byte
}

var SubsrciptionContent *Response = &Response{}

func GetSubscribeFile() bool {

	client := resty.New()
	req := client.R()

	u, err := url.Parse(SubsrciptionPath.Data.Subscribe_url)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Host", u.Host)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "ClashforWindows/0.20.9")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\"")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept-Language", "zh-CN")

	res, err := req.Execute("GET", SubsrciptionPath.Data.Subscribe_url)
	if err != nil {
		log.Infoln("get subscribe file failed, retry!")
		return false
	}

	body := res.Body()
	tempBody := string(res.Body())

	SubsrciptionContent.Header = res.Header()
	fileName := SubsrciptionContent.Header.Get("Content-Disposition")
	re := regexp.MustCompile(`filename\*=UTF-8''(.*)`)
	fileName = re.ReplaceAllString(fileName, "filename*=UTF-8''"+u.Host)
	SubsrciptionContent.Header.Set("Content-Disposition", fileName)

	tempBody = strings.Replace(tempBody, `allow-lan: false`, `allow-lan: true`, 1)
	tempBody = strings.Replace(tempBody, `127.0.0.1:9090`, `0.0.0.0:9090`, 1)

	SubsrciptionContent.Body = body

	body2 := []byte(tempBody)

	if !utils.Exists(Config.SubsrciptionPath) {
		dir := filepath.Dir(Config.SubsrciptionPath)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Panicln(err)
		}
		f, err := os.OpenFile(Config.SubsrciptionPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		defer f.Close()
		if err != nil {
			log.Panicln(err)
		}
		_, err = f.Write(body)
		if err != nil {
			log.Panicln(err)
		}
	} else {
		f, err := os.OpenFile(Config.SubsrciptionPath, os.O_WRONLY|os.O_TRUNC, 0777)
		defer f.Close()
		if err != nil {
			log.Panicln(err)
		}
		_, err = f.Write(body)
		if err != nil {
			log.Panicln(err)
		}
	}

	f, err := os.OpenFile("/app/clash_client/config.yaml", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		log.Panicln(err)
	}
	_, err = f.Write(body2)
	if err != nil {
		log.Panicln(err)
	}
	return true
}
