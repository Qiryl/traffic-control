package server

import (
    "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
    "net/http"
    "runtime"
    "reflect"

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

func NewServer(config *config.Config) (*Server, error) {

    entryRepo := entryStorage.NewEntryRepository(config.File.Path)

    s := &Server{
		config:  config,
        router:  mux.NewRouter(),
        logger:  logrus.New(),
        entryUC: entryUseCase.NewEntryUseCase(entryRepo),
    }

    s.configureRouter()
    if err := s.configureLogger(); err != nil {
        return nil, err
    }

    return s, nil
}

// TODO: Make error handling
func (s *Server) configureRouter() {
    h := entryHttp.NewHandler(s.entryUC)
    s.router.StrictSlash(true)
    s.router.HandleFunc("/create", log(h.CreateEntry))
    s.router.HandleFunc("/all", log(h.GetAll))
    s.router.HandleFunc("/", log(tmp))
    s.router.HandleFunc("/number/{number}", log(h.GetByCarNumber))
    s.router.HandleFunc("/date/{date}", log(h.GetByDate))
    s.router.HandleFunc("/velocity/{velocity}", log(h.GetByVelocity))
    s.router.HandleFunc("/limit/{date}/{velocity}", log(h.GetGreaterByDate))
    s.router.HandleFunc("/minmax/{date}", log(h.GetMinMaxByDate))
}

func tmp(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(201)
    w.Write([]byte("Check-Check"))
}

func log(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
        logrus.Info("Request received in: ", name)
        h(w, r)
    }
}

func (s *Server) configureLogger() error {
    level, err := logrus.ParseLevel(s.config.Logger.Level)
    if err != nil {
        return err
    }
    s.logger.SetLevel(level)
    return nil
}

func (s *Server) Start() error {
    s.logger.Info("Starting server")
    return http.ListenAndServe(s.config.Server.Port, s.router)
}
