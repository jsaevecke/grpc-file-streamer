package cerrors

import "fmt"

type ErrConfiguration struct {
	Issue error

	Caller  string
	Calling string
}

const areaErrConfiguration = "Configuration"

func (e *ErrConfiguration) Error() string {
	var res [4]string

	res[0] = fmt.Sprintf("\nArea: %s", areaErrConfiguration)
	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Calling: %s", e.Calling)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
