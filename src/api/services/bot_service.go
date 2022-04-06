package services

import "fmt"

const (
	replyPrefix = "Chat Bot: "
	greetingMsg = replyPrefix + "Hello :)"
	dontKnowMsg = replyPrefix + "Sorry i don't know the answer :("
)

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
	return greetingMsg
}

func (b *botService) Reply(msg string) string {
	// todo implement search in elastic
	faq, err := FaqService.SearchFaq(msg)
	if err != nil {
		fmt.Printf("%s", err.Message())
		return dontKnowMsg
	}
	return replyPrefix + faq[0].Answer
}
