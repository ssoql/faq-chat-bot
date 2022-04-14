package faq

import (
	"encoding/json"
	"errors"
	"github.com/ssoql/faq-chat-bot/src/mocks"
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/services"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"github.com/ssoql/faq-chat-bot/src/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateInvalidJsonData(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq", strings.NewReader(`{"}`))
	c, _ := test_utils.GetContextMock(request, response)

	Create(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json data", apiErr.Message())
}

func TestCreateOperationError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq", strings.NewReader(`{"question":"   ","answer":"   "}`))
	c, _ := test_utils.GetContextMock(request, response)

	Create(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid question", apiErr.Message())
}

func TestCreateNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq", strings.NewReader(`{"question":"   ","answer":"   "}`))
	c, _ := test_utils.GetContextMock(request, response)

	services.FaqService = &mocks.FaqServiceMock{
		FaqErrCallback: func() (*faqs.Faq, api_errors.ApiError) {
			return &faqs.Faq{Id: 1, Question: "x", Answer: "y"}, nil
		},
	}

	Create(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result faqs.Faq
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result.Id)
	assert.EqualValues(t, "x", result.Question)
	assert.EqualValues(t, "y", result.Answer)
}

func TestUpdateInvalidJsonData(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPatch, "/faq/1", strings.NewReader(``))
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	Update(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json data", apiErr.Message())
}

func TestUpdateOperationError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq/1", strings.NewReader(`{"question":"   ","answer":"   "}`))
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		FaqErrCallback: func() (*faqs.Faq, api_errors.ApiError) {
			return nil, api_errors.NewBadRequestError("invalid question")
		},
	}

	Update(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid question", apiErr.Message())
}

func TestUpdateInvalidId(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq/1", strings.NewReader(`{"question":"x","answer":"y"}`))
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "x")

	Update(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "id must be a number", apiErr.Message())
}

func TestUpdateNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq/1", strings.NewReader(`{"question":"x","answer":"y"}`))
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		FaqErrCallback: func() (*faqs.Faq, api_errors.ApiError) {
			return &faqs.Faq{Id: 1, Question: "x", Answer: "y"}, nil
		},
	}

	Update(c)
	assert.EqualValues(t, http.StatusOK, response.Code)

	var result faqs.Faq
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result.Id)
	assert.EqualValues(t, "x", result.Question)
	assert.EqualValues(t, "y", result.Answer)
}

func TestDeleteInvalidId(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "x")

	Delete(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "id must be a number", apiErr.Message())
}

func TestDeleteOperationError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		ErrCallback: func() api_errors.ApiError {
			return api_errors.NewInternalServerError(
				"error when tying to delete faq",
				errors.New("database error"))
		},
	}

	Delete(c)
	assert.EqualValues(t, http.StatusInternalServerError, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusInternalServerError, apiErr.Status())
	assert.EqualValues(t, "error when tying to delete faq", apiErr.Message())
	assert.Contains(t, apiErr.Causes(), "database error")
}

func TestDeleteNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		ErrCallback: func() api_errors.ApiError {
			return nil
		},
	}

	Delete(c)
	assert.EqualValues(t, http.StatusOK, response.Code)

	var result map[string]string
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, "deleted", result["status"])
}

func TestGetInvalidId(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "x")

	Get(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "id must be a number", apiErr.Message())
}

func TestGetOperationError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		FaqErrCallback: func() (*faqs.Faq, api_errors.ApiError) {
			return nil, api_errors.NewNotFoundError("faq with given id does not exists")
		},
	}

	Get(c)
	assert.EqualValues(t, http.StatusNotFound, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status())
	assert.EqualValues(t, "faq with given id does not exists", apiErr.Message())
}

func TestGetNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faq/1", nil)
	c, _ := test_utils.GetContextMock(request, response)
	c.AddParam("faq_id", "1")

	services.FaqService = &mocks.FaqServiceMock{
		FaqErrCallback: func() (*faqs.Faq, api_errors.ApiError) {
			return &faqs.Faq{Id: 1, Question: "x", Answer: "y"}, nil
		},
	}

	Get(c)
	assert.EqualValues(t, http.StatusOK, response.Code)

	var result faqs.Faq
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result.Id)
	assert.EqualValues(t, "x", result.Question)
	assert.EqualValues(t, "y", result.Answer)
}

func TestCreateManyInvalidJsonData(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faqs", strings.NewReader(`{"}`))
	c, _ := test_utils.GetContextMock(request, response)

	CreateMany(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := api_errors.NewErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json data", apiErr.Message())
}

func TestCreateManyNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/faqs", strings.NewReader(`[]`))
	c, _ := test_utils.GetContextMock(request, response)

	services.FaqService = &mocks.FaqServiceMock{
		FaqsCallback: func() *faqs.CreateFaqsResponse {
			dummyFaq := &faqs.Faq{Id: 1, UniqHash: "xxx", Question: "x", Answer: "y"}
			return &faqs.CreateFaqsResponse{
				Results:    []faqs.CreateFaqsResult{{Response: dummyFaq}},
				StatusCode: http.StatusCreated,
			}
		},
	}

	CreateMany(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result faqs.CreateFaqsResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 1, result.Results[0].Response.Id)
	assert.EqualValues(t, "xxx", result.Results[0].Response.UniqHash)
	assert.EqualValues(t, "x", result.Results[0].Response.Question)
	assert.EqualValues(t, "y", result.Results[0].Response.Answer)
}
