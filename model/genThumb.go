package model

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"runtime"

	"github.com/disintegration/imaging"
	"github.com/photoServer/global"
)

func GenThumbBase64(file string, size string) string { //file: wholepath contains filename.
	runtime.GOMAXPROCS(runtime.NumCPU())
	img, err := imaging.Open(file, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}

	rectangle := img.Bounds()
	t_width := int(rectangle.Dx() * global.ThumbnailHeight / rectangle.Dy())
	thumb := imaging.Thumbnail(img, t_width, global.ThumbnailHeight, imaging.CatmullRom)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
