package diceprinter

import (
	"fmt"

	dice "github.com/chelios/appliedgo-dice"
)

func PrintRoll(sides int, comment string) {
	fmt.Printf("%s: %d\n", comment, dice.Roll(sides))
}
