package services

const (
	greetingMsg = "Hello :)"
	dontKnowMsg = "Sorry i don't know the answer :("
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
	faq, err := FaqService.SearchFaq(msg)
	if err != nil {
		return dontKnowMsg
	}
	return faq[0].Answer
}
