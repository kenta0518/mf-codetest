package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return err
	}

	tt, err := time.ParseInLocation("2006/01/02 15:04:05", strings.TrimSpace(buf), time.Local)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t *DateTime) MarshalYAML() (interface{}, error) {
	return t.Time.Format("2006/01/02 15:04:05"), nil
}

func (t *DateTime) Scan(value interface{}) error {
	tm, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}

	*t = DateTime{Time: tm}
	return nil
}

func (t DateTime) Value() (driver.Value, error) {
	return t.Time, nil
}
