package main

import (
	"web/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/username/", controller.UserHomeHandler)

	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.HandleFunc("/", controller.HomeHandler)

	http.HandleFunc("/index.html", controller.Index)

	//通过api的方式来连接其他服务
	http.HandleFunc("/api", controller.ApiHandler)

	//通过proxy的方式来避免跨域
	http.HandleFunc("/upload/:vid-id", controller.ProxyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
