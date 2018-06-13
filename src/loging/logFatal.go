package loging

import (
	"fmt"
	"log"
	"log/syslog"
)

func FatalLog() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "FatalLogAPP")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(sysLog)
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}
