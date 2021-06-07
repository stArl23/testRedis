package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var asciis = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

//返回随机的git地址
func GetRandomGit(needLen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(asciis)
	var result = ""
	for i := 0; i < needLen; i++ {
		result += string(asciis[r.Int()%l])
	}
	return "https://github.com/" + result + ".git"
}

//获取随机ip
func GetRandomIp() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result = ""
	for i := 0; i < 4; i++ {
		var s = "."
		result += strconv.Itoa(r.Int() % 255)
		if i != 3 {
			result += s
		}
	}
	return result
}

//直接使用10进制数字作为key
func Ip2Key(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0, nil
	}
	var res = 0
	for index, ip := range ips {
		tmp, _ := strconv.Atoi(ip)
		res += tmp
		if index != 3 {
			res *= 256
		}
	}
	return res, nil
}
