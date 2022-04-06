package faqs

import (
	"time"
)

type FaqDocument struct {
	Id        int64     `json:"id"`
	UniqHash  string    `json:"uniq_hash"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FaqDocuments []FaqDocument

func NewDocFromFaq(faq *Faq) *FaqDocument {
	return &FaqDocument{
		Id:        faq.Id,
		UniqHash:  faq.UniqHash,
		Question:  faq.Question,
		Answer:    faq.Answer,
		CreatedAt: faq.CreatedAt,
		UpdatedAt: faq.UpdatedAt,
	}
}
