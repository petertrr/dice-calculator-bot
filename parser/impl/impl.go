package impl

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/justinian/dice"
	parser "github.com/petertrr/dice-calc-bot/parser/generated"
)

type DiceNotationListenerImpl struct {
	*parser.BaseDiceNotationListener

	result dice.StdResult

	stack []dice.StdResult
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

func (l *DiceNotationListenerImpl) pop() dice.StdResult {
	log.Println("DEBUG: pop, size=", len(l.stack), "stack=", l.stack)
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	log.Println("DEBUG: popped, size=", len(l.stack))
	return result
}

func (l *DiceNotationListenerImpl) push(r dice.StdResult) {
	log.Println("DEBUG: push, size=", len(l.stack))
	l.stack = append(l.stack, r)
	log.Println("DEBUG: pushed, size=", len(l.stack), "stack=", l.stack)
}

func (l *DiceNotationListenerImpl) ExitNotation(ctx *parser.NotationContext) {
	if len(l.stack) != 1 {
		log.Panicln("Stack contains multiple results still: ", l.stack)
	}
	l.result = l.pop()
}

func (l *DiceNotationListenerImpl) ExitAddOp(ctx *parser.AddOpContext) {
	if len(ctx.AllADDOPERATOR()) == 0 {
		return
	} else {
		// fixme: may be more than 2
		sign := getSign(interface{}(ctx))
		result2 := l.pop()
		result1 := l.pop()
		l.push(dice.StdResult{
			Total: result1.Total + sign*result2.Total,
			Rolls: append(result1.Rolls, result2.Rolls...),
		})
	}
}

func (l *DiceNotationListenerImpl) ExitMultOp(ctx *parser.MultOpContext) {
	if len(ctx.AllMULTOPERATOR()) == 0 {
		return
	} else {
		// fixme: may be more than 2
		result2 := l.pop()
		result1 := l.pop()
		l.push(dice.StdResult{
			Total: result1.Total * result2.Total,
			Rolls: append(result1.Rolls, result2.Rolls...),
		})
	}
}

func (l *DiceNotationListenerImpl) ExitDice(ctx *parser.DiceContext) {
	sign := getSign(interface{}(ctx))
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
	l.push(dice.StdResult{
		Total: sign * newRoll,
		Rolls: []int{newRoll},
	})
	log.Println("DEBUG: Evaluating ", ctx.GetText(), ", got: ", newRoll, ", current result: ", l.stack[len(l.stack)-1])
}

/**
 * @param ctx either a [parser.DiceContext] or [parser.NumberContext].
 * Other types that have [ADDOPERATOR()] method can be added manually.
 */
func getSign(ctx interface{}) int {
	var sgn int
	var addOperator antlr.TerminalNode
	dc, ok := ctx.(*parser.DiceContext)
	if ok {
		addOperator = dc.ADDOPERATOR()
	} else {
		nc, ok := ctx.(*parser.NumberContext)
		if ok {
			addOperator = nc.ADDOPERATOR()
		} else {
			ac, ok := ctx.(*parser.AddOpContext)
			if ok && len(ac.AllADDOPERATOR()) > 0 {
				addOperator = ac.ADDOPERATOR(0)
			}
		}
	}
	if addOperator == nil {
		sgn = 1
	} else {
		isPlus := addOperator.GetText() == "+" // FixMe: get constant value from generated ANTLR code
		if isPlus {
			sgn = 1
		} else {
			sgn = -1
		}
	}
	return sgn
}
