package models

import (
	"strings"
	"time"
)

var (
	layout = "2006-01-02 00:00"
)

type Date struct{ time.Time }

func (c *Date) UnmarshalJSON(b []byte) (err error) {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		c.Time = time.Time{}
		return nil
	}
	c.Time, err = time.Parse(layout, value) //parse time
	return err
}
