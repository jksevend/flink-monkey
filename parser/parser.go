package parser

import (
	"flink_monkey/ast"
	"flink_monkey/lexer"
	"flink_monkey/token"
	"fmt"
)

// Parser takes in tokens from lexer.Lexer and constructs an AST
type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	Errors       []string
}

// NewParser creates a new Parser instance and advances the parser
// two times, so currentToken and peekToken are filled
func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, Errors: []string{}}
	parser.advance()
	parser.advance()
	return parser
}

// ParseProgram creates a root AST node and attempts to parse statements
// until an EOF token is reached
func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for !parser.currentTokenIs(token.EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.advance()
	}
	return program
}

// parseStatement checks the currentToken and attempts to parse a
// corresponding generic statement
func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.VAR:
		return parser.parseVarStatement()
	default:
		return nil
	}
}

// advance advances the parsers currentToken and peekToken
func (parser *Parser) advance() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

// parseVarStatement attempts to parse a variable statement
func (parser *Parser) parseVarStatement() *ast.VarStatement {
	// Construct the node
	statement := &ast.VarStatement{Token: parser.currentToken}

	// Identifiers are expected next
	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}
	ident := ast.NewIdentifier(parser.currentToken, parser.currentToken.Literal)
	statement.Name = &ident

	// Assign is expected
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// Then a semicolon
	//TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.advance()
	}
	return statement
}

// currentTokenIs checks whether the token-type of currentToken and
// another token matches
func (parser *Parser) currentTokenIs(token token.TokenType) bool {
	return parser.currentToken.Type == token
}

// currentTokenIs checks whether the token-type of peekToken and
// another token matches
func (parser *Parser) peekTokenIs(token token.TokenType) bool {
	return parser.peekToken.Type == token
}

// expectPeek checks peekTokenIs and advances the parser
func (parser *Parser) expectPeek(token token.TokenType) bool {
	if parser.peekTokenIs(token) {
		parser.advance()
		return true
	} else {
		parser.peekError(token)
		return false
	}
}

func (parser *Parser) peekError(token token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		token, parser.peekToken.Type)
	parser.Errors = append(parser.Errors, msg)
}
