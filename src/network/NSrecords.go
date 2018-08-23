package network

import (
	"fmt"
	"net"
)

func NSrecords(domain string) {
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
