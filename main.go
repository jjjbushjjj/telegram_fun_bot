package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jjjbushjjj/telegram_fun_bot/ansible_snippets"
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

func ProcessCommand(command string, args string) (out string) {
	switch command {
	case "help":
		Conf.Parse_mode = ""
		out = "type /sayhi or /status.\n To show calendar /calendar.\n To get exchange rates type /courses.\n To get random Fun pic type /funpic.\n To get current weather type /weather\n To get doc for ansible module type: /ansible_module <modulename>"
	case "sayhi":
		out = "Hi :)"
	case "status":
		out = "I'm ok."
	case "calendar":
		Conf.Parse_mode = ""
		cmd := exec.Command("calendar")
		output, _ := cmd.CombinedOutput()
		out = printOutput(output)
	case "courses":
		out = courses.GetCourses()
	case "funpic":
		Conf.Parse_mode = "HTML"
		out = fun_pic.GetFunPic()
	case "weather":
		Conf.Parse_mode = "HTML"
		out = weather.Get_weather()
	case "ansible_module":
		Conf.Parse_mode = ""
		out = ansible_snippets.Get_snippet(args)
	default:
		out = "I don't know that command"
	}
	return
}

type Configuration struct {
	Api_token  string
	Parse_mode string
}

var Conf = Configuration{}

func main() {

	// Read config file json based
	file, err := os.Open("fun_bot.json")
	if err != nil {
		log.Panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(Conf.Api_token)
	fmt.Println(Conf.Parse_mode)

	file.Close()
	// try to use api
	bot, err := tgbotapi.NewBotAPI(Conf.Api_token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://tgbot.bushuev.xyz/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("127.0.0.1:8080", nil)

	var msg tgbotapi.MessageConfig
	for update := range updates {
		msg.Text = ""

		switch {

		case update.ChannelPost != nil:
			log.Printf("Got channel post [%s]", update.ChannelPost.Text)
			if update.ChannelPost.IsCommand() {
				msg = tgbotapi.NewMessage(update.ChannelPost.Chat.ID, "")
				cmd := update.ChannelPost.Command()
				args := update.ChannelPost.CommandArguments()
				msg.Text = ProcessCommand(cmd, args)
			}

		case update.Message != nil:
			// This stuff work when you send messages directly to bot
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if update.Message.IsCommand() {
				cmd := update.Message.Command()
				args := update.Message.CommandArguments()
				msg.Text = ProcessCommand(cmd, args)
			}

		}
		msg.ParseMode = Conf.Parse_mode
		log.Printf("Sending [%s], parse_mode is: (%s)", msg.Text, msg.ParseMode)
		if msg.Text != "" {
			bot.Send(msg)
		}
	}
}
