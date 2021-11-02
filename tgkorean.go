package tgkorean

type (
	app struct {
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
	return &app{}
}

func (a *app) Start() {

}
