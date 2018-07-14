package main

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func ready(s *dg.Session, e *dg.Ready) {
	// set the playing status
	s.UpdateStatus(0, "Cryptocurreny")
}

func ping(s *dg.Session, m *dg.MessageCreate) {
	// only replies to pings from the creator
	if m.Content == "!ping" && m.Author.ID == creatorID {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}

func commands(s *dg.Session, m *dg.MessageCreate) {
	// the command handler for detecting and executing the bot commands. Its a just a giant swicth case
	msg, flag := isBotCommand(m.Content, ">")
	if !flag {
		return
	}

	cmd := strings.Split(msg, " ")

	// switch case for all the commands
	switch cmd[0] {

	case "info":
		s.ChannelMessageSend(m.ChannelID, "info")

	default:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%v` is not a valid command", cmd[0]))
	}

}
