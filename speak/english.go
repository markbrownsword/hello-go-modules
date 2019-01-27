package speak

import "time"

type English struct {
	day   string
	night string
}

func NewEnglish(day string, night string) *English {
	english := new(English)
	english.day = day
	english.night = night

	return english
}

func (english *English) Speak(now time.Time) string {
	hour := now.Hour()
	isNightTime := hour > 17 // 5PM

	if isNightTime {
		return english.night
	}

	return english.day
}
