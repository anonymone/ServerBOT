package main

import (
	"./modules"
	"github.com/robfig/cron"
	"log"
)

func main() {
	path := "./config.json"
	data := modules.ConfigInformation{
		Receiver:   []string{"To Emails"},
		Sender:     "From",
		ServerHost: "Email server host",
		ServerPort: 25,
		APICredit:  "Email passwd",
		TaskSend:   "0 0 0 * * *",
		TaskCheck:  "0 0 * * * *",
		TaskTry:    "*/5 * * * * ?"}

	if data.ReadJSON(path) != nil {
		data.WriteJSON(path)
		log.Fatal("Please set the config.json file for your email service.")
	}
	timer := cron.New()
	l := modules.LocalInfo{IPs: []string{"0.0.0.0"}}
	e := modules.EmailInfo{
		LocalInfo:  &l,
		Receiver:   data.Receiver,
		Sender:     data.Sender,
		ServerHost: data.ServerHost,
		ServerPort: data.ServerPort,
		APICredit:  data.APICredit,
		TaskSend:   data.TaskSend,
		TaskCheck:  data.TaskCheck,
		TaskTry:    data.TaskTry,
		Message:    ""}

	timer.AddJob(data.TaskCheck, &l)
	timer.AddJob(data.TaskSend, e)

	timer.Start()
	defer timer.Stop()
	select {}
}
