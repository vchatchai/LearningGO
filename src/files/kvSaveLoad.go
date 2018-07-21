package files

import (
	"encoding/gob"
	"fmt"
	"os"
)

func SaveGOB(file string, data interface{}) (err error) {
	fmt.Println("Saving", file)
	err = os.Remove(file)

	saveTo, err := os.Create(file)
	if err != nil {
		return
	}

	defer saveTo.Close()
	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(data)
	if err != nil {
		return
	}

	return nil
}

func LoadGOB(file string) (data string, err error) {
	fmt.Println("Loading", file)
	loadFrom, err := os.Open(file)
	defer loadFrom.Close()

	if err != nil {
		return
	}
	decoder := gob.NewDecoder(loadFrom)
	err = decoder.Decode(&data)

	return
}
