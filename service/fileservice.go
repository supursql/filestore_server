package service

import (
	"filestore_server/db"
	"filestore_server/meta"
)

func UploadFileMetaDB(meta meta.FileMeta) bool {
	return db.OnFileUploadFinished(meta)
}
