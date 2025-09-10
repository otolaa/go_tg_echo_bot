package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	botToken string
	botApi   string = "https://api.telegram.org/bot"
	botUrl   string

	suffix     string = "~"
	nbsp       string = " / "
	suffixLine string = strings.Repeat(suffix, 35)
	suffixEnd  string = "\033[0m\n"
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

func p(color int, str ...any) {
	suffixColor := "\033[3" + strconv.Itoa(color) + "m"
	fmt.Printf("%s%s%s", suffixColor, fmt.Sprint(str...), suffixEnd)
}
