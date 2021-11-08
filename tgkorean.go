package tgkorean

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/A1esandr/tgbotapi"
)

type (
	app struct {
		token  string
		chatID interface{}
	}
	App interface {
		Start()
		Send(data []byte)
	}
	AppParams struct {
		Token  string
		ChatID interface{}
	}
)

func New(params AppParams) App {
	return &app{token: params.Token, chatID: params.ChatID}
}

func (a *app) Start() {
	letters := ReadCsv("letters.csv")
	a.send(letters)
}

func (a *app) Send(data []byte) {
	letters := ReadBytesCsv(data)
	a.send(letters)
}

func (a *app) send(letters [][]string) {
	lettersMap := make(map[string]string, len(letters))
	var sb strings.Builder
	for _, l := range letters {
		if len(l) != 3 {
			continue
		}
		lettersMap[l[0]] = l[1]
		sb.WriteString(l[0])
		sb.WriteString(" - ")
		sb.WriteString(l[1])
		sb.WriteString("\n")
	}
	bot, err := tgbotapi.New(a.token)
	if err != nil {
		log.Fatal(err)
	}
	positions := []int{0, 1, 2, 3}
	resp, err := bot.SendMessage(&tgbotapi.SendMessage{
		ChatID: a.chatID,
		Text:   sb.String(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
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
			if values[i] == value {
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
