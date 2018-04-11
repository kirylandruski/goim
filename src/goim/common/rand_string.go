package common

import (
	"math/rand"
)

func GenerateRandomStrings(count int, itemLen int) []string {
	const testBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	gen := rand.New(rand.NewSource(0))

	res := make([]string, count)
	for i := range res {
		b := make([]byte, itemLen)
		for j := range b {
			b[j] = testBytes[gen.Intn(len(testBytes))]
		}
		res[i] = string(b)
	}

	return res
}
