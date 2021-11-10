package controller

import (
	"encoding/json"
	"fmt"
	"my_task/A_Service/model"
	"my_task/A_Service/views"
	"net/http"
)

func connected() (ok bool) {
	_, err := http.Get("https://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}

func ping(d model.DAOInt) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			db := d.Ping() //ping B_Service
			net := connected() //check internet connection (google.com)
			data := views.Response{
				Code: http.StatusOK,
				Body: fmt.Sprintf("connected to internet: %t, connected to database: %t", net, db),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}

