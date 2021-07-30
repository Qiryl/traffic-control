package entry

import (
    "github.com/Qiryl/traffic-control/internal/models"
)

type UseCase interface {
    CreateEntry(datetime, vehicleNumber string, velocity float32) error
    GetAll() ([]*models.Entry, error)
    GetByCarNumber(number string) ([]*models.Entry, error)
    GetByDate(date string) ([]*models.Entry, error)
    GetByVelocity(velocity float32) ([]*models.Entry, error)
    GetGreaterByDate(date string, velocity float32) ([]*models.Entry, error)
    GetMinMaxByDate(date string) ([]*models.Entry, error)
}

