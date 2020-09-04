package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api/rest"
	"io"
	"log"
	"net/http"
)

var Router = mux.NewRouter().Schemes("http").Subrouter()

var server = &http.Server{}

func root(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	_, _ = io.WriteString(w, fmt.Sprintf("User: %s, Location: %s, Visit: %s", q["user"], q["location"], q["visit"]))
}

func init() {
	_ = rest.RunREST(Router)
	//_ = apiApp.RunGraphQL(Router)

	Router.HandleFunc("/", root)
	server.Handler = Router

	log.Fatal(server.ListenAndServe())
}

func Run() {}
