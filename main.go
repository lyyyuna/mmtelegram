package main

import (
	"flag"
	"fmt"
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
	i := 1000
	for {
		j := 1
		for {
			targetUrl := fmt.Sprintf("https://img1.mm131.me/pic/%v/%v.jpg", i, j)
			_, body, _ := request.Get(targetUrl).
				Set("Referer", "https://m.mm131.net/xinggan").
				Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36").
				Set("Sec-Fetch-Mode", "no-cors").End()
			if err != nil {
				glog.Errorf("Failed to generate doc from body, the error is: %v", err)
				break
			}

			b := tgbotapi.FileBytes{
				Name:  "image.jpg",
				Bytes: []byte(body)}
			msg := tgbotapi.NewPhotoUpload(-1001257795899, b)
			msg.Caption = fmt.Sprintf("Sweet ~~ %v-%v", i, j)
			_, err = bot.Send(msg)
			if err != nil {
				glog.Errorf("Send image error, the error is: %v", err)
			}

			time.Sleep(time.Second * 1)

			j++
		}

		i++
	}
}
