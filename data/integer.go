package data

import "strconv"

type Integer struct {
	val int
}

func (i *Integer) EqualTo(other DataElement) bool {
	return i.val == other.(*Integer).val
}

func (i *Integer) GreaterThan(other DataElement) bool {
	return i.val > other.(*Integer).val
}

func (i *Integer) LessThan(other DataElement) bool {
	return i.val < other.(*Integer).val
}

func (i *Integer) ComparedTo(other DataElement) int {
	num1 := i.val
	num2 := other.(*Integer).val
	if num1 == num2 {
		return 0
	}
	if num1 > num2 {
		return 1
	}
	return -1
}

func (i *Integer) ToString() string {
	return strconv.Itoa(i.val)
}
