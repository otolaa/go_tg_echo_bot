package main

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type ResResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type DeleteWebhook struct {
	Description string `json:"description"`
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
}
