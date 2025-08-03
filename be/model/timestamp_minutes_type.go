package model

import (
	"encoding/json"
	"time"
)

type TimestampMinutes time.Time

func (tm TimestampMinutes) MarshalJSON() ([]byte, error) {
	t := time.Time(tm)
	return json.Marshal(t.Format("2006-01-02T15:04"))
}

func (tm *TimestampMinutes) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02T15:04", s)
	if err != nil {
		return err
	}
	*tm = TimestampMinutes(t)
	return nil
}
