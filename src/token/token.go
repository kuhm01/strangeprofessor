package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"교수님": 교수님,
	"주세요": 공갈,
	"줘":   공갈,
	"내놔":  공갈,

	"A": CONSTINT,
	"B": CONSTINT,
	"C": CONSTINT,
	"D": CONSTINT,
	"F": CONSTINT,

	"여석신청": 여석신청,

	"전공": FLAG,

	"교양":   IF,
	"교양필수": IFFLAG,
	"교양선택": IFFLAG,

	"졸업": EXITPOINT,

	"재학생": DYNAMICVAR,
	"전입":  STACKOPER,
	"전출":  STACKOPER,
	"전과생": DYNAMICVAR,

	"공지":   PRINTER,
	"성적발표": PRINTER,
	"점수발표": PRINTER,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	IDENT = "IDENT"

	교수님 = "교수님"
	공갈  = "공갈"

	STATICVAR  = "STATICVAR"
	DYNAMICVAR = "DYNAMICVAR"

	CONSTINT = "CONSTINT"
	INT      = "INT"
	STRING   = "STRING"

	여석신청 = "여석신청"

	SEMICOLON = "SEMICOLON"

	IF     = "IF"
	IFFLAG = "IFFLAG"

	FLAG = "FLAG"

	STACKOPER = "STACKOPER"
	PRINTER   = "PRINTER"

	DOTHAT = "DOTHAT"

	EXITPOINT = "EXITPOINT"
)
