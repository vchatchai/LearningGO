package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func runcrc32() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	fmt.Println(h.Sum32())
}

func runHash() {
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(h.Sum32())

}

func runSHA1() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

}
