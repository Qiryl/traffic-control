package models

import (
    // "fmt"
    // "strconv"
    // "strings"
)

type Entry struct {
    Date     string  `csv:"date"`
    Number   string  `csv:"number"`
    Velocity float32 `csv:"velocity"`
}

// func (e Entry) String() string {
//     return fmt.Sprintf("%s, %s, %.2f", e.Datetime, e.VehicleNumber, e.Velocity)
// }
// 
// func (e *Entry) ToEntry(data string) error {
//     values := strings.Split(data, ",")
//     // fmt.Println("Data before: ", values)
//     v, err := strconv.ParseFloat(values[2], 32)
//     if err != nil {
//         return err
//     }
//     // fmt.Println("Data after: ", values)
//     e.Datetime = values[0]
//     e.VehicleNumber = values[1]
//     e.Velocity = float32(v)
//     // fmt.Println("Values: ", e.Datetime, e.VehicleNumber, e.Velocity)
//     return nil
// }
