package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie/tools"
)

func Geo(location string, w http.ResponseWriter) ([2]float64, error) {
	var result [2]float64
	loca := ""
	for _, char := range location {
		if char == '-' {
			break
		}
		loca += string(char)
	}
	var data tools.Geodata
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", location)
	err := Fetch(url, &data)
	if err != nil {
		return result, err
	}

	// If data is available, extract the coordinates
	if len(data) > 0 {
		lat, errLat := strconv.ParseFloat(data[0].Lat, 64)
		lon, errLon := strconv.ParseFloat(data[0].Lon, 64)
		if errLat != nil || errLon != nil {
			return result, err
		}
		result = [2]float64{lat, lon}
	}
	fmt.Println("Coordinates: ", result)

	return result, nil
}
