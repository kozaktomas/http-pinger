package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var url string
	var delaySeconds int
	flag.StringVar(&url, "url", "http://localhost:8090/hello", "url to call")
	flag.IntVar(&delaySeconds, "delay", 2, "timeout between calls")
	flag.Parse()

	fmt.Printf("Going to call %q every %d seconds\n", url, delaySeconds)
	ticker := time.NewTicker(time.Duration(delaySeconds) * time.Second)
	for _ = range ticker.C {
		doCall(url)
	}
}

func doCall(url string) {
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Default().Println(fmt.Sprintf("error calling url %q: %s", url, err))
		return
	}

	if r.StatusCode != http.StatusOK {
		log.Default().Println(fmt.Sprintf("error calling url %q: %s", url, r.Status))
		return
	}

	log.Default().Println(fmt.Sprintf("call ok %q", url))
}
