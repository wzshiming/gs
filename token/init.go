package token

var (
	LookupKeywork  = newLooker()
	LookupOperator = newLooker()
)

func init() {

	for k, v := range [][]Token{
		{DEFINE, ASSIGN,
			ADD_ASSIGN, SUB_ASSIGN, MUL_ASSIGN, QUO_ASSIGN, POW_ASSIGN, REM_ASSIGN,
			AND_ASSIGN, OR_ASSIGN, XOR_ASSIGN, SHL_ASSIGN, SHR_ASSIGN, AND_NOT_ASSIGN},
		{COMMA},
		{COLON},
		{LOR},
		{LAND},
		{EQL, NEQ, LSS, LEQ, GTR, GEQ},
		{ADD, SUB, OR, XOR},
		{MUL, QUO, REM, SHL, SHR, AND, AND_NOT},
		{POW},
		{PERIOD},
	} {
		for _, v0 := range v {
			v0.setPrecedence(k + 1)
		}
	}

	for i := keyworkBeg + 1; i != keyworkEnd; i++ {
		LookupKeywork.Add([]rune(tokenMap[i]), i)
	}
	for i := operatorBeg + 1; i != operatorEnd; i++ {
		LookupOperator.Add([]rune(tokenMap[i]), i)
	}
}
