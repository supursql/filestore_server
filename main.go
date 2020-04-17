package main

import (
	"filestore_server/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", controller.UploadHandler)
	http.HandleFunc("/file/upload/suc", controller.UploadSucHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
