package home

import (
	"github.com/ssoql/faq-chat-bot/src/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowHomePage(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	c, router := test_utils.GetContextMock(request, response)
	// template mock
	templ := template.Must(template.New("index.tmpl").Parse(`{{.title}}`))
	router.SetHTMLTemplate(templ)

	ShowHomePage(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t,
		true,
		strings.Contains(response.Body.String(), "Simple Chat Bot"))
}
