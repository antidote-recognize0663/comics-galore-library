package utils

func GenerateRandomInt(min, max int) int {
	if min < 0 {
		min = 0
	}
	// Validate the range
	if max <= min {
		max = min + 10
		//panic("max must be greater than min")
	}
	// Use the globally seeded random generator
	return seededRand.Intn(max-min) + min
}

func GenerateRandomInt64(min, max int64) int64 {
	if min < 0 {
		min = 0
	}
	// Validate the range
	if max <= min {
		max = min + 10
		//panic("max must be greater than min")
	}
	// Use the globally seeded random generator
	return seededRand.Int63n(max-min) + min
}
