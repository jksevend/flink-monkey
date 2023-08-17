package lexer

import "flink_monkey/token"

// Lexer takes in source code as input and produce tokens that represents it
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// NewLexer creates a new instance of the Lexer struct and consumes
// the first characters from the input
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.advance()
	return lexer
}

// advance consumes the next byte of the input
func (lexer *Lexer) advance() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

// peek ahead by one character in the input
func (lexer *Lexer) peek() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

// eatWhitespace skips any whitespace characters and advances the Lexer
func (lexer *Lexer) eatWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.advance()
	}
}

// readIdentifier advances the source input until no letter is present
// anymore and returns the parsed characters
func (lexer *Lexer) readIdentifier() string {
	// 'Cache' the current lexer position
	position := lexer.position

	// Advance for each letter in the input
	for isLetter(lexer.char) {
		lexer.advance()
	}

	// Return a substring of the input
	return lexer.input[position:lexer.position]
}

// readNumber advances the source input checking for number until no number
// is present anymore and returns the parsed characters
func (lexer *Lexer) readNumber() string {
	// 'Cache' the current lexer position
	position := lexer.position

	// Advance for each number in the input
	for isNumber(lexer.char) {
		lexer.advance()
	}

	// Return a substring of the input
	return lexer.input[position:lexer.position]
}

// isLetter checks if input matches [A-Z][a-z]_
func isLetter(input byte) bool {
	return 'a' <= input && input <= 'z' || 'A' <= input && input <= 'Z' || input == '_'
}

// isNumber checks if input matches [0-9]
func isNumber(input byte) bool {
	return '0' <= input && input <= '9'
}

// NextToken returns a Token representation of the Lexers current character
func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token

	// Whitespace at beginning
	lexer.eatWhitespace()

	switch lexer.char {

	case ';':
		nextToken = token.NewToken(token.SEMICOLON, lexer.char)
	case '(':
		nextToken = token.NewToken(token.LPAREN, lexer.char)
	case ')':
		nextToken = token.NewToken(token.RPAREN, lexer.char)
	case ',':
		nextToken = token.NewToken(token.COMMA, lexer.char)
	case '{':
		nextToken = token.NewToken(token.LBRACE, lexer.char)
	case '}':
		nextToken = token.NewToken(token.RBRACE, lexer.char)
	case '=':
		// Check for '=='
		if lexer.peek() == '=' {
			char := lexer.char
			lexer.advance()
			nextToken = token.NewMultiCharToken(token.EQ, string(char)+string(lexer.char))
		} else {
			nextToken = token.NewToken(token.ASSIGN, lexer.char)
		}

	case '+':
		nextToken = token.NewToken(token.PLUS, lexer.char)
	case '-':
		nextToken = token.NewToken(token.MINUS, lexer.char)
	case '!':
		// Check for '!='
		if lexer.peek() == '=' {
			char := lexer.char
			lexer.advance()
			nextToken = token.NewMultiCharToken(token.NOT_EQ, string(char)+string(lexer.char))
		} else {
			nextToken = token.NewToken(token.BANG, lexer.char)
		}
	case '/':
		nextToken = token.NewToken(token.SLASH, lexer.char)
	case '*':
		nextToken = token.NewToken(token.ASTERISK, lexer.char)
	case '<':
		nextToken = token.NewToken(token.LT, lexer.char)
	case '>':
		nextToken = token.NewToken(token.GT, lexer.char)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			// Read the whole identifier and look up the type
			nextToken.Literal = lexer.readIdentifier()
			nextToken.Type = token.LookUpIdentifier(nextToken.Literal)
			return nextToken
		} else if isNumber(lexer.char) {
			// Read the number
			nextToken.Literal = lexer.readNumber()
			nextToken.Type = token.INT
			return nextToken
		} else {
			nextToken = token.NewToken(token.ILLEGAL, lexer.char)
		}
	}

	// Advance to the next character
	lexer.advance()
	return nextToken
}
