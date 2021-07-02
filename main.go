package main

import (
	"go.spider/handle"
	"go.spider/tools"
	"log"
)

func main() {
	url :=  "https://imoemei.com/"
	var req = new(tools.Request)
	var index = new(handle.IndexHandle)
	request, err := req.NewRequest("GET", url, "", index, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := request.Execute(); err != nil{
		log.Fatal(err)
	}
}
