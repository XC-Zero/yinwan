package errs

import "github.com/pkg/errors"

func ErrorListToError(errorList []error) error {
	message := ""
	for i := range errorList {
		err := errorList[i]
		if err != nil {
			message += err.Error() + "\n"
		}
	}
	if message == "" {
		return nil
	}
	return errors.New(message)
}
