package main

import (
	"net/http"
	"log"
	"weixinminiprogram/controller"
)


func main() {
	http.HandleFunc("/d1",controller.Cal)
	http.HandleFunc("/thisday",controller.Thisday)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal(err)
	}
}
