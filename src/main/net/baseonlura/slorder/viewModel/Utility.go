package viewModel

import (
	"encoding/json"
	"time"
)

var DATE_FORMAT = "2006-01-02"
var DATE_TIME_MILLIS_FORMAT = "2006-01-02 15:04:05.000"
var DATE_TIME_FORMAT = "2006-01-02 15:04:05"
var NOW_DATE = time.Now().UTC().Truncate(time.Hour * 24)
var NOW_DATE_TIME = time.Now().UTC()

func DeepCopy(src interface{}, dst interface{}) error {
	byteData, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &dst)
	if err != nil {
		return err
	}
	return nil
}

func Str2Date(dateStr string) time.Time {
	date, err := time.Parse(DATE_FORMAT, dateStr)
	if err != nil {
		return time.Time{}
	}
	return date
}
func Str2DateTime(dateTime string) time.Time {
	date, err := time.Parse(DATE_TIME_MILLIS_FORMAT, dateTime)
	if err != nil {
		return time.Time{}
	}
	return date
}
