package helpers

// Definisikan tipe baru string
type RelOp string

// Konstanta string enum
const (
	EQ    RelOp = "EQUAL"
	NOTEQ RelOp = "NOT EQUAL"
	IN    RelOp = "IN"
	NOTIN RelOp = "NOT IN"
)
