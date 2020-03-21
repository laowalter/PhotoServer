package main

import (
	"fmt"
	"os"
	"time"

	"github.com/barasher/go-exiftool"
)

func main() {
	file := os.Args[1]
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(file)

	var _cDate, _make, _model, _shutterSpeed, _iso, _aperture, _exposureCompensation, _exposureTime, _lensSpec, _lensID, _gpsPosition string
	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			//fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}
		for k, v := range fileInfo.Fields {
			fmt.Printf("key: %v, value: %v, valueType: %T\n", k, v, v)
			switch k {
			case "CreateDate":
				_cDate = fmt.Sprintf("%v", v)
			case "Make":
				_make = fmt.Sprintf("%v", v)
			case "Model":
				_model = fmt.Sprintf("%v", v)
			case "LensSpec":
				_lensSpec = fmt.Sprintf("%v", v)
			case "LensID":
				_lensID = fmt.Sprintf("%v", v)
			case "ShutterSpeed":
				_shutterSpeed = fmt.Sprintf("%v", v)
			case "ExposureTime":
				_exposureTime = fmt.Sprintf("%v", v)
			case "ISO":
				_iso = fmt.Sprintf("%v", v)
			case "Aperture":
				_aperture = "f/" + fmt.Sprintf("%v", v)
			case "ExposureCompensation":
				_exposureCompensation = fmt.Sprintf("%v", v)
			case "GPSPosition":
				_gpsPosition = fmt.Sprintf("%v", v)
			default:
			}
		}

		exifCreateDate, err := time.Parse("2006:01:02 15:04:05", _cDate)
		if err != nil {
			fmt.Printf("Opps! Cannot convert %s of %v\n", exifCreateDate, file)
		} else {
			fmt.Println(exifCreateDate)
		}

		fmt.Println(_make)
		fmt.Println(_model)
		fmt.Println(_shutterSpeed)
		fmt.Println(_iso)
		fmt.Println(_aperture)
		fmt.Println(_exposureTime)
		fmt.Println(_exposureCompensation)
		fmt.Println(_lensSpec)
		fmt.Println(_lensID)
		fmt.Println(_gpsPosition)

	}
}
