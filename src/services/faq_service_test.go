package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/olivere/elastic"
	"github.com/ssoql/faq-chat-bot/src/clients/elasticsearch"
	"github.com/ssoql/faq-chat-bot/src/datasources/faqs_db"
	"github.com/ssoql/faq-chat-bot/src/mocks"
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestCreateFaqInvalidData(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "", Answer: "yyy"}
	result, err := FaqService.CreateFaq(input)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid question", err.Message())
}

func TestCreateFaqDbError(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "xxx", Answer: "yyy"}
	db, _ := mocks.GetDbMock()
	faqs_db.Client = db

	result, err := FaqService.CreateFaq(input)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error when tying to save faq", err.Message())
}

func TestCreateFaqEsError(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "xxx", Answer: "yyy"}
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
	elasticsearch.EsClient = &mocks.EsClientMock{
		IndexCallback: func() (*elastic.IndexResponse, error) {
			return nil, api_errors.NewBadRequestError("There is no active ES node")
		}}

	result, err := FaqService.CreateFaq(input)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error when trying to save faq in ES", err.Message())
}

func TestCreateFaqNoError(t *testing.T) {
	input := &faqs.FaqCreateInput{Question: "xxx", Answer: "yyy"}
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
	elasticsearch.EsClient = &mocks.EsClientMock{
		IndexCallback: func() (*elastic.IndexResponse, error) {
			return &elastic.IndexResponse{}, nil
		}}

	result, err := FaqService.CreateFaq(input)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 1, result.Id)
	assert.EqualValues(t, "xxx", result.Question)
	assert.EqualValues(t, "yyy", result.Answer)
}

func TestGetFaqDbError(t *testing.T) {
	db, _ := mocks.GetDbMock()
	faqs_db.Client = db

	result, err := FaqService.GetFaq(0)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error when tying to fetch faq", err.Message())
}

func TestGetFaqNoError(t *testing.T) {
	db, mock := mocks.GetDbMock()

	mockedRow := sqlmock.NewRows([]string{"id", "question", "answer", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "xxx", "yyy", time.Now(), time.Now(), nil)
	mock.ExpectQuery("SELECT").
		WillReturnRows(mockedRow)

	faqs_db.Client = db

	result, err := FaqService.GetFaq(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 1, result.Id)
	assert.EqualValues(t, "xxx", result.Question)
	assert.EqualValues(t, "yyy", result.Answer)
}

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
	elasticsearch.EsClient = &mocks.EsClientMock{
		IndexCallback: func() (*elastic.IndexResponse, error) {
			return &elastic.IndexResponse{}, nil
		}}

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
