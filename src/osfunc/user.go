package osfunc

import (
	"fmt"
	"os"
	"os/user"
)

func User() {
	fmt.Println(os.Geteuid())

	var u *user.User
	u, _ = user.Current()
	fmt.Print("Group ids: ")
	groupIDs, _ := u.GroupIds()
	for _, id := range groupIDs {
		fmt.Print(id, " ")
	}
	fmt.Println()
}
