package panicRecover

import "fmt"

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recovere inside a()!", c)
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b")
	panic("Panic in b()!")
	fmt.Println("Exitingb()")
}

func PanicAndRecover() {
	a()
	fmt.Println("PanicAndRecover ended!")
}
