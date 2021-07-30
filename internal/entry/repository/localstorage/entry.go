package localstorage

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
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
        return err
    }

    if _, err := fmt.Fprintln(f, entry.String()); err != nil {
        f.Close()
        return err
    }

    return f.Close()
}

func (r EntryRepository) GetAll() ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)

	for scanner.Scan() {
        entry := new(models.Entry)
        if err := entry.ToEntry(scanner.Text()); err != nil {
            return nil, err
        }
        entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}

func (r EntryRepository) GetByCarNumber(number string) ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)

	for scanner.Scan() {
        if strings.Contains(strings.Split(scanner.Text(), ",")[1], number) {
            entry := new(models.Entry)
            if err := entry.ToEntry(scanner.Text()); err != nil {
                return nil, err
            }
            entries = append(entries, entry)
        }
	}

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}

func (r EntryRepository) GetByDate(date string) ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)

	for scanner.Scan() {
        if strings.Contains(strings.Split(scanner.Text(), ",")[0], date) {
            entry := new(models.Entry)
            if err := entry.ToEntry(scanner.Text()); err != nil {
                return nil, err
            }
            entries = append(entries, entry)
        }
	}

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}

func (r EntryRepository) GetByVelocity(velocity float32) ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)
    var values []string
    var v1 float64

	for scanner.Scan() {
        values = strings.Split(scanner.Text(), ",")
        v1, err = strconv.ParseFloat(values[2], 32)
        v2 := float32(v1)
        if err != nil {
            return nil, err
        }
        if velocity == v2 {
            entry := new(models.Entry)
            if err := entry.ToEntry(scanner.Text()); err != nil {
                return nil, err
            }
            entries = append(entries, entry)
        }
	}

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}

func (r EntryRepository) GetGreaterByDate(date string, velocity float32) ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)
    var values []string
    var v1 float64

	for scanner.Scan() {
        values = strings.Split(scanner.Text(), ",")
        v1, err = strconv.ParseFloat(values[2], 32)
        v2 := float32(v1)
        if err != nil {
            return nil, err
        }
        if strings.Contains(values[0], date) && velocity < v2 {
            entry := new(models.Entry)
            if err := entry.ToEntry(scanner.Text()); err != nil {
                return nil, err
            }
            entries = append(entries, entry)
        }
	}

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}

func (r EntryRepository) GetMinMaxByDate(date string) ([]*models.Entry, error) {
    f, err := os.OpenFile(r.filepath, os.O_RDONLY, 0644)
    if err != nil {
		return nil, err
    }
	scanner := bufio.NewScanner(f)
    entries := make([]*models.Entry, 0)
    var max, min, tmp float64
    var maxValues, minValues, values []string

    if scanner.Scan() {
        values = strings.Split(scanner.Text(), ",")
        maxValues = values
        minValues = values
        min, _ = strconv.ParseFloat(values[2], 32)
    }

	for scanner.Scan() {
        values = strings.Split(scanner.Text(), ",")
        tmp, err = strconv.ParseFloat(values[2], 32)
        if err != nil {
            return nil, err
        }
        if strings.Contains(values[0], date) {
            if max < tmp {
                max = tmp
                maxValues = values
            }
            if min > tmp {
                min = tmp
                minValues = values
            }
        }
	}
	maxEntry := new(models.Entry)
	minEntry := new(models.Entry)
    if err := maxEntry.ToEntry(strings.Join(maxValues, ",")); err != nil {
        return nil, err
    }
    if err := minEntry.ToEntry(strings.Join(minValues, ",")); err != nil {
        return nil, err
    }
    entries = append(entries, maxEntry, minEntry)

	if err := scanner.Err(); err != nil {
		f.Close()
		return nil, err
	}
	return entries, f.Close()
}
