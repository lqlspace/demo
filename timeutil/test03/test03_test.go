package test03

import (
	"fmt"
	"testing"
	"time"
)


func Test03(t *testing.T) {
	now := UTCNowRFC3339()

	fmt.Printf("now = %s\n", now)

	nowtime, _ := ParseUTC3339(now)

	mill := nowtime.UnixMilli()

	fmt.Printf("mill = %d\n", mill)

	now = UTCNowRFC3339Nano()

	fmt.Printf("now = %s\n", now)

	nowtime, _ = ParseUTC3339Nano(now)
	mill = nowtime.UnixMilli()
	fmt.Printf("mill = %d\n", mill)
}

func UTCNowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}
func UTCNowRFC3339Nano() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

func ParseUTC3339(utcS string) (time.Time, error) {
	utcT, err := time.Parse(time.RFC3339, utcS)
	if err != nil {
		return time.Time{}, err
	}
	return utcT, nil
}

func ParseUTC3339Nano(utcS string) (time.Time, error) {
	utcT, err := time.Parse(time.RFC3339Nano, utcS)
	if err != nil {
		return time.Time{}, err
	}
	return utcT, nil
}