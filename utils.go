package main

import (
	"fmt"
	"regexp"
	"strings"
)

// checks if a message is a command intended for the bot. Any message that starts with the guild's prefix or a bot mention is a command
func isBotCommand(s, prefix string) (string, bool) {
	trim := strings.TrimSpace
	botMention := fmt.Sprintf("<@%v>", botID)
	br := regexp.MustCompile(fmt.Sprintf("^%v", botMention))

	if br.MatchString(s) {
		return trim(s[len(botMention):]), true
	}

	if strings.HasPrefix(s, prefix) {
		return trim(s[len(prefix):]), true
	}

	return "", false
}
