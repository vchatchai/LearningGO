package datetime

import (
	"fmt"
	"time"
)

func ParseTime(mytime string) (d time.Time, err error) {
	d, err = time.Parse("15:04", mytime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute(), d.Second())

	} else {
		fmt.Println(err)
	}

	return
}
