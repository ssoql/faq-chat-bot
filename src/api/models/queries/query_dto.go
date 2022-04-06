package queries

type EsQuery struct {
	Equals   []FieldValue `json:"equals"`
	FullText []FieldValue `json:"full_text"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
