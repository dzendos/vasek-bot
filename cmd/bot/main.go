package main

import (
	"log"
	"time"

	"vasek-bot/internal/clients/tg"
	"vasek-bot/internal/config"
	"vasek-bot/internal/model/messages"
	"vasek-bot/internal/model/state"

	"github.com/jasonlvhit/gocron"
)

func task() {
	hours := time.Now().Hour()

	if hours == 18 {
		text := "Ва-ха-ха новые Васьки на подходе\nА вот и желающие стать Васьком сегодня!\n"
		if !state.RomkaState.WasSent {
			text += state.RomkaState.Name + "\n"
		}
		if !state.EvgState.WasSent {
			text += state.EvgState.Name + "\n"
		}
		if !state.MomState.WasSent {
			text += state.MomState.Name + "\n"
		}
		tgClient.SendMessage(text, -657739139)
	} else if hours == 0 {
		if state.MomState.WasSent == true {
			state.MomState.Score++
		}
		if state.EvgState.WasSent == true {
			state.EvgState.Score++
		}
		if state.RomkaState.WasSent == true {
			state.RomkaState.Score++
		}
		state.MomState.WasSent = false
		state.EvgState.WasSent = false
		state.RomkaState.WasSent = false
	}
}

func action() {
	s := gocron.NewScheduler()
	s.Every(1).Hours().Do(task)
	<-s.Start()
}

var tgClient *tg.Client

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err = tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	go action()

	tgClient.ListenUpdates(msgModel)
}
