package services

type botService struct{}

type BotServiceInterface interface {
	Greeting() string
	Reply(string) string
}

var BotService BotServiceInterface

func init() {
	BotService = &botService{}
}

func (b *botService) Greeting() string {
	return "Hello"
}

func (b *botService) Reply(msg string) string {
	var result string
	result = "Chat Bot: " + msg
	return result
}
