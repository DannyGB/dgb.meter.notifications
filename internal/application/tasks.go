package application

import (
	"fmt"
	"log"
	"strings"
	"time"

	"net/smtp"

	"dgb.meter.notifications/internal/configuration"
	"github.com/go-co-op/gocron"
)

func CreateTasks() {
	gc := gocron.NewScheduler(time.UTC)

	//gc.Every(10).Do(notificationTask) // for testing only
	gc.Every(1).MonthLastDay().Do(notificationTask)
	gc.StartBlocking()
}

func notificationTask() {

	conf := configuration.NewConfig()

	fmt.Println("Beginning email reminder task.")

	auth := smtp.PlainAuth("", conf.SMTP_USERNAME, conf.SMTP_PASSWORD, conf.SMTP_HOST)

	to := strings.Split(conf.RECIPIENTS, ",")

	msg := []byte(fmt.Sprintf("To: %s\r\n"+

		"Subject: %s\r\n"+

		"\r\n"+

		"You should take an electric meter reading and upload it to %s\r\n", conf.RECIPIENTS, conf.SUBJECT, conf.WEBSITE))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", conf.SMTP_HOST, conf.SMTP_PORT), auth, conf.SMTP_FROM, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ending email reminder task.")
}
