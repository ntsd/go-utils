package stringutil

import (
	"math/rand"
)

// RandStringNumber random string number by giving length
func RandStringNumber(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = NUMBERS[rand.Int63()%int64(NUMBERS_LEN)]
	}
	return string(b)
}

// RandStringLetter random string upper letter by giving length
func RandStringLetter(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = LETTERS[rand.Int63()%int64(LETTERS_LEN)]
	}
	return string(b)
}

// RandStringUpperLetter random string upper letter by giving length
func RandStringUpperLetter(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = UPPERCASE_LETTERS[rand.Int63()%int64(UPPERCASE_LETTERS_LEN)]
	}
	return string(b)
}

// RandStringLowerLetter random string upper letter by giving length
func RandStringLowerLetter(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = LOWERCASE_LETTERS[rand.Int63()%int64(LOWERCASE_LETTERS_LEN)]
	}
	return string(b)
}
