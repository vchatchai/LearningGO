package osfunc

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func RequiredVersion() {
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]

	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	fmt.Println("myVersion:", myVersion)
	fmt.Println("major:", major)
	fmt.Println("minor", minor)
	fmt.Println("m1", m1)
	fmt.Println("m2", m2)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go Version 1.8 or higher!")
	}

	fmt.Println("You are using GO version 1.8 or higher")
}
