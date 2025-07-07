package utils

import (
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")

	if str == "" || str == "null" {
		d.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + d.Time.Format(dateLayout) + "\""), nil
}
