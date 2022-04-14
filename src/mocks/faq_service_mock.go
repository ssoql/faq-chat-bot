package mocks

import (
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
)

type FaqServiceMock struct {
	FaqErrCallback func() (*faqs.Faq, api_errors.ApiError)
	ErrCallback    func() api_errors.ApiError
	FaqsCallback   func() *faqs.CreateFaqsResponse
}

func (s *FaqServiceMock) CreateFaq(input *faqs.FaqCreateInput) (*faqs.Faq, api_errors.ApiError) {
	return s.FaqErrCallback()
}

func (s *FaqServiceMock) CreateFaqs(requests []faqs.FaqCreateInput) *faqs.CreateFaqsResponse {
	return s.FaqsCallback()
}

func (s *FaqServiceMock) UpdateFaq(id int64, input *faqs.FaqUpdateInput) (*faqs.Faq, api_errors.ApiError) {
	return s.FaqErrCallback()
}

func (s *FaqServiceMock) DeleteFaq(int64) api_errors.ApiError {
	return s.ErrCallback()
}

func (s *FaqServiceMock) GetFaq(int64) (*faqs.Faq, api_errors.ApiError) {
	return s.FaqErrCallback()
}

func (s *FaqServiceMock) SearchFaq(string) (faqs.FaqDocuments, api_errors.ApiError) { return nil, nil }
func (s *FaqServiceMock) InitializeDemoFaqs()                                       {}
