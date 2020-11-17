package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type JoinLeaveHandler struct{}

func NewJoinLeaveHandler() *JoinLeaveHandler {
	return &JoinLeaveHandler{}
}

func (h *JoinLeaveHandler) HandlerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("Failed getting guild object: ", err)
		return
	}

	fmt.Printf("Member %s joined the guild %s\n", e.Member.User.String(), guild.Name)
}

func (h *JoinLeaveHandler) HandlerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		fmt.Println("Failed getting guild object: ", err)
		return
	}

	fmt.Printf("Member %s left the guild %s\n", e.Member.User.String(), guild.Name)
}
