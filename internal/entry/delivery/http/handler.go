package http

import (
    "net/http"
    // "fmt"
    "github.com/Qiryl/traffic-control/internal/entry"
    // "github.com/gorilla/mux"
)

type Entry struct {
    Datetime      string  `json:"datetime"`
    VehicleNumber string  `json:"number"`
    Velocity      float32 `json:"velocity"`
}

type Handler struct {
    useCase entry.UseCase
}

func NewHandler(useCase entry.UseCase) *Handler {
    return &Handler{
        useCase: useCase,
    }
}

func (h *Handler) CreateEntry(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Create Entry")
    h.useCase.CreateEntry("datetime", "number", 60.0)
}

// func ConfigureEntryEndpoints(router *mux.Router, eu entry.UseCase) {
//     handler := NewHandler(eu)
//     
// }

