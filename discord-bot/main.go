package main

import(
	"fmt"
	"github.com/helloabhii/go-only/discord-bot/bot"
	"github.com/helloabhii/go-only/discord-bot/config"
)
func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	bot.Start()
	<-make(chan struct{})
	return
}
