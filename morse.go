package main

import (
	"fmt"
)

var morseCodeMap = map[rune]string{
	'A': ".-", 'a': ".-", 'B': "-...", 'b': "-...", 'C': "-.-.", 'c': "-.-.",
	'D': "-..", 'd': "-..", 'E': ".", 'e': ".", 'F': "..-.", 'f': "..-.",
	'G': "--.", 'g': "--.", 'H': "....", 'h': "....", 'I': "..", 'i': "..",
	'J': ".---", 'j': ".---", 'K': "-.-", 'k': "-.-", 'L': ".-..", 'l': ".-..",
	'M': "--", 'm': "--", 'N': "-.", 'n': "-.", 'O': "---", 'o': "---",
	'P': ".--.", 'p': ".--.", 'Q': "--.-", 'q': "--.-", 'R': ".-.", 'r': ".-.",
	'S': "...", 's': "...", 'T': "-", 't': "-", 'U': "..-", 'u': "..-",
	'V': "...-", 'v': "...-", 'W': ".--", 'w': ".--", 'X': "-..-", 'x': "-..-",
	'Y': "-.--", 'y': "-.--", 'Z': "--..", 'z': "--..",
}

func Morse() {
	text := "Introduction to GO"
	morseCode := encodeToMorse(text)
	fmt.Println("Text:", text)
	fmt.Println("Morse Code: \n", morseCode)
}

func encodeToMorse(text string) string {
	var morse string
	for _, char := range text {
		if char == ' ' {
			morse += " || \n"
		} else {
			morse += morseCodeMap[char] + " | "
		}
	}
	return morse
}
