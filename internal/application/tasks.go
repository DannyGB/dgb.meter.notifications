package application

import (
	"fmt"
	"log"
	"strings"
	"time"

	"net/smtp"

	"dgb/meter.notifications/internal/configuration"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateTasks(conf configuration.Configuration) {

	gc := gocron.NewScheduler(time.UTC)

	startTelegramBot(conf, gc)

	if conf.DISABLE_EMAIL == "true" {
		return
	}

	sendTestEmail(conf)

	gc.Every(1).MonthLastDay().Do(func() {
		notificationTask(conf)
	})

	gc.StartAsync()
}

func notificationTask(conf configuration.Configuration) {

	fmt.Println("Beginning email reminder task.")

	auth := smtp.PlainAuth("", conf.SMTP_USERNAME, conf.SMTP_PASSWORD, conf.SMTP_HOST)

	to := strings.Split(conf.RECIPIENTS, ",")

	msg := []byte(fmt.Sprintf("To: %s\r\n"+

		"Subject: %s\r\n"+

		"\r\n"+

		"You should take an electric meter reading and upload it to %s\r\n or upload via the Telegram bot @DannygbReadingsBot by typing '/add n:<reading> d:<reading>'", conf.RECIPIENTS, conf.SUBJECT, conf.WEBSITE))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", conf.SMTP_HOST, conf.SMTP_PORT), auth, conf.SMTP_FROM, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ending email reminder task.")
}

func sendTelegramReminder(conf configuration.Configuration, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(conf.TELEGRAM_CHANNEL_ID, fmt.Sprintf("You should take an electric meter reading and upload it to %s\r\n or or upload via the Telegram bot @DannygbReadingsBot by typing '/add n:<reading> d:<reading>'", conf.WEBSITE))
	bot.Send(msg)
}

func startTelegramBot(conf configuration.Configuration, gc *gocron.Scheduler) {
	bot, err := tgbotapi.NewBotAPI(conf.TELEGRAM_BOT_TOKEN)

	if err != nil {
		log.Panic(err)
	}

	gc.Every(1).MonthLastDay().Do(func() {
		sendTelegramReminder(conf, bot)
	})
}

func sendTestEmail(conf configuration.Configuration) {
	notificationTask(conf)
}
