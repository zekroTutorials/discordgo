package commands

import "github.com/bwmarrin/discordgo"

// [prefix][invoke/alias] [1st arg] [2nd arg] [3rd agrg]
// ;;kick zekro ist dumm

type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Args    []string
	Handler *CommandHandler
}
