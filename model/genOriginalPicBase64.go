package model

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"runtime"

	"github.com/disintegration/imaging"
)

func GenOriginalPicBase64(file string) string { //file: wholepath contains filename.
	runtime.GOMAXPROCS(runtime.NumCPU())
	img, err := imaging.Open(file, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
