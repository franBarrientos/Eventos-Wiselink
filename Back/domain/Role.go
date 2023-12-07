package domain

import (
	"database/sql/driver"
	"fmt"
)

type Role string

const (
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

func (self *Role) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*self = Role(string(v))
	case string:
		*self = Role(v)
	default:
		return fmt.Errorf("unexpected type for Role: %T", value)
	}
	return nil
}

func (self Role) Value() (driver.Value, error) {
	return string(self), nil
}
