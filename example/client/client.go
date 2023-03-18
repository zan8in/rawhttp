package main

import (
	"log"
	"net/http"

	"github.com/remeh/sizedwaitgroup"
	"github.com/zan8in/rawhttp"
)

func main() {
	swg := sizedwaitgroup.New(25)
	pipeOptions := rawhttp.DefaultPipelineOptions
	pipeOptions.Host = "127.0.0.1:10000"
	pipeOptions.MaxConnections = 1
	pipeclient := rawhttp.NewPipelineClient(pipeOptions)
	for i := 0; i < 50; i++ {
		swg.Add()
		go func(swg *sizedwaitgroup.SizedWaitGroup) {
			defer swg.Done()
			req, err := http.NewRequest("GET", "http://127.0.0.1:10000/headers", nil)
			if err != nil {
				log.Printf("Error sending request to API endpoint. %+v", err)
				return
			}
			req.Host = "127.0.0.1:10000"
			req.Header.Set("Host", "127.0.0.1:10000")
			resp, err := pipeclient.Do(req)
			if err != nil {
				log.Printf("Error sending request to API endpoint. %+v", err)
				return
			}
			log.Printf("%+v\n", resp)
			_ = resp
		}(&swg)
	}

	swg.Wait()

}
