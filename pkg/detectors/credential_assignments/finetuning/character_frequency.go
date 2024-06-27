package finetuning

import (
	"math"
	"unicode"
)

func CheckProperCharacterFrequencyDistribution(fileExtension string, variableName string, secretValue string) bool {
	statistic := getCharacterStatistics(secretValue)

	// Password is probably a placeholder when the complete password does not contain a single lower case letter or number
	if statistic.countOfLowerCaseCharacters == 0 && statistic.countOfNumberCharacters == 0 {
		return false
	}

	// if a character is very frequent then it is probably not a password
	// the threshold 40% is very high to be on the safe side to not prune out real passwords
	for _, charCount := range statistic.countPerCharacter {
		if float32(charCount)/float32(statistic.totalNumberOfCharacters) > 0.4 {
			return false
		}
	}

	return true
}

func getCharacterStatistics(value string) *characterStatistics {
	var lowerCase, upperCase, numbers, specialChars int

	countPerCharacter := make(map[rune]int)
	observedSpecialCharacters := make(map[rune]struct{})

	for _, char := range value {
		switch {
		case unicode.IsLower(char):
			lowerCase++
		case unicode.IsUpper(char):
			upperCase++
		case unicode.IsDigit(char):
			numbers++
		default:
			observedSpecialCharacters[char] = struct{}{}
			specialChars++
		}

		if _, found := countPerCharacter[char]; found {
			countPerCharacter[char]++
		} else {
			countPerCharacter[char] = 1
		}
	}

	return &characterStatistics{
		countOfLowerCaseCharacters:            lowerCase,
		countOfUpperCaseCharacters:            upperCase,
		countOfNumberCharacters:               numbers,
		countOfSpecialCharacters:              specialChars,
		totalNumberOfCharacters:               len(value),
		countPerCharacter:                     countPerCharacter,
		numberOfDifferentObservedSpecialChars: len(observedSpecialCharacters),
	}
}

type characterStatistics struct {
	countOfLowerCaseCharacters            int
	countOfUpperCaseCharacters            int
	countOfNumberCharacters               int
	countOfSpecialCharacters              int
	totalNumberOfCharacters               int
	numberOfDifferentObservedSpecialChars int
	countPerCharacter                     map[rune]int
}

func (statistic *characterStatistics) GetShannonEntropy() float64 {
	entropy := 0.0
	for _, count := range statistic.countPerCharacter {
		p := float64(count) / float64(statistic.totalNumberOfCharacters)
		entropy -= p * math.Log2(p)
	}
	return entropy
}
