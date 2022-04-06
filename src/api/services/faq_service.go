package services

import (
	"github.com/ssoql/faq-chat-bot/src/api/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/api/models/queries"
	"github.com/ssoql/faq-chat-bot/src/api/utils/api_errors"
	"github.com/ssoql/faq-chat-bot/src/api/utils/crypto_utils"
	"strings"
)

var FaqService FaqServiceInterface = &faqService{}

type faqService struct{}

type FaqServiceInterface interface {
	CreateFaq(*faqs.FaqCreateInput) (*faqs.Faq, api_errors.ApiError)
	UpdateFaq(int64, *faqs.FaqUpdateInput) (*faqs.Faq, api_errors.ApiError)
	DeleteFaq(int64) api_errors.ApiError
	GetFaq(int64) (*faqs.Faq, api_errors.ApiError)
	SearchFaq(string) (faqs.FaqDocuments, api_errors.ApiError)
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

func MigrateFaqs() {

}
