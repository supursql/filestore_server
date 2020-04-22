package db

import (
	myDB "filestore_server/db/mysql"
	"filestore_server/meta"
	"fmt"
)

//文件上传持久化
func OnFileUploadFinished(meta meta.FileMeta) bool {

	db := myDB.DBConn()

	meta.Status = 1

	if db.NewRecord(meta) {
		db.Create(meta)
		return true
	} else {
		fmt.Printf("File with hash:%s has been uploaded before", meta.FileSha1)
		return true
	}

	return false
}
