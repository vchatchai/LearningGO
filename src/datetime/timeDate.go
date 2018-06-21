package datetime

import (
	"fmt"
	"regexp"
	"time"
)

func TimeDate() {
	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200]325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] \"GET /CVEN		HTTP/1.1\" 200 12531 \"-\" \"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
	}

	r := regexp.MustCompile(`.*\[(\d\d\/\w+\/\d\d\d\d:\d\d:\d\d:\d\d\s+\+\d\d\d\d.*)\].*`)
	for _, logEntry := range logs {
		fmt.Println(logEntry)
		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)
			fmt.Println("match:", match)
			for i, m := range match {
				fmt.Println("m", i, m)
			}
			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("Not a valid date time format!")
			}
			fmt.Println("Date", dt)

		} else {
			fmt.Println("Not a match!")
		}
	}
}
