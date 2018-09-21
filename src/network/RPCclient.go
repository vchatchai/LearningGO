package network

import (
	"fmt"
	"net/rpc"
)

func RPCClient(address string) {
	c, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := MyFloats{16, -0.5}

	var reply float64

	err = c.Call("RPCInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("RPCInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
