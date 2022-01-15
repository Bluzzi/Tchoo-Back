package promises

type DatabaseEntry struct {
	PetNonce int64 `json:"pet_nonce" bson:"pet_nonce"`
	ExecutionTimestamp int64 `json:"execution_timestamp" bson:"execution_timestamp"`
	Type string `json:"type" bson:"type"`
	Identifier string `json:"identifier" bson:"identifier"`
	UniqueIdentifier string `json:"unique_identifier" bson:"unique_identifier"`
	Value float64 `json:"value" bson:"value"`
	Field string `json:"field" bson:"field"`
}

var (
	FieldPetNonce = "pet_nonce"
	FieldExecutionTimestamp = "execution_timestamp"
	FieldIdentifier = "identifier"
	FieldUniqueIdentifier = "unique_identifier"
	FieldType = "type"
	FieldValue = "value"
	FieldField = "field"

	TypeIncrement = "increment"
	TypeDecrement = "decrement"
)