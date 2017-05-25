package proto

import (
	"errors"
	"strconv"
)

func (s *SayHelloRequest) SetYourName(v string) error {
	// NotBlank
	if len(v) == 0 {
		return errors.New("validate fail! NotBlank, expected:<=5, actual:" + strconv.Itoa(len(v)))
	}
	// MaxLength
	if len(v) > 5 {
		return errors.New("validate fail! MaxLength, expected:<=5, actual:" + strconv.Itoa(len(v)))
	}
	s.YourName = v
	return nil
}
