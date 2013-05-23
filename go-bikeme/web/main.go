package web

import (
	"fmt"
	"net/http"
	"go-bikeme/bikeshareservice"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/update_stations", update_stations)
}

func update_stations(w http.ResponseWriter, r *http.Request) {
	services := []bikeshareservice.Service{bikeshareservice.NewBicingService(), bikeshareservice.NewCapitalBikeShareService(), bikeshareservice.NewTelOFunService()}
	for _, service := range services {
		stations, err := service.Stations()
		if err != nil {
			fmt.Fprintf(w, "#main() received an error: '%s'\n", err.Error())
			return
		}
		fmt.Fprintf(w,"There are %d stations in the %T system!\n", len(stations), service)
	}
	fmt.Fprint(w, "done")
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
