package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	connURL := "admin/admin@192.168.101.248:12084"
	hosts, ports, users, pwds := ParseConnURL(connURL)

	for i, s := range hosts {
		fmt.Println(i, s, ports[i], users[i], pwds[i])
	}

}

//parse conn url
func ParseConnURL(connStr string) (
	hosts []string, ports []int, users []string, pwds []string) {

	sl := len(connStr)

	if sl == 0 {
		return hosts, ports, users, pwds
	}

	ss := strings.Split(connStr, ";")

	size := len(ss)

	hosts = make([]string, size)
	ports = make([]int, size)
	users = make([]string, size)
	pwds = make([]string, size)

	startIndex := 0
	aIndex := 0

	for index, r := range connStr {
		switch r {
		case '/':
			users[aIndex] = connStr[startIndex:index]
			startIndex = index + 1
		case '@':
			pwds[aIndex] = connStr[startIndex:index]
			startIndex = index + 1
		case ':':
			hosts[aIndex] = connStr[startIndex:index]
			startIndex = index + 1
		case ';':
			ports[aIndex],_ = strconv.Atoi(connStr[startIndex:index])
			startIndex = index + 1
			aIndex++
		default:
			continue
		}
	}

	if startIndex < sl {
		ports[aIndex],_ = strconv.Atoi(connStr[startIndex:sl])
	}

	return hosts, ports, users, pwds

}
