package models

import (
    "fmt"
)

type Entry struct {
    Datetime      string // time.Time
    VehicleNumber string
    Velocity      float32
}

func (e Entry) String() string {
    return fmt.Sprintf("%s, %s, %.2f", e.Datetime, e.VehicleNumber, e.Velocity)
}

