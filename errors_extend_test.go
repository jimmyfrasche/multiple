package multiple_test

import (
	"fmt"

	"github.com/jimmyfrasche/multiple"
)

func ExampleErrors_extend() {
	var errs1, errs2 multiple.Errors

	for i := 0; i < 2; i++ {
		errs1.Append(fmt.Errorf("errs1: error %d", i+1))
	}
	for i := 0; i < 3; i++ {
		errs2.Append(fmt.Errorf("errs2: error %d", i+1))
	}

	errs1.Append(errs2)

	for _, err := range errs1 {
		fmt.Println(err)
	}
	// Output:
	// errs1: error 1
	// errs1: error 2
	// errs2: error 1
	// errs2: error 2
	// errs2: error 3
}
