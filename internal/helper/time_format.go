package helper

import (
	"encoding/base64"
	"time"
)

// DecodeTime will decode time
func DecodeTime(encodedTime string, timeFormat string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeTime will encode time
func EncodeTime(t time.Time, timeFormat string) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
