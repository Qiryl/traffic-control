package localstorage

import (
    "os"
    "fmt"
    "log"
    "github.com/Qiryl/traffic-control/internal/models"
)

type EntryRepository struct {
    filepath string
}

func NewEntryRepository(filepath string) *EntryRepository {
    return &EntryRepository{
        filepath: filepath,
    }
}

func (r EntryRepository) CreateEntry(entry *models.Entry) error {
    f, err := os.OpenFile(r.filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

func (r EntryRepository) GetEntries(entry *models.Entry) error {
    
    return nil
}
