package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	channel, err := s.Channel(e.ChannelID)
	if err != nil {
		fmt.Println("Failed getting channel: ", err)
		return
	}

	fmt.Printf("%s said in channel %s: %s\n", e.Author.String(), channel.Name, e.Content)
}
