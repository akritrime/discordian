package main

import (
	dg "github.com/bwmarrin/discordgo"
)

func ready(s *dg.Session, e *dg.Ready) {
	// set the playing status
	s.UpdateStatus(0, "Cryptocurreny")
}

func ping(s *dg.Session, m *dg.MessageCreate) {
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
