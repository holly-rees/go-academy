// Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
// 7 and 11 are to be called NATURAL
// 2 is called SNAKE-EYES-CRAPS
// 3 and 12 is called LOSS-CRAPS
// Any other combination is called NEUTRAL.

package assignment7

import (
	"fmt"
	"math/rand/v2"
)

func Assignment7() {
	dice := &RealDice{}
	for i := 0; i < 50; i++ {
		RollDice(dice)
	}
}

func RollDice(dice Dice) {

	firstDiceResult, secondDiceResult := dice.Roll()
	result := firstDiceResult + secondDiceResult

	resultString := fmt.Sprintf("Rolled a %v and a %v: ", firstDiceResult, secondDiceResult)

	switch result {
	case 7, 11:
		resultString += "NATURAL!"
	case 2:
		resultString += "SNAKE-EYES-CRAPS!"
	case 3, 12:
		resultString += "LOSS-CRAPS!"
	default:
		resultString += "NEUTRAL!"
	}
	fmt.Println(resultString)

}

type Dice interface {
	Roll() (int, int)
}

type RealDice struct {
}

func (d RealDice) Roll() (int, int) {
	return rand.IntN(6) + 1, rand.IntN(6) + 1
}
