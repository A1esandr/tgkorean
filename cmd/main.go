package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/A1esandr/tgkorean"
)

var tokenFlag = flag.String("token", "", "Bot token")

func main() {
	flag.Parse()
	chatIDStr := os.Getenv("CHAT_ID")
	token := os.Getenv("TOKEN")
	if token == "" && tokenFlag != nil {
		token = *tokenFlag
	}
	if token == "" {
		log.Fatal("token is empty!")
	}
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tgkorean.New(tgkorean.AppParams{Token: token, ChatID: chatID}).Start()
}
