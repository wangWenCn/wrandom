package wrandom

import (
	"errors"
	"math/rand"
	"time"
)

func GenerateRandomDigitsWithLen(count, length int) []string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	codes := make([]string, count)
	for i := range codes {
		digits := make([]rune, length)
		for j := range digits {
			digits[j] = rune(rng.Intn(10) + '0')
		}
		codes[i] = string(digits)
	}
	return codes
}

func GenerateBytesList(l int) ([]byte, error) {
	if l <= 0 {
		return nil, errors.New("list len is zero")
	}
	keys := make([]byte, l)
	_, err := rand.Read(keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
