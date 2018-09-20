package network

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

const welcome = "Welcome to the Key-value store!\n"

var DATA = make(map[string]myElement)
var DATAFILE = "/tmp/dataFile.gob"

func KVTCP(address string) {

	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	err = load()
	if err != nil {
		fmt.Println(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		go handleConnectionTCP(c)
	}

}

func handleConnectionTCP(c net.Conn) {
	c.Write([]byte(welcome))
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		command := strings.TrimSpace(string(netData))
		tokens := strings.Fields(command)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")

		}

		switch tokens[0] {
		case "STOP":
			err = save()
			if err != nil {
				fmt.Println(err)
			}
			c.Close()
			return

		case "PRINT":
			PRINT(c)
		case "DELETE":
			if !DELETE(tokens[1]) {
				netData := "Delete operation failed!\n"
				c.Write([]byte(netData))
			} else {
				netData := "Delete operation successful!\n"
				c.Write([]byte(netData))
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}

			if !ADD(tokens[1], n) {
				netData := "Add operation failed]n"
				c.Write([]byte(netData))
			} else {
				netData := "Add operation successful!\n"
				c.Write([]byte(netData))
			}
			err = save()
			if err != nil {
				fmt.Println(err)
			}

		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				netData := fmt.Sprintf("%v\n", *n)
				c.Write([]byte(netData))
			} else {
				netData := "Did not find key!\n"
				c.Write([]byte(netData))
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				netData = "Update operation failed!\n"
				c.Write([]byte(netData))
			} else {
				netData = "Update operation successful!\n"
				c.Write([]byte(netData))
			}
			err = save()
			if err != nil {
				fmt.Println(err)
			}
		default:
			netData := "Unknow command - please try again!\n"
			c.Write([]byte(netData))
		}
	}

}

func CHANGE(key string, data myElement) bool {
	DATA[key] = data
	return true
}

func LOOKUP(key string) *myElement {
	data, ok := DATA[key]

	if ok {
		return &data
	} else {
		return nil
	}
}

func DELETE(key string) bool {

	if LOOKUP(key) != nil {
		delete(DATA, key)
		return true
	}

	return false
}
func PRINT(c net.Conn) {
	for k, d := range DATA {
		netData := fmt.Sprintf("key: %s value: %v\n", k, d)
		c.Write([]byte(netData))
	}
}

func ADD(k string, data myElement) bool {

	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = data
		return true
	}

	return false

}

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)

	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}

	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	dataFile, err := os.Open(DATAFILE)
	defer dataFile.Close()

	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(dataFile)
	decoder.Decode(&DATA)
	return nil
}
