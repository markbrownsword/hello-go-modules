// Package speak implements greeting functionality
package speak

import "time"

type Speaker interface {
	Speak(now time.Time) string
}

func Greeting(s Speaker, now time.Time) string {
	return s.Speak(now)
}
