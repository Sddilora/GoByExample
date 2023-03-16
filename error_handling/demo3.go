package error_handling

import (
	"errors"
	"fmt"
)

func GuessWhat(guess int) (string, error) {

	theNumberInMyMınd := 50

	if guess < 1 || guess > 100 {
		return "", errors.New("make your guess between 1 and 100 :)")
	}

	if guess > theNumberInMyMınd {
		return "Guess a smaller number", nil
	} else if guess < theNumberInMyMınd {
		return "Guess a bigger number", nil
	} else if guess == theNumberInMyMınd {
		return "Congratulations you guessed my number correctly", nil
	}
	return "", nil
}

func Demo3() {
	fmt.Println(GuessWhat(50))
	fmt.Println(GuessWhat(101))
	fmt.Println(GuessWhat(-12))
	fmt.Println(GuessWhat(56))
}

func Demo3_2() {
	message, err := GuessWhat(50) //If I would only want to print message
	fmt.Println(message)
	if err != nil {
		fmt.Println(err)
	}
}
