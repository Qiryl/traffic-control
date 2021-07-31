package http

import (
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
	"github.com/gocarina/gocsv"

    "github.com/Qiryl/traffic-control/internal/entry"
)

type Handler struct {
    useCase entry.UseCase
}

func NewHandler(useCase entry.UseCase) *Handler {
    return &Handler{
        useCase: useCase,
    }
}

func (h *Handler) CreateEntry(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    h.useCase.CreateEntry(vars["date"], vars["number"], vars["velocity"])
}

// TODO: Add error respponse
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
    entries, err := h.useCase.GetAll()
    if err != nil {
       //
    }

	entiesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }

    w.Header().Set("Content-Type", "text/csv")
    w.Write(entiesCsv)
}

func (h *Handler) GetByCarNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    entries, err := h.useCase.GetByCarNumber(vars["number"])
    if err != nil {
       //
    }

    entriesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }

    w.Header().Set("Content-Type","text/csv")
    w.Write(entriesCsv)
}

func (h *Handler) GetByDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    entries, err := h.useCase.GetByDate(vars["date"])
    if err != nil {
       //
    }

    entriesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }
    w.Header().Set("Content-Type","text/csv")
    w.Write(entriesCsv)
}

func (h *Handler) GetByVelocity(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    v, err := strconv.ParseFloat(vars["velocity"], 32)
    velocity := float32(v)
    if err != nil {
       //
    }

    entries, err := h.useCase.GetByVelocity(velocity)
    if err != nil {
       //
    }

    entriesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }
    w.Header().Set("Content-Type","text/csv")
    w.Write(entriesCsv)
}

func (h *Handler) GetGreaterByDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    v, err := strconv.ParseFloat(vars["velocity"], 32)
    velocity := float32(v)
    if err != nil {
       //
    }

    entries, err := h.useCase.GetGreaterByDate(vars["date"], velocity)
    if err != nil {
       //
    }

    entriesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }
    w.Header().Set("Content-Type","text/csv")
    w.Write(entriesCsv)
}

func (h *Handler) GetMinMaxByDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    entries, err := h.useCase.GetMinMaxByDate(vars["date"])
    if err != nil {
       //
    }

    entriesCsv, err := gocsv.MarshalBytes(entries)
    if err != nil {
       //
    }
    w.Header().Set("Content-Type","text/csv")
    w.Write(entriesCsv)
}
