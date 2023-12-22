// this file works is to reading all data inside config.json
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config) //fr unmarshal always pass reference

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
