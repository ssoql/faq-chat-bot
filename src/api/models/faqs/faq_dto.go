package faqs

import (
	"github.com/ssoql/faq-chat-bot/src/api/utils/api_errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Faq struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	UniqHash  string         `gorm:"unique" json:"uniq_hash"`
	Question  string         `gorm:"not null" json:"question"`
	Answer    string         `gorm:"not null" json:"answer"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Faqs []Faq

type FaqCreateInput struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type FaqUpdateInput FaqCreateInput

func (faq *Faq) Validate() api_errors.ApiError {
	faq.Question = strings.TrimSpace(faq.Question)
	if faq.Question == "" {
		return api_errors.NewBadRequestError("invalid question")
	}

	faq.Answer = strings.TrimSpace(faq.Answer)
	if faq.Answer == "" {
		return api_errors.NewBadRequestError("invalid answer")
	}
	return nil
}
