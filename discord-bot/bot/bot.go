package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/helloabhii/go-only/discord-bot/config"
)

// this file main work is to send messges as bot on server
var BotID string
var goBot *discordgo.Session //when we create a new session  store in gobot

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token) //here we using New function from discordgo package
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me") //create bot as a user
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler) //this help us to read and write the message in channel

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return //do nothing if both are same
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong") //we don't have  to do anything with these values so  capture these in empty var
	}
}
