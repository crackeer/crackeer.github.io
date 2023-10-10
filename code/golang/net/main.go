package main

import (
	"context"
	"fmt"
	"net"
)

func main() {
	ips, err := net.DefaultResolver.LookupIPAddr(context.Background(), "baidu.com")
	fmt.Println(ips, err)
}
