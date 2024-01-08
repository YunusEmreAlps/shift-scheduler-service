package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// for support jsonb in postgres, create interface
type JSONB []interface{}

// Value Marshal
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan Unmarshal
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j) // Remove the "&" here
}
