package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"교수님": PROFESSOR,
	"주세요": GIVEME,
	"줘":   GIVEME,
	"내놔":  GIVEME,

	"A": CONSTINT,
	"B": CONSTINT,
	"C": CONSTINT,
	"D": CONSTINT,
	"F": CONSTINT,

	"여석신청": YEOSEOK,
	"자리":   JARI,

	"전공":  FLAG,
	"재수강": JUMPPOINT,

	"교양":   IF,
	"교양필수": IFFLAG,
	"교양선택": ELSEFLAG,

	"졸업": EXITPOINT,

	"재학생": DYNAMICVAR,
	"전입":  STACKOPER,
	"전출":  STACKOPER,
	"전과생": DYNAMICVAR,

	"공지":   PRINTER,
	"성적발표": PRINTER,
	"점수발표": PRINTER,
}

func isRequireProfessor(ident string) bool {
	flag := true
	for i, tt := range ident {
		switch i {
		case 0:
			if string(tt) != "교" {
				flag = false
			}
		case 3:
			if string(tt) != "수" {
				flag = false
			}
		case 6:
			if string(tt) != "님" {
				flag = false
			}
		default:
			if string(tt) != "!" && string(tt) != "?" {
				flag = false
			}
		}
	}

	return flag
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else if isRequireProfessor(ident) {
		return STATICVAR
	}

	return ILLEGAL
}

const (
	IDENT = "IDENT"

	PROFESSOR = "교수님"
	GIVEME    = "공갈"

	STATICVAR  = "STATICVAR"
	DYNAMICVAR = "DYNAMICVAR"

	CONSTINT = "CONSTINT"
	INT      = "INT"
	STRING   = "STRING"

	YEOSEOK = "여석신청"
	JARI    = "자리"

	SEMICOLON = ";"
	COLON     = ","

	IF       = "IF"
	IFFLAG   = "IFFLAG"
	ELSEFLAG = "ELSEFLAG"

	FLAG      = "FLAG"
	JUMPPOINT = "JUMPPOINT"

	STACKOPER = "STACKOPER"
	PRINTER   = "PRINTER"

	BANG     = "!"
	QUESTION = "?"

	EXITPOINT = "EXITPOINT"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)
