package main

import (
	"filestore_server/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", controller.UploadHandler)
	http.HandleFunc("/file/upload/suc", controller.UploadSucHandler)
	http.HandleFunc("/file/query", controller.FileQueryHandler)
	http.HandleFunc("/file/meta", controller.GetFileMetaHandler)
	http.HandleFunc("/file/download", controller.DownloadHandler)
	http.HandleFunc("/file/update", controller.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", controller.FileDeleteHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
