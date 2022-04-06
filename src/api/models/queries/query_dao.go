package queries

import "github.com/olivere/elastic"

func (q EsQuery) BuildFullTextQuery() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.FullText {
		equalsQueries = append(equalsQueries, elastic.NewMatchPhraseQuery(eq.Field, eq.Value))
	}
	query.Must(equalsQueries...)
	return query
}
