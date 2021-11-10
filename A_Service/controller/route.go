package controller

import (
	"my_task/A_Service/model"
	"net/http"
)

func Register(d model.DAOInt, c *model.Cache) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", ping(d))
	mux.HandleFunc("/", crud(d, c))
	return mux
}

