package main

import (
	"fmt"
	"os"
	"reflect"
)

/* error types */
type onlyOneArgNegError struct {
	neg int
	err string
}

type prodUnderFiftyError struct {
	x, y, prod int
	err        string
}

type oneArgAboveFiftyError struct {
	above int
	err   string
}

func main() {

	// call the function and print error if any - type switch to print the errpr type
	res, err := productAboveFiftyOnly(2, 16)
	if err != nil {
		switch err.(type) {
		case *onlyOneArgNegError:
			fmt.Println(reflect.TypeOf(err))
		case *prodUnderFiftyError:
			fmt.Println(reflect.TypeOf(err))
		case *oneArgAboveFiftyError:
			fmt.Println(reflect.TypeOf(err))
		}
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

}

/* function mapping all types of error to each error situation*/
func productAboveFiftyOnly(x, y int) (int, error) {
	// only one negative arg
	if x < 0 {
		err := onlyOneArgNegError{neg: x}
		return 0, &err
	}
	if y < 0 {
		err := onlyOneArgNegError{neg: y}
		return 0, &err
	}

	// product under fifty
	if prod := x * y; prod <= 50 {
		err := prodUnderFiftyError{
			x:    x,
			y:    y,
			prod: x * y,
		}
		return 0, &err
	}

	// one arg is above fifty
	if x > 50 {
		err := oneArgAboveFiftyError{above: x}
		return 0, &err
	}
	if y > 50 {
		err := oneArgAboveFiftyError{above: y}
		return 0, &err
	}

	return x * y, nil
}

/* Error methods for each error type */
func (e *onlyOneArgNegError) Error() string {
	e.err = fmt.Sprintf("only one argument %d is negative", e.neg)
	return e.err
}

func (e *prodUnderFiftyError) Error() string {
	e.err = fmt.Sprintf("product %d of %d and %d is under 50", e.prod, e.x, e.y)
	return e.err
}

func (e *oneArgAboveFiftyError) Error() string {
	e.err = fmt.Sprintf("argument %d is above 50", e.above)
	return e.err
}
