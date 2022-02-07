package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"과제": 과제,
	"공지": 공지,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	IDENT = "IDENT"

	과제 = "과제" //변수 바인딩 명령문 let
	공지 = "공지"
)
