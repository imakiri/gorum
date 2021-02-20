package web

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/gate"
	"net/http"
)

type Service struct {
	gate        gate.GeneralService
	config      cfg.Web
	Server      *http.Server
	RedirServer *http.Server
}

func NewService(c cfg.EI, g gate.GeneralService) (*Service, error) {
	var s Service
	var err error

	s.gate = g
	s.config = c.Web

	router := mux.NewRouter()
	redirRouter := mux.NewRouter()

	redirRouter.HandleFunc("/", redirect)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", s.Root)

	s.Server = &http.Server{}
	s.RedirServer = &http.Server{}
	s.Server.Handler = router
	s.RedirServer.Handler = redirRouter

	return &s, err
}

// Redirect from HTTP to HTTPS
func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

// Internal ServiceInt Error Response
func ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
