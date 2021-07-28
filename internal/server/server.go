package server

import (
    "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "github.com/Qiryl/traffic-control/config"
    "github.com/Qiryl/traffic-control/internal/entry"
    entryStorage "github.com/Qiryl/traffic-control/internal/entry/repository/localstorage"
    entryHttp "github.com/Qiryl/traffic-control/internal/entry/delivery/http"
    entryUseCase "github.com/Qiryl/traffic-control/internal/entry/usecase"
)

type Server struct {
	config  *config.Config
    router  *mux.Router
    logger  *logrus.Logger
    entryUC entry.UseCase
}


func NewServer(config *config.Config) *Server {

    entryRepo := entryStorage.NewEntryRepository(config.File.Path)

    s := &Server{
		config:  config,
        router:  mux.NewRouter(),
        logger:  logrus.New(),
        entryUC: entryUseCase.NewEntryUseCase(entryRepo),
    }

    s.configureRouter()
    return s
}

func (s *Server) configureRouter() {
    h := entryHttp.NewHandler(s.entryUC)
    s.router.StrictSlash(true)
    s.router.HandleFunc("/", tmp)
    s.router.HandleFunc("/create", h.CreateEntry)
}

func tmp(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test")
}

// func (s *Server) configureLogger() {
// }

func (s *Server) Start() error {
    return http.ListenAndServe(s.config.Server.Port, s.router)
}


