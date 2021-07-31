package localstorage

import (
    "os"
    "strings"
    "io/ioutil"

	"github.com/gocarina/gocsv"

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
    if _, err := os.Stat(r.filepath); os.IsNotExist(err) {
        err := ioutil.WriteFile(r.filepath, []byte("date,number,velocity\n"), 0755)
        if err != nil {
            return err
        }
    }

    file, err := os.OpenFile(r.filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }

    if err := gocsv.MarshalWithoutHeaders([]*models.Entry{entry}, file); err != nil {
        file.Close()
        return err
    }

    return file.Close()
}

func (r EntryRepository) GetAll() ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	return entries, file.Close()
}

func (r EntryRepository) GetByCarNumber(number string) ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
    result := make([]*models.Entry, 0)
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	for _, entry := range entries {
		if entry.Number == number {
			result = append(result, entry)
		}
	}

	return result, file.Close()
}

func (r EntryRepository) GetByDate(date string) ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
    result := make([]*models.Entry, 0)
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	for _, entry := range entries {
		if strings.Contains(entry.Date, date) {
			result = append(result, entry)
		}
	}

	return result, file.Close()
}

func (r EntryRepository) GetByVelocity(velocity float32) ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
    result := make([]*models.Entry, 0)
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	for _, entry := range entries {
		if entry.Velocity == velocity {
			result = append(result, entry)
		}
	}

	return result, file.Close()
}

func (r EntryRepository) GetGreaterByDate(date string, velocity float32) ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
    result := make([]*models.Entry, 0)
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	for _, entry := range entries {
		if strings.Contains(entry.Date, date) && velocity < entry.Velocity {
			result = append(result, entry)
		}
	}
	return result, file.Close()
}

func (r EntryRepository) GetMinMaxByDate(date string) ([]*models.Entry, error) {
    file, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }

    entries := make([]*models.Entry, 0)
    result := make([]*models.Entry, 0)
	var max, min float32
	var maxEntry, minEntry *models.Entry

	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		file.Close()
		return nil, err
	}

	for _, entry := range entries {
		if strings.Contains(entry.Date, date) {
			if max < entry.Velocity {
				max = entry.Velocity
				maxEntry = entry
			}
			if min > entry.Velocity || min == 0 {
				min = entry.Velocity
				minEntry = entry
			}
		}
	}

	return append(result, minEntry, maxEntry), file.Close()
}
