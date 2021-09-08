package config

import (
	"fmt"
)

type MissingValue struct {
	Val interface{}
}

func (s *MissingValue) Error() string {
	return fmt.Sprintf("Provided value is not set? %#v.", s.Val)
}

type UnsupportedValue struct {
	Val interface{}
}

func (s *UnsupportedValue) Error() string {
	return fmt.Sprintf("Provided value is unsupported: %#v.", s.Val)
}
