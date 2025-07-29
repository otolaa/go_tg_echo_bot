package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	botToken string
	botApi   string = "https://api.telegram.org/bot"
	botUrl   string
)

func init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "TOKEN=") {
			botToken = strings.ReplaceAll(s, "TOKEN=", "")
		}
	}

	botUrl = botApi + botToken

	err = dellWebhook()
	if err != nil {
		panic(err)
	}
}
