package loging

import (
	"fmt"
	"log"
	"path/filepath"
	"os"
	"log/syslog"
)

func Syslog(){
	programName := filepath.Base(os.Args[0])

	sysLog, err := syslog.New(syslog.LOG_INFO , programName)

	if err != nil {
		log.Fatal(err)
	}else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL: Logging in GO")
	fmt.Println("will you see this ?")

}