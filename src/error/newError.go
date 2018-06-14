package error

import (
	"errors"
	"fmt"
)

func ReturnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func NewError() {
	err := ReturnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = ReturnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
