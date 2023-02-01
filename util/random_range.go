package util

import (
	"math/rand"
	"time"
)

/*
 * Function: RandomRange
 * Generate a random number greater than or equal to min and less than max
 *
 * Params:
 * min: Minimum number in the range
 * max: Maximum number in the range + 1
 *
 * Return:
 * Random number greater than or equal to min and less than max
 */
func RandomRange(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}
