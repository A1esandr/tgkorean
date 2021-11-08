package main

import (
	_ "embed"
	"flag"
	"log"
	"os"

	"github.com/A1esandr/tgkorean"
)

//go:embed token.txt
var embedToken string

//go:embed chat.txt
var embedChatID string

//go:embed letters.csv
var embedLetters []byte

func main() {
	flag.Parse()
	chatID := os.Getenv("CHAT_ID")
	token := os.Getenv("TOKEN")
	if token == "" {
		token = embedToken
	}
	if token == "" {
		log.Fatal("token is empty!")
	}
	if chatID == "" {
		chatID = embedChatID
	}
	if chatID == "" {
		log.Fatal("chat id is empty!")
	}
	tgkorean.New(tgkorean.AppParams{Token: token, ChatID: chatID}).Send(embedLetters)
}
