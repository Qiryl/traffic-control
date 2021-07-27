package entry

import (
    "os"
    "fmt"
    "log"
    "github.com/Qiryl/traffic-control/internal/models"
)

type EntryRepository struct {
    filePath string
}

func NewEntryRepository(filePath string) *EntryRepository {
    return &EntryRepository{
        filePath: filePath,
    }
}

func (r EntryRepository) CreateEntry(entry *models.Entry) error {
    f, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Open File: %v", err)
        return err
    }

    if _, err := fmt.Fprintln(f, entry.String()); err != nil {
        log.Fatalf("Write to File: %v", err)
        f.Close()
        return err
    }

    return f.Close()
}

func GetEntries() error {

    return nil
}
