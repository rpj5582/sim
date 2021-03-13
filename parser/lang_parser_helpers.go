package parser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func lineTerminatorAhead(p *SimParser) bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	if possibleIndexEosToken < 0 {
		return false
	}

	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return false
	}

	if ahead.GetTokenType() == SimParserNEWLINE {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == SimParserWHITESPACE {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		if possibleIndexEosToken < 0 {
			return false
		}

		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == SimParserBLOCK_COMMENT && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) || (_type == SimParserNEWLINE)
}

func checkPreviousTokenText(p *SimParser, text string) bool {
	stream := p.GetTokenStream()
	return stream.LT(1).GetText() == text
}
