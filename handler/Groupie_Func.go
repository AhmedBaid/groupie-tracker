package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	tools "groupie/tools"
)

func GroupieFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errore := tools.ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		// execute the not found  template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data: ", err)
		return
	}
	defer res.Body.Close()

	var post []tools.Artists
	Fetch_locations("https://groupietrackers.herokuapp.com/api/locations")
	Fetch_dates("https://groupietrackers.herokuapp.com/api/dates")
	Fetch_relation("https://groupietrackers.herokuapp.com/api/relation")
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		return
	}

	tools.Tp.ExecuteTemplate(w, "index.html", post)
}
