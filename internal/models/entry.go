package models

type Entry struct {
    Date     string  `csv:"date"`
    Number   string  `csv:"number"`
    Velocity float32 `csv:"velocity"`
}
