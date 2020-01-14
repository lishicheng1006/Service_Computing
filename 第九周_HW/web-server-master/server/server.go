package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	// New mux and server
	r := mux.NewRouter()
	server := negroni.Classic()

	// Add router
	initRouter(r, formatter)

	server.UseHandler(r)
	return server
}

func initRouter(r *mux.Router, formatter *render.Render) {
	r.HandleFunc("/server/{info}", helloHandler(formatter)).Methods("GET")
}

func helloHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		info := vars["info"]
		formatter.JSON(w, http.StatusOK, map[string]string{"info": info})
	}
}
