package parser_test

import (
	"testing"

	"github.com/justinian/dice"
	"github.com/petertrr/dice-calc-bot/parser"
)

func TestBasicRolls(t *testing.T) {
	roller := parser.NewAntrl4BasedRoller(
		func(x int) int { return x/2 + 1 },
	)
	var result dice.RollResult

	result, _, _ = roller.Roll("d4")
	shouldBeTotal(result, 3, t)
	shouldBeRolls(result.(dice.StdResult), []int{3}, t)

	result, _, _ = roller.Roll("d4*d6+d8")
	shouldBeTotal(result, 17, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 4, 5}, t)

	result, _, _ = roller.Roll("-d2")
	shouldBeTotal(result, -2, t)
	shouldBeRolls(result.(dice.StdResult), []int{2}, t)

	result, _, _ = roller.Roll("d4-d2")
	shouldBeTotal(result, 1, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 2}, t)

	result, _, _ = roller.Roll("d4+d2-d6")
	shouldBeTotal(result, 1, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 2, 4}, t)
}

func TestRollsWithMultipliers(t *testing.T) {
	roller := parser.NewAntrl4BasedRoller(
		func(x int) int { return x/2 + 1 },
	)
	var result dice.RollResult

	result, _, _ = roller.Roll("2d4")
	shouldBeTotal(result, 6, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 3}, t)

	result, _, _ = roller.Roll("2*d4")
	shouldBeTotal(result, 6, t)
	shouldBeRolls(result.(dice.StdResult), []int{3}, t)

	result, _, _ = roller.Roll("1d6*2d4")
	shouldBeTotal(result, 24, t)
	shouldBeRolls(result.(dice.StdResult), []int{4, 3, 3}, t)
}

func TestRollsWithNumbers(t *testing.T) {
	roller := parser.NewAntrl4BasedRoller(
		func(x int) int { return x/2 + 1 },
	)
	var result dice.RollResult

	result, _, _ = roller.Roll("2d4+2")
	shouldBeTotal(result, 8, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 3}, t)

	result, _, _ = roller.Roll("2d4-2")
	shouldBeTotal(result, 4, t)
	shouldBeRolls(result.(dice.StdResult), []int{3, 3}, t)

	result, _, _ = roller.Roll("2*d4+2")
	shouldBeTotal(result, 8, t)
	shouldBeRolls(result.(dice.StdResult), []int{3}, t)

	result, _, _ = roller.Roll("2*d4-2")
	shouldBeTotal(result, 4, t)
	shouldBeRolls(result.(dice.StdResult), []int{3}, t)
}

func TestKeepModifier(t *testing.T) {
	numInvocations := 0
	roller := parser.NewAntrl4BasedRoller(
		func(x int) int {
			numInvocations++
			if numInvocations%2 == 0 {
				return x / 4
			} else {
				return 3 * x / 4
			}
		},
	)
	var result dice.RollResult

	result, _, _ = roller.Roll("2d20kh1")
	shouldBeTotal(result, 15, t)
	shouldBeRolls(result.(dice.StdResult), []int{15, 5}, t)
	shouldBeDroppedRolls(result.(dice.StdResult), []int{5}, t)

	result, _, _ = roller.Roll("2d20kl1")
	shouldBeTotal(result, 5, t)
	shouldBeRolls(result.(dice.StdResult), []int{15, 5}, t)
	shouldBeDroppedRolls(result.(dice.StdResult), []int{15}, t)

	result, _, _ = roller.Roll("2d20ih1")
	shouldBeTotal(result, 5, t)
	shouldBeRolls(result.(dice.StdResult), []int{15, 5}, t)
	shouldBeDroppedRolls(result.(dice.StdResult), []int{15}, t)

	result, _, _ = roller.Roll("2d20il1")
	shouldBeTotal(result, 15, t)
	shouldBeRolls(result.(dice.StdResult), []int{15, 5}, t)
	shouldBeDroppedRolls(result.(dice.StdResult), []int{5}, t)

	result, _, _ = roller.Roll("2*2d20kh1")
	shouldBeTotal(result, 30, t)
	shouldBeRolls(result.(dice.StdResult), []int{15, 5}, t)
	shouldBeDroppedRolls(result.(dice.StdResult), []int{5}, t)
}

func shouldBeTotal(d dice.RollResult, total int, t *testing.T) {
	if d.(dice.StdResult).Total != total {
		t.Helper()
		t.Error("Expected total to be ", total, ", but found ", d.(dice.StdResult).Total)
	}
}

func shouldBeRolls(d dice.StdResult, rolls []int, t *testing.T) {
	if !arrayEquals(d.Rolls, rolls) {
		t.Helper()
		t.Error("Expected rolls to be ", rolls, ", but found ", d.Rolls)
	}
}

func shouldBeDroppedRolls(d dice.StdResult, droppedRolls []int, t *testing.T) {
	if !arrayEquals(d.Dropped, droppedRolls) {
		t.Helper()
		t.Error("Expected droppedRolls to be ", droppedRolls, ", but found ", d.Dropped)
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
