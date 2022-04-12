package faq

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/faq-chat-bot/src/models/faqs"
	"github.com/ssoql/faq-chat-bot/src/services"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var faq faqs.FaqCreateInput

	if err := c.ShouldBindJSON(&faq); err != nil {
		apiErr := api_errors.NewBadRequestError("invalid json data")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, opErr := services.FaqService.CreateFaq(&faq)
	if opErr != nil {
		c.JSON(opErr.Status(), opErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateMany(c *gin.Context) {
	var input []faqs.FaqCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		apiErr := api_errors.NewBadRequestError("invalid json data")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result := services.FaqService.CreateFaqs(input)
	c.JSON(result.StatusCode, result)
}

func Update(c *gin.Context) {
	faqId, idErr := getId(c.Param("faq_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	var faq faqs.FaqUpdateInput
	if err := c.ShouldBindJSON(&faq); err != nil {
		apiErr := api_errors.NewBadRequestError("invalid json data")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, opErr := services.FaqService.UpdateFaq(faqId, &faq)
	if opErr != nil {
		c.JSON(opErr.Status(), opErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	faqId, idErr := getId(c.Param("faq_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.FaqService.DeleteFaq(faqId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Get(c *gin.Context) {
	faqId, idErr := getId(c.Param("faq_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	faq, opErr := services.FaqService.GetFaq(faqId)
	if opErr != nil {
		c.JSON(opErr.Status(), opErr)
		return
	}
	c.JSON(http.StatusOK, faq)
}

func getId(idParam string) (int64, api_errors.ApiError) {
	id, userErr := strconv.ParseInt(idParam, 10, 64)
	if userErr != nil {
		return 0, api_errors.NewBadRequestError("id must be a number")
	}
	return id, nil
}
