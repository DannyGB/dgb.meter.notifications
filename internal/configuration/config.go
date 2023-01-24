package configuration

import (
	"os"
	"strconv"
)

type Configuration struct {
	SMTP_USERNAME       string
	SMTP_PASSWORD       string
	SMTP_HOST           string
	SMTP_PORT           string
	SMTP_FROM           string
	RECIPIENTS          string
	SUBJECT             string
	WEBSITE             string
	HTTP_PORT           string
	DISABLE_EMAIL       string
	TELEGRAM_BOT_TOKEN  string
	TELEGRAM_CHANNEL_ID int64
}

func NewConfig() Configuration {

	configuration := &Configuration{}
	configuration.SMTP_USERNAME = os.Getenv("METER_READINGS_NOTIFICATIONS_SMTP_USERNAME")
	configuration.SMTP_PASSWORD = os.Getenv("METER_READINGS_NOTIFICATIONS_SMTP_PASSWORD")
	configuration.SMTP_HOST = os.Getenv("METER_READINGS_NOTIFICATIONS_SMTP_HOST")
	configuration.SMTP_PORT = os.Getenv("METER_READINGS_NOTIFICATIONS_SMTP_PORT")
	configuration.SMTP_FROM = os.Getenv("METER_READINGS_NOTIFICATIONS_SMTP_FROM")
	configuration.RECIPIENTS = os.Getenv("METER_READINGS_NOTIFICATIONS_RECIPIENTS")
	configuration.WEBSITE = os.Getenv("METER_READINGS_NOTIFICATIONS_WEBSITE")
	configuration.SUBJECT = os.Getenv("METER_READINGS_NOTIFICATIONS_SUBJECT")
	configuration.HTTP_PORT = os.Getenv("METER_READINGS_NOTIFICATIONS_HTTP_PORT")
	configuration.DISABLE_EMAIL = os.Getenv("METER_READINGS_NOTIFICATIONS_DISABLE_EMAIL")
	configuration.TELEGRAM_BOT_TOKEN = os.Getenv("METER_READINGS_NOTIFICATIONS_TELEGRAM_BOT_TOKEN")
	configuration.TELEGRAM_CHANNEL_ID, _ = strconv.ParseInt(os.Getenv("METER_READINGS_NOTIFICATIONS_TELEGRAM_CHANNEL_ID"), 10, 64)

	return *configuration
}
