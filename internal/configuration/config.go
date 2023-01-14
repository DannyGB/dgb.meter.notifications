package configuration

import "os"

type Configuration struct {
	SMTP_USERNAME string
	SMTP_PASSWORD string
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_FROM     string
	RECIPIENTS    string
	SUBJECT       string
	WEBSITE       string
	HTTP_PORT     string
}

func NewConfig() Configuration {

	configuration := &Configuration{}
	configuration.SMTP_USERNAME = os.Getenv("METER_READINGS_SMTP_USERNAME")
	configuration.SMTP_PASSWORD = os.Getenv("METER_READINGS_SMTP_PASSWORD")
	configuration.SMTP_HOST = os.Getenv("METER_READINGS_SMTP_HOST")
	configuration.SMTP_PORT = os.Getenv("METER_READINGS_SMTP_PORT")
	configuration.SMTP_FROM = os.Getenv("METER_READINGS_SMTP_FROM")
	configuration.RECIPIENTS = os.Getenv("METER_READINGS_RECIPIENTS")
	configuration.WEBSITE = os.Getenv("METER_READINGS_WEBSITE")
	configuration.SUBJECT = os.Getenv("METER_READINGS_SUBJECT")
	configuration.HTTP_PORT = os.Getenv("METER_READINGS_HTTP_PORT")

	return *configuration
}
