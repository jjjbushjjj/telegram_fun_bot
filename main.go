package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jjjbushjjj/telegram_fun_bot/courses"
	"github.com/jjjbushjjj/telegram_fun_bot/fun_pic"
	"github.com/jjjbushjjj/telegram_fun_bot/weather"
)

func printOutput(outs []byte) string {
	if len(outs) > 0 {
		return string(outs)
	}
	return ""
}

func ProcessCommand(command string) (out string) {
	switch command {
	case "help":
		out = "type /sayhi or /status.\n To show calendar /calendar.\n To get exchange rates type /courses.\n To get random Fun pic type /funpic.\n To get current weather type /weather"
	case "sayhi":
		out = "Hi :)"
	case "status":
		out = "I'm ok."
	case "calendar":
		cmd := exec.Command("calendar")
		output, _ := cmd.CombinedOutput()
		out = printOutput(output)
	case "courses":
		out = courses.GetCourses()
	case "funpic":
		out = fun_pic.GetFunPic()
	case "weather":
		out = weather.Get_weather()
	default:
		out = "I don't know that command"
	}
	return
}

func main() {

	// proxyUrl, err := url.Parse("socks5://794845572:xsn2GdDa@phobos.public.opennetwork.cc:1090")
	// proxyUrl, err := url.Parse("http://51.255.115.231:8080") // some proxy from internet This one is French
	proxyUrl, err := url.Parse("http://167.99.74.125:3128") // some proxy from internet This one is French
	//Create custom http client and transport for telegram access
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	tg_client := &http.Client{Transport: tr}

	// This will overwrite ALL http transports so all http requests will use proxy server
	// http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}

	resp, err := tg_client.Get("https://api.telegram.org/bot609231646:AAErrYiuODkTI6UgvruuIBzSmydsqKku59U/getMe")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	// try to use api
	bot, err := tgbotapi.NewBotAPIWithClient("609231646:AAErrYiuODkTI6UgvruuIBzSmydsqKku59U", tg_client)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	var msg tgbotapi.MessageConfig
	// var doc tgbotapi.DocumentConfig
	// msg = tgbotapi.NewMessageToChannel("@test_zabbix_bot", "Hello everyone!")
	// doc = tgbotapi.NewDocumentUpload(794845572, "IMG_2019-01-19_130949.jpg")
	// bot.Send(msg)
	// bot.Send(doc)
	// msg = tgbotapi.NewMessage(794845572, "https://cdn.jpg.wtf/futurico/33/5c/1455974066-335c822a23665a3d7f74727a7a15f3e3.jpeg")
	// bot.Send(msg)
	for update := range updates {
		msg.Text = ""

		switch {

		case update.ChannelPost != nil:
			log.Printf("Got channel post [%s]", update.ChannelPost.Text)
			if update.ChannelPost.IsCommand() {
				msg = tgbotapi.NewMessage(update.ChannelPost.Chat.ID, "")
				cmd := update.ChannelPost.Command()
				// if cmd == "Maxbithday" {
				// 	msg_photo := tgbotapi.NewSetChatPhotoUpload(update.ChannelPost.Chat.ID, "IMG_2019-01-19_130949.jpg")
				// 	bot.SetChatPhoto(msg_photo)
				// 	continue
				// }
				msg.Text = ProcessCommand(cmd)
			}

		case update.Message != nil:
			// This stuff work when you send messages directly to bot
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			if update.Message.From.UserName != "jjjbushjjj" {
				log.Printf("This must be rejected Got message from  [%s] unknown user", update.Message.From.UserName)
				msg.Text = "Sorry i would accept commands only from my creator you are not him"
			}

			if update.Message.IsCommand() {
				cmd := update.Message.Command()
				msg.Text = ProcessCommand(cmd)
			}

		}
		log.Printf("Sending [%s]", msg.Text)
		if msg.Text != "" {
			bot.Send(msg)
		}
	}
}
