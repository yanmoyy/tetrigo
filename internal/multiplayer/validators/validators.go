package validators

import (
	"errors"
	"fmt"
	"strconv"
)

type ValidateFunc func(string) error

func Nickname() ValidateFunc {
	return func(s string) error {
		if len(s) == 0 {
			return errors.New("empty username not allowed")
		}
		return nil
	}
}

func RoomID() ValidateFunc {
	return func(s string) error {
		if len(s) == 0 {
			return errors.New("empty host address not allowed")
		}
		return nil
	}
}

func Port() ValidateFunc {
	return func(s string) error {
		if len(s) == 0 {
			return errors.New("empty port not allowed")
		}
		if _, err := strconv.Atoi(s); err != nil {
			return fmt.Errorf("invalid port: %w", err)
		}
		return nil
	}
}
