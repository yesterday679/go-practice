package main

import (
	"flag"
	"fmt"
	"learn/crawler_distributed/rpcsupport"
	"learn/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 9000, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
