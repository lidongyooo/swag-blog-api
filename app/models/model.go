package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type BaseModel struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id" uri:"id"`

	CreatedAt LocalTime `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt LocalTime `gorm:"column:updated_at;index;not null" json:"updated_at"`
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return errors.New(fmt.Sprintf("can not convert %v to timestamp", v))
}