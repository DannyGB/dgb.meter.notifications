package main

import (
	"dgb.meter.notifications/internal/application"
	"dgb.meter.notifications/internal/configuration"
)

func main() {
	conf := configuration.NewConfig()
	application.SendTestEmail(conf)
	application.CreateTasks(conf)
	application.HandleRequests(conf)
}
