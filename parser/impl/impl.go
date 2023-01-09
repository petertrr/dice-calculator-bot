package impl

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/justinian/dice"
	parser "github.com/petertrr/dice-calc-bot/parser/generated"
)

type DiceNotationListenerImpl struct {
	*parser.BaseDiceNotationListener

	result dice.StdResult

	stack *lls.Stack
}

func NewDiceNotationListenerImpl() DiceNotationListenerImpl {
	return DiceNotationListenerImpl{
		result: dice.StdResult{
			Total:   0,
			Rolls:   []int{},
			Dropped: nil,
		},
		stack: lls.New(),
	}
}

func (l *DiceNotationListenerImpl) GetResult() dice.RollResult {
	return l.result
}

func (l *DiceNotationListenerImpl) pop() dice.StdResult {
	log.Println("DEBUG: pop, size=", l.stack.Size(), "stack=", l.stack)
	result, _ := l.stack.Pop()
	return result.(dice.StdResult)
}

func (l *DiceNotationListenerImpl) push(r dice.StdResult) {
	l.stack.Push(interface{}(r))
	log.Println("DEBUG: pushed, size=", l.stack.Size(), "stack=", l.stack)
}

func (l *DiceNotationListenerImpl) ExitNotation(ctx *parser.NotationContext) {
	if l.stack.Size() == 0 && ctx.GetText() == "" {
		log.Println("DEBUG: Attempt to get result when stack is empty, expression is [", ctx.GetText(), "]")
		return
	} else if l.stack.Size() != 1 {
		log.Panicln("Stack contains multiple results still: ", l.stack)
	}
	l.result = l.pop()
}

func (l *DiceNotationListenerImpl) ExitAddOp(ctx *parser.AddOpContext) {
	log.Println("DEBUG: Exiting AddOp node [", ctx.GetText(), "]")
	numAdds := len(ctx.AllADDOPERATOR())
	if numAdds == 0 {
		return
	} else {
		var results []dice.StdResult = []dice.StdResult{}
		for i := 0; i < numAdds+1; i++ {
			results = append(results, l.pop())
		}
		lastIndex := len(results) - 1
		var result dice.StdResult = results[lastIndex]
		for i := 0; i < numAdds; i++ {
			sign := getSign(ctx.ADDOPERATOR(i))
			result2 := results[lastIndex-i-1]
			result = dice.StdResult{
				Total: result.Total + sign*result2.Total,
				Rolls: append(result.Rolls, result2.Rolls...),
			}
		}
		l.push(result)
	}
}

/**
 * FixMe: `N * DX` should be equal to `N` independent rollings instead of multiplication
 */
func (l *DiceNotationListenerImpl) ExitMultOp(ctx *parser.MultOpContext) {
	log.Println("DEBUG: Exiting MultOp node [", ctx.GetText(), "]")
	numMults := len(ctx.AllMULTOPERATOR())
	if numMults == 0 {
		return
	} else {
		for i := numMults - 1; i >= 0; i-- {
			result2 := l.pop()
			result1 := l.pop()
			l.push(dice.StdResult{
				Total: result1.Total * result2.Total,
				Rolls: append(result1.Rolls, result2.Rolls...),
			})
		}
	}
}

func (l *DiceNotationListenerImpl) ExitDice(ctx *parser.DiceContext) {
	sign := getSign(ctx.ADDOPERATOR())
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
	lastResult, _ := l.stack.Peek()
	log.Println("DEBUG: Evaluating ", ctx.GetText(), ", got: ", newRoll, ", current result: ", lastResult)
}

/**
 * @param ctx either a [parser.DiceContext] or [parser.NumberContext].
 * Other types that have [ADDOPERATOR()] method can be added manually.
 */
func getSign(addOperator antlr.TerminalNode) int {
	var sgn int
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
