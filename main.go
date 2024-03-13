package main

import (
	"fmt"
	"net"
	"strings"
	"ui"
	"web"

	"github.com/gin-gonic/gin"
)

var ip = ""

func main() {
	ip, e := my_ip_address()
	if e != nil {
		return
	}
	gin.SetMode(gin.ReleaseMode)

	go web.Start()
	ui.Start(ip)
}

func my_ip_address() ([]string, error) {
	n := make([]string, 0)
	iface, e := net.InterfaceAddrs()
	if e != nil {
		return nil, e
	}
	for _, i := range iface {
		ip := i.String()
		if strings.Contains(ip, "/") {
			a := strings.Split(ip, "/")
			if len(a) > 0 {
				if strings.HasPrefix(a[0], "::") || strings.HasPrefix(a[0], "127") {
					continue
				}
				n = append(n, "http://"+a[0]+":8083/")
			}
		}
	}
	fmt.Println(n)
	return n, nil
}
