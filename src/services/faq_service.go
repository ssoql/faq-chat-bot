package services

import (
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/models/queries"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"github.com/ssoql/faq-chat-bot/src/utils/crypto_utils"
	"log"
	"net/http"
	"strings"
	"sync"
)

var FaqService FaqServiceInterface = &faqService{}

type faqService struct{}

type FaqServiceInterface interface {
	CreateFaq(*faqs.FaqCreateInput) (*faqs.Faq, api_errors.ApiError)
	CreateFaqs([]faqs.FaqCreateInput) *faqs.CreateFaqsResponse
	UpdateFaq(int64, *faqs.FaqUpdateInput) (*faqs.Faq, api_errors.ApiError)
	DeleteFaq(int64) api_errors.ApiError
	GetFaq(int64) (*faqs.Faq, api_errors.ApiError)
	SearchFaq(string) (faqs.FaqDocuments, api_errors.ApiError)
	InitializeDemoFaqs()
}

func (s *faqService) CreateFaq(input *faqs.FaqCreateInput) (*faqs.Faq, api_errors.ApiError) {
	var faq = faqs.Faq{Question: input.Question, Answer: input.Answer}

	if err := faq.Validate(); err != nil {
		return nil, err
	}
	if err := faq.Save(); err != nil {
		return nil, err
	}
	faqDoc := faqs.NewDocFromFaq(&faq)
	if err := faqDoc.Save(); err != nil {
		return nil, err
	}
	return &faq, nil
}

func (s *faqService) CreateFaqs(requests []faqs.FaqCreateInput) *faqs.CreateFaqsResponse {
	inChan := make(chan faqs.CreateFaqsResult)
	outChan := make(chan faqs.CreateFaqsResponse)
	defer close(outChan)

	var wg sync.WaitGroup
	go s.handleFaqCreateInputs(&wg, inChan, outChan)

	for _, request := range requests {
		wg.Add(1)
		go s.CreateFaqConcurrent(&request, inChan)
	}
	wg.Wait()
	close(inChan)

	result := <-outChan
	succeeded := 0
	for _, res := range result.Results {
		if res.Response != nil {
			succeeded++
		}
	}

	if succeeded == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if succeeded != len(requests) {
		result.StatusCode = http.StatusPartialContent
	} else {
		result.StatusCode = http.StatusCreated
	}

	return &result
}

func (s *faqService) handleFaqCreateInputs(wg *sync.WaitGroup, inChan chan faqs.CreateFaqsResult, outChan chan faqs.CreateFaqsResponse) {
	var results faqs.CreateFaqsResponse
	for item := range inChan {
		faqResult := faqs.CreateFaqsResult{
			Response: item.Response,
			Error:    item.Error,
		}
		results.Results = append(results.Results, faqResult)
		wg.Done()
	}
	// processing only if input channel is closed
	outChan <- results
}

func (s *faqService) CreateFaqConcurrent(input *faqs.FaqCreateInput, outChan chan faqs.CreateFaqsResult) {
	result, err := s.CreateFaq(input)
	if err != nil {
		outChan <- faqs.CreateFaqsResult{Error: err}
		return
	}
	outChan <- faqs.CreateFaqsResult{Response: result}
}

func (s *faqService) UpdateFaq(id int64, input *faqs.FaqUpdateInput) (*faqs.Faq, api_errors.ApiError) {
	var faq = &faqs.Faq{Id: id}
	if err := faq.Get(); err != nil {
		return nil, err
	}

	faq.Question = input.Question
	faq.Answer = input.Answer

	if err := faq.Validate(); err != nil {
		return nil, err
	}
	faq.UniqHash = crypto_utils.GetMd5(strings.ToLower(faq.Question))
	if err := faq.Update(); err != nil {
		return nil, err
	}

	return faq, nil
}

func (s *faqService) DeleteFaq(id int64) api_errors.ApiError {
	var faq = &faqs.Faq{Id: id}
	if err := faq.Get(); err != nil {
		return err
	}
	if err := faq.Delete(); err != nil {
		return err
	}

	query := &queries.EsQuery{
		Delete: []queries.FieldValue{{Field: "id", Value: faq.Id}},
	}
	faqDoc := &faqs.FaqDocument{}

	if err := faqDoc.Delete(query); err != nil {
		return err
	}
	return nil
}

func (s *faqService) GetFaq(id int64) (*faqs.Faq, api_errors.ApiError) {
	var faq = &faqs.Faq{Id: id}
	if err := faq.Get(); err != nil {
		return nil, err
	}
	return faq, nil
}

func (s *faqService) SearchFaq(q string) (faqs.FaqDocuments, api_errors.ApiError) {
	query := &queries.EsQuery{
		FullText: []queries.FieldValue{{Field: "question", Value: q}},
	}
	faqDoc := &faqs.FaqDocument{}

	return faqDoc.Search(query)
}

func (s *faqService) InitializeDemoFaqs() {
	demoData := map[string]string{
		"Hi!":                             "Hello again :)",
		"What is your name?":              "My name is Bot. Chat Bot.",
		"Is Earth flat?":                  "No it isn't.",
		"How many languages do you know?": "Unfortunately i know only one language - english.",
		"How are you?":                    "I'm great :D",
	}

	for qust, ansr := range demoData {
		input := &faqs.FaqCreateInput{Question: qust, Answer: ansr}
		_, err := s.CreateFaq(input)
		if err != nil {
			log.Println(err.Message())
		}
	}
}
