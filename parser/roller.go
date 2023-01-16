package parser

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/justinian/dice"
	parser "github.com/petertrr/dice-calc-bot/parser/generated"
	impl "github.com/petertrr/dice-calc-bot/parser/impl"
)

type Antrl4BasedRoller struct {
	rand func(int) int
}

func NewAntrl4BasedRoller(rand func(int) int) Antrl4BasedRoller {
	return Antrl4BasedRoller{
		rand: rand,
	}
}

func (roller *Antrl4BasedRoller) Roll(desc string) (dice.RollResult, string, error) {
	log.Println("DEBUG: Rolling [", desc, "]")

	parser.DiceNotationParserInit()

	is := antlr.NewInputStream(desc)
	lexer := parser.NewDiceNotationLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	_parser := parser.NewDiceNotationParser(stream)
	listener := impl.NewDiceNotationListenerImpl(roller.rand)

	antlr.ParseTreeWalkerDefault.Walk(&listener, _parser.Notation())
	return listener.GetResult(), "", nil
}
