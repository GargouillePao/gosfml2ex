package utils

import (
	"fmt"
)

type util struct {
	Math     mathUtil
	Vector   vectorUtil
	Graphics graphicsUtil
}

func Utils() util {
	return util{
		Math:     mathUtil{},
		Vector:   vectorUtil{},
		Graphics: graphicsUtil{},
	}
}

type errors struct {
	NotTheSameType string
	CannotMius     string
	NilAttribute   string
	OutOfRange     string
}

func Errors() errors {
	return errors{
		NotTheSameType: "NotTheSameTypeError",
		CannotMius:     "CannotMiusError",
		NilAttribute:   "NilAttributeError",
		OutOfRange:     "OutOfRangeError",
	}
}

type ErrorUtil struct {
	err string
}

func (u ErrorUtil) Error() string {
	return fmt.Sprintf("Chatch error:%s", u.err)
}
