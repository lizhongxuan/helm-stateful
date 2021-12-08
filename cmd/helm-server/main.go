package main

import (
	log "github.com/golang/glog"
	"helm-statuful/api"
)

func main() {
	api.StartApiServer(":8080")
	log.Error("main end...")
}


