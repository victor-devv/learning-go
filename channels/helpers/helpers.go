package helpers

import (
	"math/rand"
	"time"
)

//create random number from pool of size n
func RandomNumber (n int) int {
	//seed
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)

	return value
}
