package cerrors

import (
	"fmt"
)

type ErrValidation struct {
	Issue error // to report issue plus cause

	Caller         string
	WithoutLineEnd bool
}

const areaErrValidation = "Validation"

func (e ErrValidation) Error() string {
	var res [3]string

	res[0] = fmt.Sprintf("Area: %s", areaErrValidation)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Issue: %s", e.Issue.Error()) // possible panic here

	if e.WithoutLineEnd {
		return res[0] + _space + res[1] + _space + res[2]
	}

	return "\n" + res[0] + "\n" + res[1] + "\n" + res[2] + "\n"
}

func (e ErrValidation) Is(err error) bool {
	_, ok := err.(ErrValidation)

	return ok
}
