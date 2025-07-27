package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// update request
func getUpdate(offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	restResponse := &ResResponse{}
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

// response
func sendResponse(update Update) error {
	botMessage := &BotMessage{
		ChatID: update.Message.Chat.ChatID,
		Text:   update.Message.Text,
	}

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	resp, err := http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}

// entry point
func main() {
	//~~~~
	fmt.Println("BOT_TOKEN → " + botToken)
	fmt.Println("BOT_URL → " + botUrl)
	//~~~~

	offset := 0
	for {
		updates, err := getUpdate(offset)
		if err != nil {
			log.Println("→", err.Error())
		}

		for _, update := range updates {
			err = sendResponse(update)
			if err != nil {
				log.Println("→", err.Error())
			}

			offset = update.UpdateID + 1
		}

		t := time.Now()
		formatted := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())

		//~~~~
		fmt.Println(updates)
		fmt.Println(formatted)
		//~~~~

		time.Sleep(2 * time.Second)
	}
}
