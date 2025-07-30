package model

type EntryType string

const (
	INCOME EntryType = "INCOME"
	EXPENSE EntryType = "EXPENSE"
)

type Entry struct {
	Amount float64
	Category string
	Type EntryType
}