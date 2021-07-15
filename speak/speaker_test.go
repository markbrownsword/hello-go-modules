package speak_test

import (
    "github.com/markbrownsword/hello-go-modules/language"
	"github.com/markbrownsword/hello-go-modules/speak"
	"testing"
	"time"
)

func TestSpeaker(t *testing.T) {
	day := "Good Day"
	night := "Good Evening"
	english := language.NewEnglish(day, night)
	dayTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.UTC)
	nightTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 0, 0, 0, time.UTC)

	tests := []struct {
		speaker  speak.Speaker
		time     time.Time
		expected string
	}{
		{english, dayTime, day},
		{english, nightTime, night},
	}

	for _, test := range tests {
		actual := speak.Greeting(test.speaker, test.time)
		if actual != test.expected {
			t.Errorf("speak.Greeting(%s, %s): expected '%s', actual '%s'", test.speaker, test.time, test.expected, actual)
		}
	}
}
