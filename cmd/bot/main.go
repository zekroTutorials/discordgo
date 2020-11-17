package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTutorials/discordgo/internal/config"
	"github.com/zekroTutorials/discordgo/internal/events"
)

func main() {
	const fileName = "./config/config.json"

	cfg, err := config.ParseConfigFromJSONFile(fileName)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers |
			discordgo.IntentsGuildMessages)

	registerEvents(s)

	if err = s.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	s.Close()
}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
}
