package server

import (
    "net/http"
    "runtime"
    "reflect"
    "time"

    "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"

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

func (s *Server) configureRouter() {
    h := entryHttp.NewHandler(s.entryUC)
    s.router.StrictSlash(true)
    s.router.HandleFunc("/create/{date}/{number}/{velocity}", log(h.CreateEntry))
    s.router.HandleFunc("/all", log(s.access(h.GetAll)))
    s.router.HandleFunc("/number/{number}", log(s.access(h.GetByCarNumber)))
    s.router.HandleFunc("/date/{date}", log(s.access(h.GetByDate)))
    s.router.HandleFunc("/velocity/{velocity}", log(s.access(h.GetByVelocity)))
    s.router.HandleFunc("/limit/{date}/{velocity}", log(s.access(h.GetGreaterByDate)))
    s.router.HandleFunc("/minmax/{date}", log(s.access(h.GetMinMaxByDate)))
}

func log(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
        logrus.Info("Request received in: ", name)
        h(w, r)
    }
}

func (s *Server) access(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        now := time.Now().Format("15:04")
        if now > s.config.Time.Start && now < s.config.Time.End {
            h(w, r)
        } else {
            http.Error(w, http.StatusText(403), 403)
        }
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
