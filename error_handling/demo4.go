package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

// This is a method for borderException struct
// We can increase the performance by giving borderException with a pointer
// (The pointer is the star at the beginning of the word).
// func (b *borderException) Error() string {
func (b borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func GuessWhat2(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", &borderException{guess, "Out of range"}
	}
	return "You guessed right", nil
}
