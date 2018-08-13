package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	port     = flag.Int("port", 1113, "set the port")
	rootPath = flag.String("root", "./", "set the root dir path")
)

func main() {
	flag.Parse()
	portStr := strconv.Itoa(*port)
	fmt.Printf("Listening on port %s\n", portStr)
	fmt.Printf("Root directory: %s\n", *rootPath)
	fmt.Printf("Please open http://localhost:%s\n", portStr)
	fmt.Printf("Press [Ctrl+C] to exit\n")
	log.Fatal(http.ListenAndServe(":"+portStr, http.FileServer(http.Dir(*rootPath))))
}
