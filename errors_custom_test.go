package multiple_test

import (
	"errors"
	"fmt"

	"github.com/jimmyfrasche/multiple"
)

type Errors struct {
	multiple.Errors
}

func NewErrors(errs multiple.Errors) error {
	if errs.Err() == nil {
		return nil
	}
	return Errors{errs}
}

func (e Errors) Error() string {
	switch len(e.Errors) {
	case 0:
		return ""
	case 1:
		return e.First().Error()
	default:
		return fmt.Sprintf("%s (and %d more)", e.First(), len(e.Errors)-1)
	}
}

func ExampleErrors_customFormat() {
	var errs multiple.Errors
	errs.Append(errors.New("an error"))
	errs.Append(errors.New("another error"))
	fmt.Println(NewErrors(errs))

	// Output:
	// an error (and 1 more)
}
