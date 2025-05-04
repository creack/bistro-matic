package lexer

/*
#define	OP_OPEN_PARENT_IDX	0
#define	OP_CLOSE_PARENT_IDX	1
#define	OP_PLUS_IDX		2
#define	OP_SUB_IDX		3
#define	OP_NEG_IDX		3
#define OP_MULT_IDX		4
#define OP_DIV_IDX		5
#define OP_MOD_IDX		6
#define	SYNTAXE_ERROR_MSG	"syntax error\n"
*/

const (
	OpOpenParentIdx int = iota
	OpCloseParentIdx
	OpPlusIdx
	OpSubIdx
	OpMultIdx
	OpDivIdx
	OpModIdx

	OpNegId = OpSubIdx
)

const syntaxErrorMsg = "syntax error\n"

const (
	DefaultBase      = "0123456789"
	DefaultOperators = "()+-*/%"
)

var OpTable = map[int]TokenType{
	OpOpenParentIdx:  TokParenLeft,
	OpCloseParentIdx: TokParenRight,
	OpPlusIdx:        TokPlus,
	OpSubIdx:         TokMinus,
	OpMultIdx:        TokMultiply,
	OpDivIdx:         TokDivide,
	OpModIdx:         TokModulo,
}
