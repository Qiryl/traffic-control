package usecase

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
        Date: datetime,
        Number: vehicleNumber,
        Velocity: velocity,
    }
    return e.entryRepo.CreateEntry(entry)
}

func (e EntryUseCase) GetAll() ([]*models.Entry, error) {
    return e.entryRepo.GetAll()
}

func (e EntryUseCase) GetByCarNumber(number string) ([]*models.Entry, error) {
    return e.entryRepo.GetByCarNumber(number)
}

func (e EntryUseCase) GetByDate(date string) ([]*models.Entry, error) {
    return e.entryRepo.GetByDate(date)
}

func (e EntryUseCase) GetByVelocity(velocity float32) ([]*models.Entry, error) {
    return e.entryRepo.GetByVelocity(velocity)
}

func (e EntryUseCase) GetGreaterByDate(date string, velocity float32) ([]*models.Entry, error) {
    return e.entryRepo.GetGreaterByDate(date, velocity)
}

func (e EntryUseCase) GetMinMaxByDate(date string) ([]*models.Entry, error) {
    return e.entryRepo.GetMinMaxByDate(date)
}
