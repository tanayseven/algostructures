package data

type DataElement interface {
	EqualTo(other DataElement) bool
	GreaterThan(other DataElement) bool
	LessThan(other DataElement) bool
	ComparedTo(other DataElement) int
	ToString() string
}
