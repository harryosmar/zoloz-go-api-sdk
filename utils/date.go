package utils

import (
	"fmt"
	"regexp"
	"time"
)

func GenerateZolozTimeStr(t time.Time) string {
	str := t.Format(time.RFC3339)
	return GenerateZolozTimeStrFromStr(str)
}

func GenerateZolozTimeStrFromStr(str string) string {
	m := regexp.MustCompile(`^(.*)Z$`)
	matches := m.FindStringSubmatch(str)
	if len(matches) == 2 {
		return fmt.Sprintf("%s+0000", matches[1])
	}

	m = regexp.MustCompile(`(\d{2}):(\d{2})$`)
	str = m.ReplaceAllString(str, "$1$2")
	return str
}
