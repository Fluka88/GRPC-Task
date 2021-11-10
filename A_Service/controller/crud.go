package controller

import (
	"encoding/json"
	"fmt"
	"my_task/A_Service/model"
	"my_task/A_Service/views"
	"net/http"
)

func crud(d model.DAOInt, c *model.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodPost:
			data := views.Movie{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := d.Create(data.Title, data.Year); err != nil{
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
			break
		case http.MethodGet:
			name := r.URL.Query().Get("name")
			data, err := d.Read(name, c)
			if err != nil{
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			break
		case http.MethodDelete:
			name := r.URL.Path[1:]
			if err := d.Delete(name); err != nil{
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct{
				Status string `json:status`
			}{"Item Deleted"})
			break
		default:

		}
	}
}