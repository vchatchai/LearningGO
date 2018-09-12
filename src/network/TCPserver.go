package network

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

func TCPserver(port string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("->", string(netData))
		t := time.Now()

		var buffer bytes.Buffer

		myTime := t.Format(time.RFC3339)
		buffer.WriteString(myTime)
		buffer.WriteByte('\n')

		c.Write(buffer.Bytes())
		fmt.Println("end loop")
	}

}
