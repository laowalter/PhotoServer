package photoServer

import (
	"time"
)

type Document struct { //mongodb Database: album, collection: pic
	FileName    string    `bson:"filename"`
	Path        string    `bson:"path"`
	CreateTime  time.Time `bson:"created_at"`
	ContentType string    `bson:"content_type,omitempty"`
	Thumbnail   []byte    `bson:"binary"`
	Md5         string    `bson:"md5"`
}

func VIDEO() []string {
	return []string{".mp4", ".mov", ".avi"}
}

func PIC() []string {
	return []string{".jpg", ".png", ".tiff", ".gif", ".jpeg", ".bmp"}
}

func RAW() []string {
	return []string{".nef", ".arw", ".cr2"}
}

func PICRAW() []string {
	return append(PIC(), RAW()...)
}

func PICVIDEO() []string {
	return append(append(PIC(), VIDEO()...), RAW()...)
}
