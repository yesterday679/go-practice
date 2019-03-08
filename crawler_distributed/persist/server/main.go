package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"learn/crawler_distributed/config"
	"learn/crawler_distributed/persist"
	"learn/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	err := serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex)
	log.Fatal(err)
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
