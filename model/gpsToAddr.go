package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/photoServer/global"
)

func ReverseGeocoding(lat, lon string) global.GPSJson {

	lat = GPSdegreeToDecimal(lat)
	lon = GPSdegreeToDecimal(lon)

	baseUrl := "https://nominatim.openstreetmap.org/reverse?format="
	format := "jsonv2"
	baseUrl = baseUrl + format + "&lat=" + lat + "&lon=" + lon

	result, _ := getJson2(baseUrl)
	return result
}

func getJson2(url string) (global.GPSJson, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	var gps global.GPSJson
	resp, err := client.Get(url)
	if err != nil {
		return gps, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return gps, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gps, fmt.Errorf("Read body: %v", err)
	}

	err = json.Unmarshal(data, &gps)
	if err != nil {
		panic(err)
	}

	return gps, nil
}

func GPSdegreeToDecimal(deg string) string {
	// lon := `121 deg 18' 32.00" E`
	// GPSdegreeToDecimal(lon)
	//fmt.Printf("Degree deg:%v\n", deg)

	re, _ := regexp.Compile(`(?P<deg>[0-9]+)\s+deg\s+(?P<min>[0-9]+)'\s+(?P<sec>[\.0-9]+)"\s+(?P<dir>[EWNS]{1})$`)
	match := re.FindStringSubmatch(deg)

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
			fmt.Println(result["name"])
		}
	}

	decDeg, _ := strconv.ParseFloat(result["deg"], 64)
	decMin, _ := strconv.ParseFloat(result["min"], 64)
	decSec, _ := strconv.ParseFloat(result["sec"], 64)
	decimal := decDeg + decMin/60.0 + decSec/3600.0
	if result["dir"] == "W" || result["dir"] == "S" {
		decimal = decimal * (-1)

	}
	return fmt.Sprintf("%.10f", decimal)
}
