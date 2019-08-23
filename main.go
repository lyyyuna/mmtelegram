package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/glog"
	"github.com/parnurzeal/gorequest"
)

func main() {
	token := flag.String("t", "6882", "token")
	flag.Parse()
	defer glog.Flush()

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		glog.Exitf("error: %v", err)
	}

	bot.Debug = true
	request := gorequest.New()
	i := 1500
	for {
		j := 1
		for {
			targetUrl := fmt.Sprintf("https://img1.mm131.me/pic/%v/%v.jpg", i, j)
			resp, body, _ := request.Get(targetUrl).
				Set("Referer", "https://m.mm131.net/xinggan").
				Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36").
				Set("Sec-Fetch-Mode", "no-cors").End()

			if resp.StatusCode == http.StatusNotFound {
				glog.Errorf("404 not found, %v - %v", i, j)
				break
			}

			b := tgbotapi.FileBytes{
				Name:  "image.jpg",
				Bytes: []byte(body)}
			msg := tgbotapi.NewPhotoUpload(-1001257795899, b)
			msg.Caption = fmt.Sprintf("Sweet ~~ %v-%v", i, j)
			_, err := bot.Send(msg)
			if err != nil {
				glog.Errorf("Send image error, the error is: %v", err)
				glog.Error(resp.StatusCode)
				glog.Errorf("%v - %v", i, j)
			}

			time.Sleep(time.Second * 30)

			j++
		}

		i++
	}
}
