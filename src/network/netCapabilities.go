package network

import (
	"fmt"
	"net"
)

func NetCapabilities() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		fmt.Println("Interface HW Address", i.HardwareAddr)
		fmt.Println()
	}
}
