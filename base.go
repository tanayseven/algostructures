package main

type DataElement interface {
	equalTo(other DataElement) bool
	greaterThan(other DataElement) bool
	lessThan(other DataElement) bool
	comparedTo(other DataElement) int
	toSting() string
}

type Integer struct {
	num int
}
