// Code generated from /home/ryanj/sim/SimParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SimParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 128,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2,
	3, 2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 3, 3, 3, 7, 3, 23, 10,
	3, 12, 3, 14, 3, 26, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14, 3, 56, 11,
	3, 5, 3, 58, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 67,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 83, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 96, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 116, 10, 4, 12, 4, 14, 4, 119, 11, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 5, 6, 126, 10, 6, 3, 6, 2, 3, 6, 7, 2, 4, 6, 8, 10, 2, 8,
	4, 2, 10, 11, 39, 39, 4, 2, 16, 17, 20, 20, 3, 2, 18, 19, 3, 2, 29, 32,
	3, 2, 27, 28, 3, 2, 21, 26, 2, 151, 2, 17, 3, 2, 2, 2, 4, 82, 3, 2, 2,
	2, 6, 95, 3, 2, 2, 2, 8, 120, 3, 2, 2, 2, 10, 125, 3, 2, 2, 2, 12, 13,
	5, 4, 3, 2, 13, 14, 5, 10, 6, 2, 14, 16, 3, 2, 2, 2, 15, 12, 3, 2, 2, 2,
	16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 3, 3, 2,
	2, 2, 19, 17, 3, 2, 2, 2, 20, 24, 7, 35, 2, 2, 21, 23, 5, 4, 3, 2, 22,
	21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2,
	2, 25, 27, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 83, 7, 36, 2, 2, 28, 29,
	7, 4, 2, 2, 29, 30, 5, 6, 4, 2, 30, 31, 5, 4, 3, 2, 31, 83, 3, 2, 2, 2,
	32, 33, 7, 5, 2, 2, 33, 83, 5, 4, 3, 2, 34, 35, 7, 5, 2, 2, 35, 36, 5,
	6, 4, 2, 36, 37, 5, 4, 3, 2, 37, 83, 3, 2, 2, 2, 38, 39, 7, 5, 2, 2, 39,
	40, 7, 40, 2, 2, 40, 41, 7, 21, 2, 2, 41, 42, 5, 6, 4, 2, 42, 43, 7, 6,
	2, 2, 43, 44, 5, 6, 4, 2, 44, 45, 5, 4, 3, 2, 45, 83, 3, 2, 2, 2, 46, 47,
	7, 3, 2, 2, 47, 48, 7, 40, 2, 2, 48, 57, 7, 33, 2, 2, 49, 54, 5, 6, 4,
	2, 50, 51, 7, 38, 2, 2, 51, 53, 5, 6, 4, 2, 52, 50, 3, 2, 2, 2, 53, 56,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2,
	56, 54, 3, 2, 2, 2, 57, 49, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 60, 7, 34, 2, 2, 60, 61, 7, 37, 2, 2, 61, 83, 7, 40, 2, 2,
	62, 63, 7, 40, 2, 2, 63, 66, 7, 40, 2, 2, 64, 65, 7, 21, 2, 2, 65, 67,
	5, 6, 4, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 83, 3, 2, 2, 2,
	68, 69, 7, 40, 2, 2, 69, 70, 5, 8, 5, 2, 70, 71, 5, 6, 4, 2, 71, 83, 3,
	2, 2, 2, 72, 73, 7, 7, 2, 2, 73, 83, 5, 6, 4, 2, 74, 75, 7, 15, 2, 2, 75,
	76, 7, 33, 2, 2, 76, 77, 5, 6, 4, 2, 77, 78, 7, 34, 2, 2, 78, 83, 3, 2,
	2, 2, 79, 83, 7, 7, 2, 2, 80, 83, 7, 8, 2, 2, 81, 83, 7, 9, 2, 2, 82, 20,
	3, 2, 2, 2, 82, 28, 3, 2, 2, 2, 82, 32, 3, 2, 2, 2, 82, 34, 3, 2, 2, 2,
	82, 38, 3, 2, 2, 2, 82, 46, 3, 2, 2, 2, 82, 62, 3, 2, 2, 2, 82, 68, 3,
	2, 2, 2, 82, 72, 3, 2, 2, 2, 82, 74, 3, 2, 2, 2, 82, 79, 3, 2, 2, 2, 82,
	80, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 85, 8, 4, 1,
	2, 85, 86, 7, 33, 2, 2, 86, 87, 5, 6, 4, 2, 87, 88, 7, 34, 2, 2, 88, 96,
	3, 2, 2, 2, 89, 90, 7, 19, 2, 2, 90, 96, 5, 6, 4, 12, 91, 92, 7, 14, 2,
	2, 92, 96, 5, 6, 4, 11, 93, 96, 7, 40, 2, 2, 94, 96, 9, 2, 2, 2, 95, 84,
	3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2,
	95, 94, 3, 2, 2, 2, 96, 117, 3, 2, 2, 2, 97, 98, 12, 10, 2, 2, 98, 99,
	9, 3, 2, 2, 99, 116, 5, 6, 4, 11, 100, 101, 12, 9, 2, 2, 101, 102, 9, 4,
	2, 2, 102, 116, 5, 6, 4, 10, 103, 104, 12, 8, 2, 2, 104, 105, 9, 5, 2,
	2, 105, 116, 5, 6, 4, 9, 106, 107, 12, 7, 2, 2, 107, 108, 9, 6, 2, 2, 108,
	116, 5, 6, 4, 8, 109, 110, 12, 6, 2, 2, 110, 111, 7, 12, 2, 2, 111, 116,
	5, 6, 4, 7, 112, 113, 12, 5, 2, 2, 113, 114, 7, 13, 2, 2, 114, 116, 5,
	6, 4, 6, 115, 97, 3, 2, 2, 2, 115, 100, 3, 2, 2, 2, 115, 103, 3, 2, 2,
	2, 115, 106, 3, 2, 2, 2, 115, 109, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 116,
	119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 7, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 121, 9, 7, 2, 2, 121, 9, 3, 2, 2, 2,
	122, 126, 7, 2, 2, 3, 123, 126, 6, 6, 8, 2, 124, 126, 6, 6, 9, 2, 125,
	122, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 11, 3,
	2, 2, 2, 12, 17, 24, 54, 57, 66, 82, 95, 115, 117, 125,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'function'", "'if'", "'loop'", "'to'", "'return'", "'break'", "'continue'",
	"'true'", "'false'", "'and'", "'or'", "'not'", "'print'", "'*'", "'/'",
	"'+'", "'-'", "'%'", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'('", "')'", "'{'", "'}'", "':'",
	"','",
}
var symbolicNames = []string{
	"", "FUNCTION", "IF", "LOOP", "TO", "RETURN", "BREAK", "CONTINUE", "TRUE",
	"FALSE", "AND", "OR", "NOT", "PRINT", "MULTIPLY", "DIVIDE", "ADD", "SUBTRACT",
	"MODULO", "ASSIGNMENT", "ADD_ASSIGNMENT", "SUB_ASSIGNMENT", "MUL_ASSIGNMENT",
	"DIV_ASSIGNMENT", "MOD_ASSIGNMENT", "EQUALS", "NOT_EQUALS", "GREATER",
	"LESSER", "GREATER_OR_EQUAL", "LESSER_OR_EQUAL", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "COLON", "COMMA", "NUMBER", "IDENTIFIER", "NEWLINE", "WHITESPACE",
	"LINE_COMMENT", "BLOCK_COMMENT",
}

var ruleNames = []string{
	"start", "statement", "expression", "assignment_op", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SimParser struct {
	*antlr.BaseParser
}

func NewSimParser(input antlr.TokenStream) *SimParser {
	this := new(SimParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SimParser.g4"

	return this
}

// SimParser tokens.
const (
	SimParserEOF              = antlr.TokenEOF
	SimParserFUNCTION         = 1
	SimParserIF               = 2
	SimParserLOOP             = 3
	SimParserTO               = 4
	SimParserRETURN           = 5
	SimParserBREAK            = 6
	SimParserCONTINUE         = 7
	SimParserTRUE             = 8
	SimParserFALSE            = 9
	SimParserAND              = 10
	SimParserOR               = 11
	SimParserNOT              = 12
	SimParserPRINT            = 13
	SimParserMULTIPLY         = 14
	SimParserDIVIDE           = 15
	SimParserADD              = 16
	SimParserSUBTRACT         = 17
	SimParserMODULO           = 18
	SimParserASSIGNMENT       = 19
	SimParserADD_ASSIGNMENT   = 20
	SimParserSUB_ASSIGNMENT   = 21
	SimParserMUL_ASSIGNMENT   = 22
	SimParserDIV_ASSIGNMENT   = 23
	SimParserMOD_ASSIGNMENT   = 24
	SimParserEQUALS           = 25
	SimParserNOT_EQUALS       = 26
	SimParserGREATER          = 27
	SimParserLESSER           = 28
	SimParserGREATER_OR_EQUAL = 29
	SimParserLESSER_OR_EQUAL  = 30
	SimParserLPAREN           = 31
	SimParserRPAREN           = 32
	SimParserLBRACE           = 33
	SimParserRBRACE           = 34
	SimParserCOLON            = 35
	SimParserCOMMA            = 36
	SimParserNUMBER           = 37
	SimParserIDENTIFIER       = 38
	SimParserNEWLINE          = 39
	SimParserWHITESPACE       = 40
	SimParserLINE_COMMENT     = 41
	SimParserBLOCK_COMMENT    = 42
)

// SimParser rules.
const (
	SimParserRULE_start         = 0
	SimParserRULE_statement     = 1
	SimParserRULE_expression    = 2
	SimParserRULE_assignment_op = 3
	SimParserRULE_eos           = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StartContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StartContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimParserFUNCTION)|(1<<SimParserIF)|(1<<SimParserLOOP)|(1<<SimParserRETURN)|(1<<SimParserBREAK)|(1<<SimParserCONTINUE)|(1<<SimParserPRINT))) != 0) || _la == SimParserLBRACE || _la == SimParserIDENTIFIER {
		{
			p.SetState(10)
			p.Statement()
		}
		{
			p.SetState(11)
			p.Eos()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	*StatementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SimParserBREAK, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementContext struct {
	*StatementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(SimParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalLoopStatementContext struct {
	*StatementContext
}

func NewConditionalLoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalLoopStatementContext {
	var p = new(ConditionalLoopStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ConditionalLoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalLoopStatementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(SimParserLOOP, 0)
}

func (s *ConditionalLoopStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalLoopStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ConditionalLoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterConditionalLoopStatement(s)
	}
}

func (s *ConditionalLoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitConditionalLoopStatement(s)
	}
}

func (s *ConditionalLoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitConditionalLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionStatementContext struct {
	*StatementContext
}

func NewFunctionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionStatementContext {
	var p = new(FunctionStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FunctionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStatementContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(SimParserFUNCTION, 0)
}

func (s *FunctionStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SimParserIDENTIFIER)
}

func (s *FunctionStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SimParserIDENTIFIER, i)
}

func (s *FunctionStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserLPAREN, 0)
}

func (s *FunctionStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserRPAREN, 0)
}

func (s *FunctionStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(SimParserCOLON, 0)
}

func (s *FunctionStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimParserCOMMA)
}

func (s *FunctionStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimParserCOMMA, i)
}

func (s *FunctionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitFunctionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementContext struct {
	*StatementContext
	varName antlr.Token
}

func NewAssignmentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignmentStatementContext) GetVarName() antlr.Token { return s.varName }

func (s *AssignmentStatementContext) SetVarName(v antlr.Token) { s.varName = v }

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) Assignment_op() IAssignment_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_opContext)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimParserIDENTIFIER, 0)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type InfiniteLoopStatementContext struct {
	*StatementContext
}

func NewInfiniteLoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InfiniteLoopStatementContext {
	var p = new(InfiniteLoopStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *InfiniteLoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfiniteLoopStatementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(SimParserLOOP, 0)
}

func (s *InfiniteLoopStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *InfiniteLoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterInfiniteLoopStatement(s)
	}
}

func (s *InfiniteLoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitInfiniteLoopStatement(s)
	}
}

func (s *InfiniteLoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitInfiniteLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	*StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SimParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStatementContext struct {
	*StatementContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SimParserPRINT, 0)
}

func (s *PrintStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserLPAREN, 0)
}

func (s *PrintStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserRPAREN, 0)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStatementContext struct {
	*StatementContext
}

func NewBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimParserLBRACE, 0)
}

func (s *BlockStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimParserRBRACE, 0)
}

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopStatementContext struct {
	*StatementContext
	min IExpressionContext
	max IExpressionContext
}

func NewLoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopStatementContext {
	var p = new(LoopStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *LoopStatementContext) GetMin() IExpressionContext { return s.min }

func (s *LoopStatementContext) GetMax() IExpressionContext { return s.max }

func (s *LoopStatementContext) SetMin(v IExpressionContext) { s.min = v }

func (s *LoopStatementContext) SetMax(v IExpressionContext) { s.max = v }

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(SimParserLOOP, 0)
}

func (s *LoopStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimParserIDENTIFIER, 0)
}

func (s *LoopStatementContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserASSIGNMENT, 0)
}

func (s *LoopStatementContext) TO() antlr.TerminalNode {
	return s.GetToken(SimParserTO, 0)
}

func (s *LoopStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LoopStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LoopStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	*StatementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SimParserCONTINUE, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationStatementContext struct {
	*StatementContext
	type_   antlr.Token
	varName antlr.Token
}

func NewDeclarationStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationStatementContext {
	var p = new(DeclarationStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *DeclarationStatementContext) GetType_() antlr.Token { return s.type_ }

func (s *DeclarationStatementContext) GetVarName() antlr.Token { return s.varName }

func (s *DeclarationStatementContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *DeclarationStatementContext) SetVarName(v antlr.Token) { s.varName = v }

func (s *DeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SimParserIDENTIFIER)
}

func (s *DeclarationStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SimParserIDENTIFIER, i)
}

func (s *DeclarationStatementContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserASSIGNMENT, 0)
}

func (s *DeclarationStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterDeclarationStatement(s)
	}
}

func (s *DeclarationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitDeclarationStatement(s)
	}
}

func (s *DeclarationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitDeclarationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBlockStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(SimParserLBRACE)
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimParserFUNCTION)|(1<<SimParserIF)|(1<<SimParserLOOP)|(1<<SimParserRETURN)|(1<<SimParserBREAK)|(1<<SimParserCONTINUE)|(1<<SimParserPRINT))) != 0) || _la == SimParserLBRACE || _la == SimParserIDENTIFIER {
			{
				p.SetState(19)
				p.Statement()
			}

			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(25)
			p.Match(SimParserRBRACE)
		}

	case 2:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(SimParserIF)
		}
		{
			p.SetState(27)
			p.expression(0)
		}
		{
			p.SetState(28)
			p.Statement()
		}

	case 3:
		localctx = NewInfiniteLoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Match(SimParserLOOP)
		}
		{
			p.SetState(31)
			p.Statement()
		}

	case 4:
		localctx = NewConditionalLoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(32)
			p.Match(SimParserLOOP)
		}
		{
			p.SetState(33)
			p.expression(0)
		}
		{
			p.SetState(34)
			p.Statement()
		}

	case 5:
		localctx = NewLoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(36)
			p.Match(SimParserLOOP)
		}
		{
			p.SetState(37)
			p.Match(SimParserIDENTIFIER)
		}
		{
			p.SetState(38)
			p.Match(SimParserASSIGNMENT)
		}
		{
			p.SetState(39)

			var _x = p.expression(0)

			localctx.(*LoopStatementContext).min = _x
		}
		{
			p.SetState(40)
			p.Match(SimParserTO)
		}
		{
			p.SetState(41)

			var _x = p.expression(0)

			localctx.(*LoopStatementContext).max = _x
		}
		{
			p.SetState(42)
			p.Statement()
		}

	case 6:
		localctx = NewFunctionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.Match(SimParserFUNCTION)
		}
		{
			p.SetState(45)
			p.Match(SimParserIDENTIFIER)
		}
		{
			p.SetState(46)
			p.Match(SimParserLPAREN)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(SimParserTRUE-8))|(1<<(SimParserFALSE-8))|(1<<(SimParserNOT-8))|(1<<(SimParserSUBTRACT-8))|(1<<(SimParserLPAREN-8))|(1<<(SimParserNUMBER-8))|(1<<(SimParserIDENTIFIER-8)))) != 0 {
			{
				p.SetState(47)
				p.expression(0)
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SimParserCOMMA {
				{
					p.SetState(48)
					p.Match(SimParserCOMMA)
				}
				{
					p.SetState(49)
					p.expression(0)
				}

				p.SetState(54)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(57)
			p.Match(SimParserRPAREN)
		}
		{
			p.SetState(58)
			p.Match(SimParserCOLON)
		}
		{
			p.SetState(59)
			p.Match(SimParserIDENTIFIER)
		}

	case 7:
		localctx = NewDeclarationStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(60)

			var _m = p.Match(SimParserIDENTIFIER)

			localctx.(*DeclarationStatementContext).type_ = _m
		}
		{
			p.SetState(61)

			var _m = p.Match(SimParserIDENTIFIER)

			localctx.(*DeclarationStatementContext).varName = _m
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(62)
				p.Match(SimParserASSIGNMENT)
			}
			{
				p.SetState(63)
				p.expression(0)
			}

		}

	case 8:
		localctx = NewAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(66)

			var _m = p.Match(SimParserIDENTIFIER)

			localctx.(*AssignmentStatementContext).varName = _m
		}
		{
			p.SetState(67)
			p.Assignment_op()
		}
		{
			p.SetState(68)
			p.expression(0)
		}

	case 9:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(70)
			p.Match(SimParserRETURN)
		}
		{
			p.SetState(71)
			p.expression(0)
		}

	case 10:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(72)
			p.Match(SimParserPRINT)
		}
		{
			p.SetState(73)
			p.Match(SimParserLPAREN)
		}
		{
			p.SetState(74)
			p.expression(0)
		}
		{
			p.SetState(75)
			p.Match(SimParserRPAREN)
		}

	case 11:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(77)
			p.Match(SimParserRETURN)
		}

	case 12:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(78)
			p.Match(SimParserBREAK)
		}

	case 13:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(79)
			p.Match(SimParserCONTINUE)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NegateExpressionContext struct {
	*ExpressionContext
}

func NewNegateExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateExpressionContext {
	var p = new(NegateExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateExpressionContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(SimParserSUBTRACT, 0)
}

func (s *NegateExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterNegateExpression(s)
	}
}

func (s *NegateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitNegateExpression(s)
	}
}

func (s *NegateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitNegateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewAddSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExpressionContext {
	var p = new(AddSubExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *AddSubExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *AddSubExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AddSubExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AddSubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(SimParserADD, 0)
}

func (s *AddSubExpressionContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(SimParserSUBTRACT, 0)
}

func (s *AddSubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterAddSubExpression(s)
	}
}

func (s *AddSubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitAddSubExpression(s)
	}
}

func (s *AddSubExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitAddSubExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InequalityExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewInequalityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InequalityExpressionContext {
	var p = new(InequalityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InequalityExpressionContext) GetOp() antlr.Token { return s.op }

func (s *InequalityExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *InequalityExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *InequalityExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *InequalityExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *InequalityExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *InequalityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InequalityExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InequalityExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InequalityExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SimParserGREATER, 0)
}

func (s *InequalityExpressionContext) LESSER() antlr.TerminalNode {
	return s.GetToken(SimParserLESSER, 0)
}

func (s *InequalityExpressionContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(SimParserGREATER_OR_EQUAL, 0)
}

func (s *InequalityExpressionContext) LESSER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(SimParserLESSER_OR_EQUAL, 0)
}

func (s *InequalityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterInequalityExpression(s)
	}
}

func (s *InequalityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitInequalityExpression(s)
	}
}

func (s *InequalityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitInequalityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *AndExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *AndExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AndExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SimParserAND, 0)
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SimParserNUMBER, 0)
}

func (s *LiteralExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SimParserTRUE, 0)
}

func (s *LiteralExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SimParserFALSE, 0)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableExpressionContext struct {
	*ExpressionContext
}

func NewVariableExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableExpressionContext {
	var p = new(VariableExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VariableExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimParserIDENTIFIER, 0)
}

func (s *VariableExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterVariableExpression(s)
	}
}

func (s *VariableExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitVariableExpression(s)
	}
}

func (s *VariableExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitVariableExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *EqualityExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *EqualityExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *EqualityExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SimParserEQUALS, 0)
}

func (s *EqualityExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(SimParserNOT_EQUALS, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *OrExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *OrExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OrExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SimParserOR, 0)
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensExpressionContext struct {
	*ExpressionContext
}

func NewParensExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExpressionContext {
	var p = new(ParensExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParensExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserLPAREN, 0)
}

func (s *ParensExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimParserRPAREN, 0)
}

func (s *ParensExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterParensExpression(s)
	}
}

func (s *ParensExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitParensExpression(s)
	}
}

func (s *ParensExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitParensExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivModExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewMulDivModExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModExpressionContext {
	var p = new(MulDivModExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivModExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *MulDivModExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *MulDivModExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *MulDivModExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *MulDivModExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivModExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivModExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(SimParserMULTIPLY, 0)
}

func (s *MulDivModExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(SimParserDIVIDE, 0)
}

func (s *MulDivModExpressionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SimParserMODULO, 0)
}

func (s *MulDivModExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterMulDivModExpression(s)
	}
}

func (s *MulDivModExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitMulDivModExpression(s)
	}
}

func (s *MulDivModExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitMulDivModExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SimParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, SimParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimParserLPAREN:
		localctx = NewParensExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(83)
			p.Match(SimParserLPAREN)
		}
		{
			p.SetState(84)
			p.expression(0)
		}
		{
			p.SetState(85)
			p.Match(SimParserRPAREN)
		}

	case SimParserSUBTRACT:
		localctx = NewNegateExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)
			p.Match(SimParserSUBTRACT)
		}
		{
			p.SetState(88)
			p.expression(10)
		}

	case SimParserNOT:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Match(SimParserNOT)
		}
		{
			p.SetState(90)
			p.expression(9)
		}

	case SimParserIDENTIFIER:
		localctx = NewVariableExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(SimParserIDENTIFIER)
		}

	case SimParserTRUE, SimParserFALSE, SimParserNUMBER:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(SimParserTRUE-8))|(1<<(SimParserFALSE-8))|(1<<(SimParserNUMBER-8)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MulDivModExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimParserMULTIPLY)|(1<<SimParserDIVIDE)|(1<<SimParserMODULO))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)

					var _x = p.expression(9)

					localctx.(*MulDivModExpressionContext).right = _x
				}

			case 2:
				localctx = NewAddSubExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AddSubExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimParserADD || _la == SimParserSUBTRACT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)

					var _x = p.expression(8)

					localctx.(*AddSubExpressionContext).right = _x
				}

			case 3:
				localctx = NewInequalityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*InequalityExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*InequalityExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimParserGREATER)|(1<<SimParserLESSER)|(1<<SimParserGREATER_OR_EQUAL)|(1<<SimParserLESSER_OR_EQUAL))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*InequalityExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)

					var _x = p.expression(7)

					localctx.(*InequalityExpressionContext).right = _x
				}

			case 4:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*EqualityExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(105)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimParserEQUALS || _la == SimParserNOT_EQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(106)

					var _x = p.expression(6)

					localctx.(*EqualityExpressionContext).right = _x
				}

			case 5:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AndExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(108)
					p.Match(SimParserAND)
				}
				{
					p.SetState(109)

					var _x = p.expression(5)

					localctx.(*AndExpressionContext).right = _x
				}

			case 6:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OrExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimParserRULE_expression)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(111)
					p.Match(SimParserOR)
				}
				{
					p.SetState(112)

					var _x = p.expression(4)

					localctx.(*OrExpressionContext).right = _x
				}

			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignment_opContext is an interface to support dynamic dispatch.
type IAssignment_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_opContext differentiates from other interfaces.
	IsAssignment_opContext()
}

type Assignment_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_opContext() *Assignment_opContext {
	var p = new(Assignment_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimParserRULE_assignment_op
	return p
}

func (*Assignment_opContext) IsAssignment_opContext() {}

func NewAssignment_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_opContext {
	var p = new(Assignment_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimParserRULE_assignment_op

	return p
}

func (s *Assignment_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_opContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserASSIGNMENT, 0)
}

func (s *Assignment_opContext) ADD_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserADD_ASSIGNMENT, 0)
}

func (s *Assignment_opContext) SUB_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserSUB_ASSIGNMENT, 0)
}

func (s *Assignment_opContext) MUL_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserMUL_ASSIGNMENT, 0)
}

func (s *Assignment_opContext) DIV_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserDIV_ASSIGNMENT, 0)
}

func (s *Assignment_opContext) MOD_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SimParserMOD_ASSIGNMENT, 0)
}

func (s *Assignment_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterAssignment_op(s)
	}
}

func (s *Assignment_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitAssignment_op(s)
	}
}

func (s *Assignment_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitAssignment_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimParser) Assignment_op() (localctx IAssignment_opContext) {
	localctx = NewAssignment_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimParserRULE_assignment_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimParserASSIGNMENT)|(1<<SimParserADD_ASSIGNMENT)|(1<<SimParserSUB_ASSIGNMENT)|(1<<SimParserMUL_ASSIGNMENT)|(1<<SimParserDIV_ASSIGNMENT)|(1<<SimParserMOD_ASSIGNMENT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimParserVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(SimParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(121)

		if !(lineTerminatorAhead(p)) {
			panic(antlr.NewFailedPredicateException(p, "lineTerminatorAhead(p)", ""))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(122)

		if !(checkPreviousTokenText(p, "}")) {
			panic(antlr.NewFailedPredicateException(p, "checkPreviousTokenText(p, \"}\")", ""))
		}

	}

	return localctx
}

func (p *SimParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 4:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return lineTerminatorAhead(p)

	case 7:
		return checkPreviousTokenText(p, "}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
