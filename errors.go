// Package multiple provides a container for collecting multiple errors.
package multiple

import "strings"

// Errors is a list of multiple errors.
//
// To use it as an error, call Err().
type Errors []error

// Append err to e only if err != nil.
//
// If err is itself Errors, each individual error in err is appended.
//
// Append returns err.
func (e *Errors) Append(err error) error {
	if e == nil {
		panic("cannot append to nil multiple.Errors")
	}
	if err == nil {
		return nil
	}
	if errs, ok := err.(Errors); ok {
		*e = append(*e, errs...)
		return err
	}
	*e = append(*e, err)
	return err
}

// First returns the first error in e or nil if there are none.
func (e Errors) First() error {
	// Return first nonnil error.
	for _, err := range e {
		if err != nil {
			return err
		}
	}
	return nil
}

// Err returns e as an error or returns nil if no errors were collected.
func (e Errors) Err() error {
	// Only return self if we have at least one nonnil error.
	for _, err := range e {
		if err != nil {
			return e
		}
	}
	return nil
}

// Error returns an error string by joining
// all of its nonnil errors with newlines.
func (e Errors) Error() string {
	if len(e) == 0 {
		return ""
	}
	es := make([]string, 0, len(e))
	for _, err := range e {
		if err != nil {
			es = append(es, err.Error())
		}
	}
	return strings.Join(es, "\n")
}
