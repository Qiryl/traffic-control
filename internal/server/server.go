package server

import (
    "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "github.com/Qiryl/traffic-control/config"
)

type Server struct {
	config *config.Config
    router *mux.Router
    logger *logrus.Logger
}


func NewServer(config *config.Config) *Server {
    s := &Server{
		config: config,
        router: mux.NewRouter(),
        logger: logrus.New(),
    }
    s.configureRouter()
    return s
}

func (s *Server) configureRouter() {
    s.router.StrictSlash(true)
    s.router.HandleFunc("/", tmp)
}

func tmp(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test")
}

func (s *Server) configureLogger() {
}

func (s *Server) Start() error {
    // s.configureRouter()
    return http.ListenAndServe(s.config.Server.Port, s.router)
}


