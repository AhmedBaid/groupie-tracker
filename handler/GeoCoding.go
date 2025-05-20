package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func GeoCoding(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	// Parse artist ID from query parameter
	Id := r.URL.Query().Get("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}
	var artistFound *tools.Artists
	errFetch := helpers.Fetch(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id), &artistFound)
	if errFetch != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Fetch geocoding data
	var locations tools.Locations
	if err := helpers.Fetch(artistFound.Locations, &locations); err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	geoLocations := make(map[string][2]float64)
	for _, location := range locations.Locations {
		coords, err := helpers.Geo(location, w)
		if err != nil {
			helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
			return
		}
		geoLocations[location] = coords
	}
	// Send geolocation data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(geoLocations)
}
