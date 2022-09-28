package main

import (
	"github.com/atedesch1/csi-flix/db"
	"github.com/atedesch1/csi-flix/server"
)

func main() {
	db.Init()
	server.Init()
}
