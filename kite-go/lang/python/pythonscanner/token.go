package pythonscanner

import "strconv"

// Token represents a lexical symbol in the Python programming language.
type Token int

// The list of tokens.
const (
	// Special tokens

	// Illegal represents a lexer error
	Illegal Token = iota
	// BadToken represents a parser error
	BadToken
	// Cursor represents the cursor position
	Cursor
	EOF
	Comment
	Magic // %magic lines in IPython

	whitespacebegin
	LineContinuation
	NewLine
	Indent
	Dedent
	whitespaceend

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)

	literalbegin
	Ident  // main
	Int    // 123
	Long   // 123L
	Float  // 1.23
	Imag   // 1.23i
	String // "abc"
	literalend

	// Operators and delimiters

	operatorbegin
	Add     // +
	Sub     // -
	Mul     // *
	Pow     // **
	Div     // /
	Truediv // //
	Pct     // %

	BitAnd    // &
	BitOr     // |
	BitXor    // ^
	BitNot    // ~
	BitLshift // <<
	BitRshift // >>

	AddAssign     // +=
	SubAssign     // -=
	MulAssign     // *=
	PowAssign     // **
	DivAssign     // /=
	TruedivAssign // //=
	PctAssign     // %=

	BitAndAssign    // &=
	BitOrAssign     // |=
	BitXorAssign    // ^=
	BitLshiftAssign // <<=
	BitRshiftAssign // >>=

	Eq // ==
	Lt // <
	Gt // >
	Le // <=
	Ge // >=
	Lg // <>
	Ne // !=

	Assign // =

	Lparen // (
	Lbrack // [
	Lbrace // {
	Comma  // ,
	Period // .

	Rparen    // )
	Rbrack    // ]
	Rbrace    // }
	Semicolon // ;
	Colon     // :
	At        // @
	Backtick  // `
	Arrow     // ->
	operatorend

	// Keywords; Any keywords added here also has to be added to the keywordMapping map in mappings.go

	keywordbegin
	And
	In
	Is
	Not
	Or

	As
	Assert
	Break
	Class
	Continue
	Def
	Del
	Elif
	Else
	Except
	Finally
	For
	From
	Global
	If
	Import
	Lambda
	Pass
	Raise
	Return
	Try
	While
	With
	Yield

	NonLocal
	Async
	Await
	keywordend

	// These are never generated by the scanner but are created by the parser as operators

	NotIn
	IsNot

	tokencount
)

var tokenMap = [...]string{
	Illegal:  "Illegal",
	BadToken: "BadToken",
	Cursor:   "Cursor",

	EOF:     "EOF",
	Comment: "Comment",
	Magic:   "Magic",

	LineContinuation: "LineContinuation",
	NewLine:          "NewLine",
	Indent:           "Indent",
	Dedent:           "Dedent",

	Ident:  "Ident",
	Int:    "Int",
	Long:   "Long",
	Float:  "Float",
	Imag:   "Imag",
	String: "String",

	Add:     "+",
	Sub:     "-",
	Mul:     "*",
	Pow:     "**",
	Div:     "/",
	Truediv: "//",
	Pct:     "%",

	BitAnd:    "&",
	BitOr:     "|",
	BitXor:    "^",
	BitNot:    "~",
	BitLshift: "<<",
	BitRshift: ">>",

	AddAssign:     "+=",
	SubAssign:     "-=",
	MulAssign:     "*=",
	PowAssign:     "**=",
	DivAssign:     "/=",
	TruedivAssign: "//=",
	PctAssign:     "%=",

	BitAndAssign:    "&=",
	BitOrAssign:     "|=",
	BitXorAssign:    "^=",
	BitLshiftAssign: "<<=",
	BitRshiftAssign: ">>=",

	Eq: "==",
	Lt: "<",
	Gt: ">",
	Le: "<=",
	Ge: ">=",
	Lg: "<>",
	Ne: "!=",

	Assign: "=",

	Lparen: "(",
	Lbrack: "[",
	Lbrace: "{",
	Comma:  ",",
	Period: ".",

	Rparen:    ")",
	Rbrack:    "]",
	Rbrace:    "}",
	Semicolon: ";",
	Colon:     ":",
	At:        "@",
	Backtick:  "`",
	Arrow:     "->",

	And: "and",
	Or:  "or",
	Not: "not",
	In:  "in",
	Is:  "is",

	As:       "as",
	Assert:   "assert",
	Break:    "break",
	Class:    "class",
	Continue: "continue",
	Def:      "def",
	Del:      "del",
	Elif:     "elif",
	Else:     "else",
	Except:   "except",
	Finally:  "finally",
	For:      "for",
	From:     "from",
	Global:   "global",
	If:       "if",
	Import:   "import",
	Lambda:   "lambda",
	Pass:     "pass",
	Raise:    "raise",
	Return:   "return",
	Try:      "try",
	While:    "while",
	With:     "with",
	Yield:    "yield",
	NonLocal: "nonlocal",
	Async:    "async",
	Await:    "await",

	// These are never generated by the scanner but are created by the parser as operators
	NotIn: "not in",
	IsNot: "is not",
}

// Keywords maps from the string literal for a python keyword to the Token for that keyword.
var Keywords map[string]Token
var inverseTokens map[string]Token

var (
	// LiteralTokens contains all literal tokens
	LiteralTokens []Token
	// OperatorTokens contains all operator tokens
	OperatorTokens []Token
	// KeywordTokens is a list of all keyword tokens
	KeywordTokens []Token
	// WhitespaceTokens is a list of all whitespace tokens
	WhitespaceTokens []Token
	// Tokens contains all tokens
	Tokens []Token
)

func init() {
	Keywords = make(map[string]Token)
	for i := keywordbegin + 1; i < keywordend; i++ {
		Keywords[tokenMap[i]] = i
	}

	inverseTokens = make(map[string]Token)
	for i, str := range tokenMap {
		if str != "" {
			inverseTokens[str] = Token(i)
		}
	}

	// create list of all tokens
	for i := Illegal; i < tokencount; i++ {
		Tokens = append(Tokens, i)
	}
	for i := literalbegin + 1; i < literalend; i++ {
		LiteralTokens = append(LiteralTokens, i)
	}
	for i := operatorbegin + 1; i < operatorend; i++ {
		OperatorTokens = append(OperatorTokens, i)
	}
	for i := keywordbegin + 1; i < keywordend; i++ {
		KeywordTokens = append(KeywordTokens, i)
	}
	for i := whitespacebegin + 1; i < whitespaceend; i++ {
		WhitespaceTokens = append(WhitespaceTokens, i)
	}
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token Add, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token Ident, the string is "Ident").
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokenMap)) {
		s = tokenMap[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator tok. If tok is not a binary operator, the result
// is LowestPrecedence.
func (tok Token) Precedence() int {
	switch tok {
	case Lambda:
		return 1
	case Or:
		return 2
	case And:
		return 3
	case Not:
		return 4
	case In, Is, Eq, Ne, Lt, Le, Gt, Ge, Lg, NotIn, IsNot:
		return 5
	case BitOr:
		return 6
	case BitAnd:
		return 7
	case BitLshift:
		return 8
	case BitRshift:
		return 9
	case Add, Sub:
		return 10
	case Mul, Div, Truediv, Pct:
		return 11
	case BitNot:
		return 12
	case Pow:
		return 13
	}
	return LowestPrec
}

// Lookup maps an identifier to its keyword token or Ident (if not a keyword).
func Lookup(ident string) Token {
	if tok, iskeyword := Keywords[ident]; iskeyword {
		return tok
	}
	return Ident
}

// ByName maps a token name to its integer representation
func ByName(name string) Token {
	if tok, ok := inverseTokens[name]; ok {
		return tok
	}
	return Illegal
}

// IsWhitespace returns true for tokens corresponding to whitespace
func (tok Token) IsWhitespace() bool { return whitespacebegin < tok && tok < whitespaceend }

// IsLiteral returns true for tokens corresponding to identifiers and primitives
func (tok Token) IsLiteral() bool { return literalbegin < tok && tok < literalend }

// IsOperator returns true for tokens corresponding to operators and delimiters
func (tok Token) IsOperator() bool { return operatorbegin < tok && tok < operatorend }

// IsKeyword returns true for tokens corresponding to keywords
func (tok Token) IsKeyword() bool { return keywordbegin < tok && tok < keywordend }