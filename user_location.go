package user_location

import (
	"encoding/json"
	"net/http"
)

type Location struct {
	Country string
	Region  string
	City    string
	LatLong []string
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	location := Location{r.Header.Get("X-AppEngine-Country"),
		r.Header.Get("X-AppEngine-Region"),
		r.Header.Get("X-AppEngine-City"),
		[]string{r.Header.Get("X-AppEngine-CityLatLong")}}
	js, err := json.Marshal(location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
