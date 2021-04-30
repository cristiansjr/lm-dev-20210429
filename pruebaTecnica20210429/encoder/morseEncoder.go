package encoder

import (
	"fmt"
	"strings"
)

type morseEncoder struct {}

func NewMorseEncoder() *morseEncoder { return &morseEncoder{}}

func (me *morseEncoder) EncodeMorse(text string) {
	textConv := convertToMorse(text)
	fmt.Println(textConv)
}

func convertToMorse(text string) interface{} {
	textToConvert := strings.ToUpper(text)
	var convertedText string
	tableConversionMorse := map[string]string {
		"A": ".- ",
		"B": "-... ",
		"C": "-.-. ",
		"D": "-.. ",
		"E": ". ",
		"F": "..-. ",
		"G": "--. ",
		"H": ".... ",
		"I": ".. ",
		"J": ".--- ",
		"K": "-.- ",
		"L": ".-.. ",
		"M": "-- ",
		"N": "-. ",
		"O": "--- ",
		"P": ".--. ",
		"Q": "--.- ",
		"R": "-. ",
		"S": "... ",
		"T": "- ",
		"U": "..- ",
		"W": ".-- ",
		"X": "-..- ",
		"Y": "-.-- ",
		"Z": "--.. ",
		" ": "   ",
	}
	for _, c := range textToConvert {
		convertedText = convertedText + tableConversionMorse[string(c)]
	}
	return convertedText
}