package random

import (
	"math/rand"
	"time"
)

const (
	alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"
)

func NewRandomString(size int) string {
	res := make([]byte, size)

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < size; i++ {
		randomIndex := random.Intn(len(alphabet))
		res[i] = alphabet[randomIndex]
	}
	return string(res)
}
