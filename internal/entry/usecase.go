package entry

type UseCase interface {
    CreateEntry(datetime, vehicleNumber string, velocity float32) error
    // GetEntries() error
}

