package query

type QueryOperator string

const (
	QueryOperatorEqual           QueryOperator = "eq"
	QueryOperatorNotEqual        QueryOperator = "neq"
	QueryOperatorMoreThan        QueryOperator = "mr"
	QueryOperatorEqualOrMoreThan QueryOperator = "mroeq"
	QueryOperatorLessThan        QueryOperator = "ls"
	QueryOperatorEqualOrLessThan QueryOperator = "lsoeq"
	QueryOperatorContain         QueryOperator = "cn"
	QueryOperatorNotContain      QueryOperator = "ncn"
	QueryOperatorIn              QueryOperator = "in"
	QueryOperatorNotIn           QueryOperator = "nin"
	QueryOperatorEmpty           QueryOperator = "empt"
	QueryOperatorNotEmpty        QueryOperator = "nempt"
)

func GetAllStringQueryOperator() []QueryOperator {

	return []QueryOperator{
		QueryOperatorEqual,
		QueryOperatorNotEqual,
		QueryOperatorContain,
		QueryOperatorNotContain,
		QueryOperatorEmpty,
		QueryOperatorNotEmpty,
		QueryOperatorIn,
		QueryOperatorNotIn,
	}
}

func GetAllQueryOperator() []QueryOperator {

	return []QueryOperator{
		QueryOperatorEqual,
		QueryOperatorNotEqual,
		QueryOperatorContain,
		QueryOperatorNotContain,
		QueryOperatorMoreThan,
		QueryOperatorEqualOrMoreThan,
		QueryOperatorLessThan,
		QueryOperatorEqualOrLessThan,
		QueryOperatorEmpty,
		QueryOperatorNotEmpty,
		QueryOperatorIn,
		QueryOperatorNotIn,
	}
}

type Operand struct {
	Value any
}

func NewOperand(value any) *Operand { return &Operand{Value: value} }

type Query interface {
	GetName() string
	SetName(string)
	GetOperator() QueryOperator
	GetOperand() *Operand
}

type QueryInfo interface {
	GetQuery() []Query
	IsAnd() bool
}

type QueryDefinition interface {
	GetName() string
	GetSupportedOperators() []QueryOperator
	GetType() DataType
}
