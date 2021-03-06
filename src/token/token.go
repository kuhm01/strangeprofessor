package token

import "unicode/utf8"

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
	"교양은":  BRANCHFLAG,
	"교양필수": IFFLAG,
	"교양선택": ELSEFLAG,
	"교양끝":  IFFLAGEXIT,

	"졸업": EXITPOINT,

	"재학생": DYNAMICVAR,
	"전과생": QUEUE,

	"전입":     STACKOPER,
	"전출":     STACKOPER,
	"입학했습니다": IPHOK,
	"수료했습니다": SURYO,

	"공지":   PRINTER,
	"성적발표": P2PRINTER,
	"점수발표": P2PRINTER,
	"평점제출": SUBMITSCORE,

	"평점합": MBUFFEROPER,
	"평점곱": MBUFFEROPER,

	"보강": PLUSCLASS,
	"휴강": MINUSCLASS,

	"신규임용": NEWEMPLOY,
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
			if string(tt) != "!" && string(tt) != "?" && string(tt) != "." {
				flag = false
			}
		}
	}

	if utf8.RuneCountInString(ident) < 4 {
		flag = false
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

	PROFESSOR   = "교수님"
	GIVEME      = "공갈"
	SUBMITSCORE = "SUBMITSCORE"

	STATICVAR  = "STATICVAR"
	DYNAMICVAR = "DYNAMICVAR"
	QUEUE      = "QUEUE"

	CONSTINT = "CONSTINT"
	INT      = "INT"
	STRING   = "STRING"

	YEOSEOK = "여석신청"
	JARI    = "자리"

	SEMICOLON = ";"
	COLON     = ","

	IF         = "IF"
	IFFLAG     = "IFFLAG"
	ELSEFLAG   = "ELSEFLAG"
	BRANCHFLAG = "BRANCHFLAG"

	FLAGIN  = "["
	FLAGOUT = "]"

	IFFLAGEXIT = "IFFLAGEXIT"

	FLAG      = "FLAG"
	JUMPPOINT = "JUMPPOINT"

	STACKOPER = "STACKOPER"
	IPHOK     = "IPHOK"
	SURYO     = "SURYO"

	MBUFFEROPER = "MBUFFEROPER"

	PRINTER   = "PRINTER"
	P2PRINTER = "P2PRINTER"

	MINUSCLASS = "MINUSCLASS"
	PLUSCLASS  = "PLUSCLASS"

	NEWEMPLOY = "NEWEMPLOY"

	BANG     = "!"
	QUESTION = "?"

	EXITPOINT = "EXITPOINT"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)
