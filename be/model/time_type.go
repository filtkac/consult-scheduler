package model

// Custom type for GORM and for JSON (un)marshaling only time values in format HH:mm
// adopted from https://jem72.medium.com/custom-time-types-in-go-with-gorm-and-gin-c2da116f1f46

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type TimeOnly struct {
	time.Time
}

func (TimeOnly) GormDataType() string {
	return "time"
}

func (TimeOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

func (timeOnly TimeOnly) Value() (driver.Value, error) {
	if !timeOnly.IsZero() {
		return timeOnly.GetTime().Format("15:04:05"), nil
	} else {
		return nil, nil
	}
}

func (timeOnly *TimeOnly) GetTime() time.Time {
	return timeOnly.Time
}

func (timeOnly *TimeOnly) Scan(value interface{}) error {
	scannedString := value.(string)
	scannedTime, err := time.Parse("15:04:05", scannedString)
	if err == nil {
		*timeOnly = TimeOnly{scannedTime}
	}
	return err
}

func (timeOnly TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(timeOnly.GetTime().Format("15:04"))
}

func (timeOnly *TimeOnly) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("15:04", s, time.UTC)
	if err != nil {
		return err
	}
	*timeOnly = TimeOnly{t}
	return nil
}
