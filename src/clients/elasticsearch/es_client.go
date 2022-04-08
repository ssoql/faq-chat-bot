package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/ssoql/faq-chat-bot/src/config"
	"time"
)

const envElasticHosts = "ES_HOSTS"

var EsClient esClientInterface = &esClient{}

type esClient struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(*elastic.Client)
	IndexData(string, esDocumentInterface) (*elastic.IndexResponse, error)
	GetById(string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query, int) (*elastic.SearchResult, error)
	DeleteByQuery(string, elastic.Query) error
}

type esDocumentInterface interface {
	GetStringId() string
}

func init() {
	client, err := elastic.NewClient(
		elastic.SetURL(config.GetEsHosts()),
		elastic.SetHealthcheckInterval(1*time.Second),
		//elastic.SetErrorLog(log),
		//elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	EsClient.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) IndexData(indexName string, document esDocumentInterface) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Id(document.GetStringId()).
		Index(indexName).
		BodyJson(document).
		Do(ctx)

	if err != nil {
		//logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) GetById(indexName string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(indexName).
		Id(id).
		Do(ctx)
	if err != nil {
		//logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(indexName string, query elastic.Query, maxResults int) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(indexName).
		Query(query).
		RestTotalHitsAsInt(true).
		Size(maxResults).
		Do(ctx)
	if err != nil {
		//logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) DeleteByQuery(indexName string, query elastic.Query) error {
	ctx := context.Background()
	_, err := c.client.DeleteByQuery(indexName).
		Query(query).
		Do(ctx)
	if err != nil {
		//logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return err
	}
	return nil
}
