package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	dg "github.com/bwmarrin/discordgo"
)

var (
	botID string
)

const (
	creatorID = "399951813237014528"
)

func main() {

	// Get the bot token from the environment variables
	tk := os.Getenv("TOKEN")
	if tk == "" {
		fmt.Println("VAR TOKEN NOT SET. Please create a discord application as bot user and use the generated token.")
		return
	}

	// prefix the token with bot and use it to create a new bot session
	bot, err := dg.New(fmt.Sprintf("Bot %v", tk))

	if err != nil {
		fmt.Println("Error in bot", err)
		return
	}

	// initalize the global variable botID
	if u, err := bot.User("@me"); err != nil {
		fmt.Println("Error in initialising the botID", err)
	} else {
		botID = u.ID
	}

	// add the handlers:

	//   ready: updates status when the bot comes live
	bot.AddHandler(ready)
	//   ping: initial pingpong tests
	bot.AddHandler(ping)
	// commands: register all the available bot command
	bot.AddHandler(commands)

	// opens and closes the gateway
	err = bot.Open()
	defer bot.Close()

	fmt.Println("Bot is online. Please, use Ctrl-C to exit.")
	// Wait for a Signal to close the bot
	scn := make(chan os.Signal, 1)
	signal.Notify(scn, syscall.SIGINT, os.Interrupt, os.Kill)
	<-scn

	fmt.Println("Bot is going offline.")
}
