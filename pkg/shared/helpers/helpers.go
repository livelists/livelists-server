package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func FalseIfNil(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
