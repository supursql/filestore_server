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

		fmt.Println("1", db.HasTable(&meta))
		fmt.Println("2", db.HasTable(&meta.FileMeta{}))
		fmt.Println("3", db.HasTable("tbl_file"))

		db.Debug().Create(meta)
		return true
	} else {
		fmt.Printf("File with hash:%s has been uploaded before", meta.FileSha1)
		return true
	}

	return false
}
