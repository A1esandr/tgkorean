package tgkorean

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/A1esandr/tgbotapi"
)

type (
	app struct {
		token  string
		chatID int64
	}
	App interface {
		Start()
	}
	AppParams struct {
		Token  string
		ChatID int64
	}
)

func New(params AppParams) App {
	return &app{token: params.Token, chatID: params.ChatID}
}

func (a *app) Start() {
	letters := ReadCsv("letters.csv")
	lettersMap := make(map[string]string, len(letters))
	for _, l := range letters {
		if len(l) != 3 {
			continue
		}
		lettersMap[l[0]] = l[1]
	}
	bot, err := tgbotapi.New(a.token)
	if err != nil {
		log.Fatal(err)
	}
	positions := []int{0, 1, 2, 3}
	for key, value := range lettersMap {
		fmt.Println(key)
		fmt.Println(value)

		answers := make([]string, 4)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(positions),
			func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })
		values := make([]string, 0, len(lettersMap))
		for k, v := range lettersMap {
			if k == key {
				continue
			}
			values = append(values, v)
		}
		rand.Shuffle(len(values),
			func(i, j int) { values[i], values[j] = values[j], values[i] })
		for i := 0; i < 4; i++ {
			if i == positions[0] {
				answers[i] = value
				continue
			}
			answers[i] = values[i]
		}
		resp, err := bot.SendPoll(&tgbotapi.SendPoll{
			ChatID:          a.chatID,
			Question:        key,
			Options:         answers,
			Type:            "quiz",
			CorrectOptionID: positions[0],
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))

		time.Sleep(time.Duration(5) * time.Second)
	}
}
