package parser_test

import (
	"math/rand"
	"testing"

	"github.com/justinian/dice"
	"github.com/petertrr/dice-calc-bot/parser"
)

func Test(t *testing.T) {
	roller := parser.Antrl4BasedRoller{}
	rand.Seed(0)

	var result dice.RollResult

	result, _, _ = roller.Roll("d4")
	shouldBeTotal(result, 3, t)
	shouldBeRolls(result.(dice.StdResult), []int{3}, t)

	result, _, _ = roller.Roll("d4*d6+d8")
	shouldBeTotal(result, 9, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 2, 3}, t)

	result, _, _ = roller.Roll("-d2")
	if result.Int() > 0 {
		t.Error("Should be negative, but got ", result)
	}
	shouldBeRolls(result.(dice.StdResult), []int{2}, t)

	result, _, _ = roller.Roll("d4-d2")
	shouldBeTotal(result, -1, t)
	shouldBeRolls(result.(dice.StdResult), []int{1, 2}, t)

	result, _, _ = roller.Roll("d4+d2-d6")
	shouldBeTotal(result, 2, t)
	shouldBeRolls(result.(dice.StdResult), []int{2, 1, 1}, t)

	roller.Roll("2d4")
	roller.Roll("2*d4")
	roller.Roll("1d6*2d4")
	roller.Roll("2d4+2")
	roller.Roll("2d4-2")
	roller.Roll("2*d4+2")
	roller.Roll("2*d4-2")
}

func shouldBeTotal(d dice.RollResult, total int, t *testing.T) {
	if d.Int() != total {
		t.Error("Expected total to be ", total, ", but found ", d.Int())
	}
}

func shouldBeRolls(d dice.StdResult, rolls []int, t *testing.T) {
	if !arrayEquals(d.Rolls, rolls) {
		t.Error("Expected rolls to be ", rolls, ", but found ", d.Rolls)
	}
}

func arrayEquals(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i1, e1 := range a1 {
		if e1 != a2[i1] {
			return false
		}
	}
	return true
}
