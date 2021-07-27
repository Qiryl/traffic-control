package entry

type UseCase interface {
    CreateEntry(time, vehicleNumber string, velocity float32) error
    GetEntries() error
}

