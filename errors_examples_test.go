package multiple_test

import (
	"errors"
	"fmt"

	"github.com/jimmyfrasche/multiple"
)

func NoError() error {
	return nil
}

func Error() error {
	return errors.New("an error")
}

func ExampleErrors() {
	var errs multiple.Errors

	fmt.Println("Adding nil error")
	errs.Append(NoError())
	fmt.Println(len(errs), errs.Err() == nil)
	fmt.Println(errs)
	fmt.Println("first:", errs.First())

	fmt.Println()
	fmt.Println("Adding first error")
	errs.Append(Error())
	fmt.Println(len(errs), errs.Err() == nil)
	fmt.Println(errs)
	fmt.Println("first:", errs.First())

	fmt.Println()
	fmt.Println("Adding second error")
	errs.Append(Error())
	fmt.Println(len(errs), errs.Err() == nil)
	fmt.Println(errs)
	fmt.Println("first:", errs.First())

	// Output:
	// Adding nil error
	// 0 true
	//
	// first: <nil>
	//
	// Adding first error
	// 1 false
	// an error
	// first: an error
	//
	// Adding second error
	// 2 false
	// an error
	// an error
	// first: an error
}
