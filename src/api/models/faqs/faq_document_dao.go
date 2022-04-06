package faqs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ssoql/faq-chat-bot/src/api/clients/elasticsearch"
	"github.com/ssoql/faq-chat-bot/src/api/models/queries"
	"github.com/ssoql/faq-chat-bot/src/api/utils/api_errors"
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

	fmt.Println("###")
	fmt.Printf("%v\n", result.Hits.Hits)
	fmt.Printf("%v\n", result.Hits.TotalHits)
	fmt.Println("###")

	documents := make([]FaqDocument, result.TotalHits())
	for index, hit := range result.Hits.Hits {

		bytes, _ := hit.Source.MarshalJSON()
		fmt.Printf("HITTT: %v\n", string(bytes))
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

func (f *FaqDocument) GetStringId() string {
	return f.UniqHash
}
