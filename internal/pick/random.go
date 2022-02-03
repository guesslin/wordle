package pick

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Pick(pool []string) string {
	return pool[rand.Intn(len(pool))]
}
