package impl

import (
	"errors"
	"strconv"
)

type KeepMode int

const (
	Keep KeepMode = iota
	Drop
)

type SortMode int

const (
	Lowest SortMode = iota
	Highest
)

type KeepModifier struct {
	Mode     KeepMode
	SortMode SortMode
	Number   int
}

func ParseModifier(desc string) *KeepModifier {
	var KeepMode KeepMode
	keepModeRune := rune(desc[0])
	if keepModeRune == rune('k') {
		KeepMode = Keep
	} else if keepModeRune == rune('i') {
		KeepMode = Drop
	} else {
		panic(errors.New("Unknown character for keep mode <" + string(keepModeRune) + "> in <" + desc + ">"))
	}

	var SortMode SortMode
	sortModeRune := rune(desc[1])
	if sortModeRune == 'h' {
		SortMode = Highest
	} else if sortModeRune == 'l' {
		SortMode = Lowest
	} else {
		panic(errors.New("Unknown character for sort mode <" + string(sortModeRune) + "> in <" + desc + ">"))
	}

	number, err := strconv.Atoi(desc[2:])
	if err != nil {
		panic(err)
	}

	return &KeepModifier{
		Mode:     KeepMode,
		SortMode: SortMode,
		Number:   number,
	}
}
