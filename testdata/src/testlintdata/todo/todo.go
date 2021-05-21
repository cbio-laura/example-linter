package todo

import (
	"fmt"
)

type myErrors struct {
}

func (e myErrors) New(code int64, text string) error {
	return nil
}

func errorFx(code int64, childErr error, text string) error {
	return nil
}

// want "Error number 1001 has already been seen"
func SomeFunc1() {
	cberr := errorFx(1001, nil, "Sample error")
	cberr = errorFx(1001, nil, "Sample error")
	fmt.Printf("Error: %s", cberr)
}

// want "Error number 1002 has already been seen"
func SomeFunc2() {
	errors := myErrors{}

	cberr := errors.New(1002, "Sample error")
	cberr = errors.New(1002, "Sample error")
	fmt.Printf("Error: %s", cberr)
}
