package faqs

import (
	"encoding/json"
	"errors"
	"github.com/ssoql/faq-chat-bot/src/clients/elasticsearch"
	"github.com/ssoql/faq-chat-bot/src/models/queries"
	"github.com/ssoql/faq-chat-bot/src/utils/api_errors"
)

const indexFaqsName = "faqs"

func (f *FaqDocument) Save() api_errors.ApiError {
	_, err := elasticsearch.EsClient.IndexData(indexFaqsName, f)
	if err != nil {
		return api_errors.NewInternalServerError(
			"error when trying to save faq in ES",
			errors.New("database error"))

	}
	return nil
}

func (f *FaqDocument) Search(query *queries.EsQuery) ([]FaqDocument, api_errors.ApiError) {
	result, err := elasticsearch.EsClient.Search(indexFaqsName, query.BuildFullTextQuery(), 1)
	if err != nil {
		return nil, api_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	documents := make([]FaqDocument, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var document FaqDocument
		if err := json.Unmarshal(bytes, &document); err != nil {
			return nil, api_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		documents[index] = document
	}

	if len(documents) == 0 {
		return nil, api_errors.NewNotFoundError("no documents found matching given criteria")
	}
	return documents, nil
}

func (f *FaqDocument) Delete(query *queries.EsQuery) api_errors.ApiError {
	err := elasticsearch.EsClient.DeleteByQuery(indexFaqsName, query.BuildDeleteQuery())
	if err != nil {
		return api_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}
	return nil
}

func (f *FaqDocument) GetStringId() string {
	return f.UniqHash
}
