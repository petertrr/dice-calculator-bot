package impl

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/justinian/dice"
	parser "github.com/petertrr/dice-calc-bot/parser/generated"
)

type DiceNotationListenerImpl struct {
	*parser.BaseDiceNotationListener

	result dice.StdResult
}

func NewDiceNotationListenerImpl() DiceNotationListenerImpl {
	return DiceNotationListenerImpl{
		result: dice.StdResult{
			Total:   0,
			Rolls:   []int{},
			Dropped: nil,
		},
	}
}

func (l *DiceNotationListenerImpl) GetResult() dice.RollResult {
	return l.result
}

func (l *DiceNotationListenerImpl) ExitDice(ctx *parser.DiceContext) {
	ctx.ADDOPERATOR()
	digits := ctx.AllDIGIT()
	var mult, sides int
	if len(digits) == 1 {
		mult = 1
		sides, _ = strconv.Atoi(digits[0].GetText())
	} else {
		mult, _ = strconv.Atoi(digits[0].GetText())
		sides, _ = strconv.Atoi(digits[1].GetText())
	}
	roll := rand.Intn(int(sides)) + 1
	newRoll := mult * roll
	l.result = dice.StdResult{
		Total: l.result.Total + newRoll,
		Rolls: append(l.result.Rolls, newRoll),
	}
	fmt.Println("DEBUG: Evaluating ", ctx.GetText(), ", got: ", newRoll, ", current result: ", l.result)
	// fmt.Println("DEBUG: ", ctx.ToStringTree())
}
