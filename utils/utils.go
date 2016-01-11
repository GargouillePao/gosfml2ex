package utils

import (
	"fmt"
)

type util struct {
	Math     mathUtil
	Vector   vectorUtil
	Graphics graphicsUtil
	Slice    sliceUtil
}

func Utils() util {
	return util{
		Math:     mathUtil{},
		Vector:   vectorUtil{},
		Graphics: graphicsUtil{},
		Slice:    sliceUtil{},
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

type errorUtil struct {
	err string
}

func NewError(err string) errorUtil {
	return errorUtil{err: err}
}

func (u errorUtil) Error() string {
	return fmt.Sprintf("Chatch error:%s", u.err)
}
