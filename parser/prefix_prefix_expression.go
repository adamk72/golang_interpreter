package parser

import "monkey/ast"


func (p *Parser) parsePrefixExpression() ast.Expression {
	// defer untrace(trace("parsePrefixExpression"))
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken() // advance so we can parse the expression again.
	expression.Right = p.parseExpression(PREFIX)

	return expression
}