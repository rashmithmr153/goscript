package lexer

type TokenType int

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	INT
	FLOAT
	STRING
	ASSIGN
	OPERATOR
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	SEMICOLON
	COMMA
	EOF
)
var keywords = map[string]TokenType{
    "let": KEYWORD,
    "if":  KEYWORD,
	"else":KEYWORD,
    "for": KEYWORD,
	"func":KEYWORD,
	"return": KEYWORD,
}

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Input   string
	CurPos  int
	NextPos int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:   input,
		CurPos:  0,
		NextPos: 0,
	}
}
func (L *Lexer) advance() byte {
	if L.NextPos >= len(L.Input) {
		return 0
	}
	L.CurPos = L.NextPos
	L.NextPos++
	return L.Input[L.CurPos]
}
func isLetter(ch byte) bool {
    return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isDigit(ch byte) bool{
	return ch>='0' && ch <='9'
}

func (L *Lexer) readIdentifier() string {
    start := L.CurPos
    for L.NextPos < len(L.Input) && isLetter(L.Input[L.NextPos]) {
        L.advance()
    }
    return L.Input[start:L.NextPos]
}

func (L *Lexer) readNumber() string{
	start := L.CurPos
    for L.NextPos < len(L.Input) && isDigit(L.Input[L.NextPos]) {
        L.advance()
    }
    return L.Input[start:L.NextPos]
}

func (L *Lexer) NextToken() Token {
	for L.NextPos < len(L.Input) && L.Input[L.NextPos] == ' ' {
    L.advance()
	}
	ch:=L.advance()
	var typ TokenType
	switch ch {
	case '(':
		typ=LPAREN
	case ')':
		typ=RPAREN
	case '{':
		typ=LBRACE
	case '}':
		typ=RBRACE
	case '=':
		if L.Input[L.NextPos]=='='{
			L.advance()
			return Token{Type: OPERATOR ,Value:"=="}
		}
		typ=ASSIGN
	case '>':
    	if L.Input[L.NextPos] == '=' {
    	    L.advance()
    	    return Token{Type: OPERATOR, Value: ">="}
    	}
    	return Token{Type: OPERATOR, Value: ">"}
	case '<':
		if L.Input[L.NextPos] == '=' {
    	    L.advance()
    	    return Token{Type: OPERATOR, Value: "<="}
    	}
    	return Token{Type: OPERATOR, Value: "<"}
	case '+', '-','*','/':
		typ=OPERATOR
	case ';':
		typ=SEMICOLON
	case ',':
    	typ = COMMA
	case 0:
		typ=EOF
	default:
	    if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
        	word := L.readIdentifier()
			if val,exists:=keywords[word];exists{
				return Token{Type: val,Value:word}
			}
        	return Token{Type: IDENTIFIER, Value: word}
    	}
		if ch >='0' && ch<='9'{
			num:=L.readNumber()
			return Token{Type: INT, Value: num}
		}
	}
	return Token{
		Type: typ,
		Value: string(ch),
	}
}