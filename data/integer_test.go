package data

import (
	"testing"
)

func TestEqualToGivesTrueResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 3}
	expectedResult := true

	// when
	actualResult := num1.EqualTo(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The two values should be equal")
	}
}
func TestEqualToGivesFalseResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 4}
	expectedResult := false

	// when
	actualResult := num1.EqualTo(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The two values should NOT be equal")
	}
}
func TestLessThanTrueResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 4}
	expectedResult := true

	// when
	actualResult := num1.LessThan(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should be less than second")
	}
}

func TestLessThanFalseResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 4}
	var num2 DataElement = &Integer{val: 3}
	expectedResult := false

	// when
	actualResult := num1.LessThan(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should NOT be less than second")
	}
}

func TestGreaterThanTrueResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 4}
	var num2 DataElement = &Integer{val: 3}
	expectedResult := true

	// when
	actualResult := num1.GreaterThan(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should be greater than second")
	}
}

func TestGreaterThanFalseResult(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 4}
	expectedResult := false

	// when
	actualResult := num1.GreaterThan(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should NOT be greater than second")
	}
}

func TestComparedToWhenLessThan(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 4}
	expectedResult := -1

	// when
	actualResult := num1.ComparedTo(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should NOT be less than second")
	}
}

func TestComparedToWhenEqualTo(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 3}
	var num2 DataElement = &Integer{val: 3}
	expectedResult := 0

	// when
	actualResult := num1.ComparedTo(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should NOT be less than second")
	}
}

func TestComparedToWhenGreaterThan(t *testing.T) {
	// given
	var num1 DataElement = &Integer{val: 4}
	var num2 DataElement = &Integer{val: 3}
	expectedResult := 1

	// when
	actualResult := num1.ComparedTo(num2)

	// then
	if actualResult != expectedResult {
		t.Errorf("The first number should NOT be less than second")
	}
}

func TestToString(t *testing.T) {
	// given
	var num DataElement = &Integer{val: 4}
	expectedResult := "4"

	// when
	actualResult := num.ToString()

	// then
	if actualResult != expectedResult {
		t.Errorf("String conversion should be successful")
	}
}
