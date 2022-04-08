package faqs

import (
	"errors"
	"github.com/ssoql/faq-chat-bot/src/datasources/faqs_db"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
	"github.com/ssoql/faq-chat-bot/src/utils/crypto_utils"
	"strings"
)

func (faq *Faq) Get() api_errors.ApiError {
	if err := faqs_db.Client.Where("id = ?", faq.Id).First(faq).Error; err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "record not found") {
			return api_errors.NewNotFoundError("faq with given id does not exists")
		}
		return api_errors.NewInternalServerError("error when tying to fetch faq", errors.New("database error: "+err.Error()))
	}
	return nil
}

func (faq *Faq) Save() api_errors.ApiError {
	faq.UniqHash = crypto_utils.GetMd5(strings.ToLower(faq.Question))
	if err := faqs_db.Client.Create(faq).Error; err != nil {
		//logger.Error("error when trying to prepare save faq statement", err.Error())
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return api_errors.NewBadRequestError("this question already exists")
		}
		return api_errors.NewInternalServerError("error when tying to save faq", errors.New("database error"))
	}
	return nil
}

func (faq *Faq) Update() api_errors.ApiError {

	if err := faqs_db.Client.Updates(&faq).Error; err != nil {
		//logger.Error("error when trying to prepare save faq statement", err.Error())
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return api_errors.NewBadRequestError("this question already exists")
		}
		return api_errors.NewInternalServerError("error when tying to save faq", errors.New("database error"))
	}
	return nil
}

func (faq *Faq) Delete() api_errors.ApiError {
	// perform soft delete
	if err := faqs_db.Client.Delete(&faq).Error; err != nil {
		return api_errors.NewInternalServerError("error when tying to delete faq", errors.New("database error"))
	}
	return nil
}
