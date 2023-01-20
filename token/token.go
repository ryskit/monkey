package token

const (
	ILLEGAL = "ILLEGAL" // トークンや文字が未知であることを表す
	EOF     = "EOF"     // ファイル終端(end of file)を表し、構文解析器にここで停止してよいと伝える

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 123456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPRREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "fn"
	LET      = "let"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
