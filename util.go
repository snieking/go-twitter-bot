// Convenient utility functions.
package main

import (
	"log"
	"math/rand"
	"time"
)

// Makes a timestamp at the current time and returns it in milliseconds.
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Makes a timestamp at current time, minus provided amount of hours.
func makeTimestampHoursBeforeNow(hours int) int64 {
	return makeTimestamp() - (int64(hours) * 3600000)
}

// Returns a pointer to a true boolean.
func newTrue() *bool {
	b := true
	return &b
}

// Returns a pointer to a false boolean.
func newFalse() *bool {
	b := false
	return &b
}

// Checks for errors, and if there is an error then it logs it as a Fatal
// together with a provided string.
// This method will kill the application if the error exists.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Removes the UserEntity with the provided index from the slice.
// It doesn't retain the order as the last element will be put at
// the index of the removed element.
func remove(s []UserEntity, i int) []UserEntity {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// Returns a random element from the slice.
func randomElementFromSlice(s []string) string {
	return s[randomNumberInRange(0, len(s))]
}

// Returns a random number in a provided range.
// For example a random number between 10-15.
func randomNumberInRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
