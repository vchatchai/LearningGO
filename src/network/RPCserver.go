package network

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
)

type RPCInterface struct {
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *RPCInterface) Multiply(arguments *MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *RPCInterface) Power(arguments *MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

func RPCServer(address string) {
	myInterface := new(RPCInterface)
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return

	}

	handler := func(l net.Conn) {
		fmt.Printf("%s\n", l.RemoteAddr())
		rpc.ServeConn(l)
	}

	for {
		con, err := l.Accept()
		if err != nil {
			continue
		}

		go handler(con)
	}

}
