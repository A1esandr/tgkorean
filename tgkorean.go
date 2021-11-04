package tgkorean

import "fmt"

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
	for key, value := range lettersMap {
		fmt.Println(key)
		fmt.Println(value)
	}
}
