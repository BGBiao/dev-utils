package main

import (
	"fmt"
	"os/user"
	"strings"
)

func main() {
	cUser, _ := user.Current()

	fmt.Println(cUser.Uid, cUser.Gid, cUser.Username, cUser.Name, cUser.HomeDir)
	gids, _ := cUser.GroupIds()
	fmt.Printf("current user ids:%v\n", strings.Join(gids, ","))

	root, _ := user.Lookup("root")
	// user.LookupId(id string)
	fmt.Println(root.Uid, root.Gid, root.Username, root.Name, root.HomeDir)
}
