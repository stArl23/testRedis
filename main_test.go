package main

import (
	"fmt"
	"testRedis/util"
)

func main() {

	var repos []string
	for i := 0; i < 100; i++ {
		repos = append(repos, util.GetRandomGit(10))
	}

	var ips []string
	for i := 0; i < 100; i++ {
		ips = append(ips, util.GetRandomIp())

	}
	fmt.Printf("repos=%v\n", repos)

}
