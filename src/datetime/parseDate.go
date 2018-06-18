package datetime

import (
	"fmt"
	"time"
)

func ParseDate(myDate string) (date time.Time, err error) {

	date, err = time.Parse("02 January 2006", myDate)

	if err == nil {
		fmt.Println("Full", date)
		fmt.Println("Time", date.Day(), date.Month(), date.Year())
	} else {
		fmt.Println(err)
	}

	return
}
