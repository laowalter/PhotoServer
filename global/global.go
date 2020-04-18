package global

import (
	"time"
)

const (
	//Fixed Thubnail height
	ThumbnailHeight = 250

	//Photos Per Page
	PhotosPerPage = int64(100)

	//Deleted file Directory
	RemovedDir = "/data/album/delete"

	//Mongodatabase
	MongoUri = "mongodb://localhost:27017"
	DBname   = "album"
	PICcol   = "pic"
)

type Document struct {
	//保存到数据库中的基础结构
	FileName    string    `bson:"filename"`
	Path        string    `bson:"path"`
	PathBase64  string    `bson:,omitempty` //base64
	ContentType string    `bson:"content_type"`
	Thumbnail   string    `bson:"thumbnail"` //base64
	Md5         string    `bson:"md5"`
	ImportTime  time.Time `bson:"import_at"`
	Tags        []string  `bson:tag,omitempty`
	GPSPosition `bson:"inline,omitempty"`
	Exif        `bson:"inline,omitempty"`
}

type Exif struct {
	CreateDate           time.Time
	Make                 string
	Model                string
	ShutterSpeed         string
	ISO                  string
	Aperture             string
	ExposureCompensation string
	ExposureTime         string
	LensSpec             string
	LensID               string
	FocalLength          string
	FullImageSize        string
}

type GPSPosition struct {
	Latitude  string
	Longitude string
}

type Pagers struct {
	Before  []int64
	Current int64
	After   []int64
}

type YearCount struct {
	Year   int32 //return from mongodb was int32 origial
	Number int32
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
