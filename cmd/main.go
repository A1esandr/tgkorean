package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/A1esandr/tgkorean"
)

//go:embed token.txt
var embedToken string

//go:embed chat.txt
var embedChatID string

func main() {
	flag.Parse()
	chatIDStr := os.Getenv("CHAT_ID")
	token := os.Getenv("TOKEN")
	if token == "" {
		token = embedToken
	}
	if token == "" {
		log.Fatal("token is empty!")
	}
	if chatIDStr == "" {
		chatIDStr = embedChatID
	}
	if chatIDStr == "" {
		log.Fatal("chat id is empty!")
	}
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tgkorean.New(tgkorean.AppParams{Token: token, ChatID: chatID}).Start()
}
