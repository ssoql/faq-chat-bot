package mocks

import (
	"github.com/olivere/elastic"
	"github.com/ssoql/faq-chat-bot/src/clients/elasticsearch"
)

type EsClientMock struct {
}

func (c *EsClientMock) SetClient(client *elastic.Client) {}
func (c *EsClientMock) IndexData(name string, doc elasticsearch.EsDocumentInterface) (*elastic.IndexResponse, error) {
	return nil, nil
}
func (c *EsClientMock) GetById(string, string) (*elastic.GetResult, error) {
	return nil, nil
}
func (c *EsClientMock) Search(string, elastic.Query, int) (*elastic.SearchResult, error) {
	return nil, nil
}
func (c *EsClientMock) DeleteByQuery(string, elastic.Query) error {
	return nil
}
