package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-chat-bot/src/clients/elasticsearch"
	"github.com/ssoql/faq-chat-bot/src/datasources/faqs_db"
	"github.com/ssoql/faq-chat-bot/src/mocks"
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestCreateFaqConcurrentInvalidData(t *testing.T) {
	input := &faqs.FaqCreateInput{}
	outChan := make(chan faqs.CreateFaqsResult)

	service := faqService{}
	go service.CreateFaqConcurrent(input, outChan)

	result := <-outChan
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid question", result.Error.Message())
}

func TestCreateFaqConcurrentOperationError(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "xxx", Answer: "yyy"}
	outChan := make(chan faqs.CreateFaqsResult)

	db, _ := mocks.GetDbMock()
	faqs_db.Client = db

	service := faqService{}
	go service.CreateFaqConcurrent(input, outChan)

	result := <-outChan
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusInternalServerError, result.Error.Status())
	assert.EqualValues(t, "error when tying to save faq", result.Error.Message())
}

func TestCreateFaqConcurrentNoError(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "xxx", Answer: "yyy"}
	outChan := make(chan faqs.CreateFaqsResult)

	db, mock := mocks.GetDbMock()

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `faqs`.*").
		WithArgs("f561aaf6ef0bf14d4208bb46a4ccb3ad",
			input.Question, input.Answer,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	faqs_db.Client = db
	elasticsearch.EsClient = &mocks.EsClientMock{}

	service := faqService{}
	go service.CreateFaqConcurrent(input, outChan)

	result := <-outChan
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 1, result.Response.Id)
	assert.EqualValues(t, "xxx", result.Response.Question)
	assert.EqualValues(t, "yyy", result.Response.Answer)
}

func TestHandleFaqCreateInputs(t *testing.T) {
	inChan := make(chan faqs.CreateFaqsResult)
	outChan := make(chan faqs.CreateFaqsResponse)
	var wg sync.WaitGroup

	service := faqService{}
	go service.handleFaqCreateInputs(&wg, inChan, outChan)

	wg.Add(1)
	go func() {
		inChan <- faqs.CreateFaqsResult{Error: api_errors.NewBadRequestError("invalid question")}
	}()

	wg.Wait()
	close(inChan)

	result := <-outChan
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid question", result.Results[0].Error.Message())
}
