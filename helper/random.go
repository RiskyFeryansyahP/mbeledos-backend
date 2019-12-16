package helper

import (
	"math/rand"
	"time"
)

func OTP() int32 {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Int31n(2000-1000+1) + 1000
	return otp
}
