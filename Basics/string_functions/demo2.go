package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {

	name := "Dilara"

	fmt.Println(s.HasPrefix(name, "Di")) //-> true (bool)
	fmt.Println(s.HasSuffix(name, "hu")) //-> false (bool)

	fmt.Println(s.Index(name, "la")) //-> 2

	letters := []string{"s", "d", "d", "i", "l", "o", "r", "a"}
	joinedLetters := s.Join(letters, "-") //seperator string seperete the elements.
	fmt.Println(joinedLetters)

	fmt.Println(s.Replace(joinedLetters, "-", "*", -1)) //->s*d*d*i*l*o*r*a //negative value means that there is no limit on the number of replacements.

	//121212121,454545454,78787878,36363636,99999999
	fmt.Println(s.Split(joinedLetters, "-"))      //-> [s d d i l o r a]
	fmt.Println(s.SplitAfter(joinedLetters, "-")) //->[s- d- d- i- l- o- r- a]
	fmt.Println(s.Repeat(joinedLetters, 5))       //-> s-d-d-i-l-o-r-as-d-d-i-l-o-r-as-d-d-i-l-o-r-as-d-d-i-l-o-r-as-d-d-i-l-o-r-a
	fmt.Println(s.Repeat(joinedLetters+" ", 5))   //->s-d-d-i-l-o-r-a s-d-d-i-l-o-r-a s-d-d-i-l-o-r-a s-d-d-i-l-o-r-a s-d-d-i-l-o-r-a
}
