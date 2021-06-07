package main

import (
	"context"
	"fmt"
	"testRedis/util"
	"time"
)

func main() {

	var ctx = context.Background()
	//
	//var rdb=util.NewRdb()
	//
	//rdb.Set(ctx,)

	var repos []string
	for i := 0; i < 1000; i++ {
		repos = append(repos, util.GetRandomGit(10))
	}

	var ips []string
	for i := 0; i < 8; i++ {
		ips = append(ips, util.GetRandomIp())

	}
	fmt.Printf("repos=%v,ips=%v\n", repos, ips)

	//获得一个redis 实例
	rdb := util.NewRdb()
	defer rdb.Close()
	//分发仓库
	l := len(ips)
	var ip = ""
	for index, repo := range repos {
		ip = ips[index%l]
		rdb.SAdd(ctx, ip, repo)
	}

	//获得分发后的数量
	nowRepos := make(map[string]int64, len(ips))
	for _, ip := range ips {
		num := rdb.SCard(ctx, ip)
		nowRepos[ip] = num.Val()
	}

	fmt.Printf("%v\n", nowRepos)

	//消耗仓库,并发消耗
	for _, ip := range ips {
		for i := 0; i < 30; i++ {
			rdb.SPop(ctx, ip)
		}
	}

	//等5秒
	time.Sleep(time.Second * 5)
	//获得分发后的数量
	for _, ip := range ips {
		num := rdb.SCard(ctx, ip)
		nowRepos[ip] = num.Val()
	}

	fmt.Printf("%v\n", nowRepos)
}
