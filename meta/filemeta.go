package meta

import (
	"github.com/jinzhu/gorm"
	"sort"
)

/***
文件元信息结构
*/

type FileMeta struct {
	gorm.Model
	FileSha1 string
	FileName string
	FileSize int64
	FileAddr string
	Status   int
	Ext1     int
	Ext2     string
}

func (FileMeta) TableName() string {
	return "tbl_file"
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

/**
通过sha1值获取文件元信息
*/
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
