package main

import (
	"github.com/xieyuschen/Awesome-Web-Server-Building/server/sun"
)

func main() {
	server := sun.Default()
	server.Run(":8080")
}
