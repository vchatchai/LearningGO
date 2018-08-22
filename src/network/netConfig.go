package network

import (
	"fmt"
	"net"
	"reflect"
)

func NetConfig() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		adress, _ := i.Addrs()
		var ip net.IP

		for _, addr := range adress {
			fmt.Println(reflect.TypeOf(addr).String())
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println(ip.String())
		}
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		adresses, err := byName.Addrs()
		for k, v := range adresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}

}
