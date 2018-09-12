package network

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

func TCPclient(address, text string) {

	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		var buffer bytes.Buffer
		buffer.WriteString(text)
		buffer.WriteString("\n")
		fmt.Fprintf(c, buffer.String())
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
