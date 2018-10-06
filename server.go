package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	port     = flag.Int("port", 1113, "set the port")
	rootPath = flag.String("root", "./", "set the root dir path")
)

func main() {
	flag.Parse()

	portStr := strconv.Itoa(*port)
	innerIP := getInnerIPAddr()
	path, err := filepath.Abs(*rootPath)
	if err != nil {
		fmt.Println("Not an available path.")
		return
	}

	fmt.Printf("Listening on port %s...\n", portStr)
	fmt.Printf("Root directory: \"%s\"\n", path)
	fmt.Printf("Please use the following url(s):\n")
	for _, ip := range innerIP {
		fmt.Printf("\thttp://%s:%s\n", ip, portStr)
	}
	fmt.Printf("Press [Ctrl+C] to exit\n")

	srv := http.Server{Addr: portStr}
	defer srv.Close()
	log.Fatal(http.ListenAndServe(":"+portStr, http.FileServer(http.Dir(path))))
}

func getInnerIPAddr() []string {
	var innerIP []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Something went wrong while getting ip address.")
		fmt.Println(err.Error())
	}
	for _, addr := range addrs {
		ipFormat := strings.Split(addr.String(), ".")
		if len(ipFormat) == 4 {
			n, err := strconv.Atoi(ipFormat[1])
			if err != nil {
				continue
			}
			if (ipFormat[0] == "192" && ipFormat[1] == "168") || //192.168.0.0 ~ 192.168.255.255
				(ipFormat[0] == "172" && n >= 16 && n <= 31) || //172.16.00 ~ 172.31.255.255
				(ipFormat[0] == "10") { //10.0.0.0 - 10.255.255.255

				innerIP = append(innerIP, strings.Split(strings.Join(ipFormat, "."), "/")[0])

			}
		}
	}
	return innerIP
}
