package entry

import (
    "github.com/Qiryl/traffic-control/internal/entry"
    "github.com/Qiryl/traffic-control/internal/models"
)

type EntryUseCase struct {
    entryRepo entry.Repository
}

func NewEntryUseCase(entryRepo entry.Repository) *EntryUseCase {
    return &EntryUseCase{
        entryRepo: entryRepo,
    }
}

func (e EntryUseCase) CreateEntry(datetime, vehicleNumber string, velocity float32) error {
    entry := &models.Entry{
        Datetime: datetime,
        VehicleNumber: vehicleNumber,
        Velocity: velocity,
    }
    return e.entryRepo.CreateEntry(entry)
}

