package entry

import (
    "github.com/Qiryl/traffic-control/internal/models"
)

type Repository interface {
    CreateEntry(entry *models.Entry) error
    // GetEntries() error
}
