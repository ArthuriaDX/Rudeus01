// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

// _alphabet is a mapping of Alpha numeric characters to Morse code
var _alphabet map[string]string

var _reverseAlphabet map[string]string

func morseInit() {
	if _alphabet == nil {
		init_alphabet()
	}
	if _reverseAlphabet == nil {
		init_reverse()
	}
}
func init_alphabet() {
	_alphabet = map[string]string{
		"A":  ".-",
		"B":  "-...",
		"C":  "-.-.",
		"D":  "-..",
		"E":  ".",
		"F":  "..-.",
		"G":  "--.",
		"H":  "....",
		"I":  "..",
		"J":  ".---",
		"K":  "-.-",
		"L":  ".-..",
		"M":  "--",
		"N":  "-.",
		"O":  "---",
		"P":  ".--.",
		"Q":  "--.-",
		"R":  ".-.",
		"S":  "...",
		"T":  "-",
		"U":  "..-",
		"V":  "...-",
		"W":  ".--",
		"X":  "-..-",
		"Y":  "-.--",
		"Z":  "--..",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		"0":  "-----",
		" ":  "/",
		".":  ".-.-.-",  // period
		":":  "---...",  // colon
		",":  "--..--",  // comma
		";":  "-.-.-",   // semicolon
		"?":  "..--..",  // question
		"=":  "-...-",   // equals
		"'":  ".----.",  // apostrophe
		"/":  "-..-.",   // slash
		"!":  "-.-.--",  // exclamation
		"-":  "-....-",  // dash
		"_":  "..--.-",  // underline
		"\"": ".-..-.",  // quotation marks
		"(":  "-.--.",   // parenthesis (open)
		")":  "-.--.-",  // parenthesis (close)
		"()": "-.--.-",  // parentheses
		"$":  "...-..-", // dollar
		"&":  ".-...",   // ampersand
		"@":  ".--.-.",  // at
		"+":  ".-.-.",   // plus
		"Á":  ".--.-",   // A with acute accent
		"Ä":  ".-.-",    // A with diaeresis
		"É":  "..-..",   // E with acute accent
		"Ñ":  "--.--",   // N with tilde
		"Ö":  "---.",    // O with diaeresis
		"Ü":  "..--",    // U with diaeresis
	}
}

func init_reverse() {
	if _alphabet == nil {
		return
	}
	if _reverseAlphabet == nil {
		_reverseAlphabet = make(map[string]string)
	}
	for k, v := range _alphabet {
		_reverseAlphabet[v] = k
	}
}
