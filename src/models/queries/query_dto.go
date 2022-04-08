package queries

type EsQuery struct {
	Equals   []FieldValue `json:"equals"`
	FullText []FieldValue `json:"full_text"`
	Delete   []FieldValue `json:"delete"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
