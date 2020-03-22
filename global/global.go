package global

import (
	"time"
)

const (
	ThumbnailHeight = 250
)

type Document struct { //mongodb Database: album, collection: pic
	FileName    string `bson:"filename"`
	Path        string `bson:"path"`
	ContentType string `bson:"content_type"`
	Thumbnail   []byte `bson:"thumbnail"`
	Thumbnail64 string `bson:"-, omitempty"`
	Md5         string `bson:"md5"`
	GPSPosition `bson:"inline, omitempty"`
	Exif        `bson:"inline, omitempty"`
	Tags        []string `bson:"tags, omitempty"`
}

type Exif struct {
	CreadTime            time.Time
	Make                 string
	Model                string
	ShutterSpeed         string
	ISO                  string
	Aperture             string
	ExposureCompensation string
	ExposureTime         string
	LensSpec             string
	LensID               string
}

type GPSPosition struct {
	Latitude  string
	Longitude string
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
