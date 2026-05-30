package logger

import (
	"fmt"
)


func FormatError(component string, reason string) error {
	err := fmt.Errorf("%v error: %v.", component, reason)
	return err
}