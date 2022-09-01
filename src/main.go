package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	text := "ala"
	value := isPalindromo(text)
	value2 := isPalindromoRune(text)

	fmt.Printf("¿%s es palindromo? : %t \n", text, value)
	fmt.Printf("¿%s es palindromo2? : %t \n", text, value2)

	fmt.Println("************* Range continued *******************")

	pow2 := make([]int, 10)

	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func isPalindromo(text string) bool {
	// cuando accedemos al indice de un string (osea apuntamos al caracter), nos arrojara un int que es su codigo acsi, asi que necesitamos convertirlo a string.

	var textReverse string

	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	}

	return false
}

func isPalindromoRune(text string) bool {
	//los slices a comparar seran runes, no strings (runes es el valor en acsi de un character)
	textSlice := []rune(text)
	reverseSlice := make([]rune, len(textSlice))

	for i := len(reverseSlice) - 1; i >= 0; i-- {
		reverseSlice[len(textSlice)-1-i] = (textSlice[i])
	}

	if reflect.DeepEqual(textSlice, reverseSlice) {
		return true
	}

	return false
}
