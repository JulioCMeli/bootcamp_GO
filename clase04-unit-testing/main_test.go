// go test main_test.go

package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	main "julio.com/clase03-funciones"
)

func TestRetrieveTax_OK(t *testing.T) {
	//arrange
	value := 10000.0
	expected := float64(value) * 0.17

	// act
	result := main.RetrieveTax(value)

	//asserrt

	//Nativo
	if expected != result {
		t.Errorf("NO ok")
	}
	//testify
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestRetrieveTax_NO_OK(t *testing.T) {
	//arrange
	value := 10000.0
	expected := float64(value) * 0.50

	// act
	result := main.RetrieveTax(value)

	//asserrt
	if expected == result {
		t.Errorf("NO ok")
	}
}

func TestAverageFunc_OK(t *testing.T) {
	//arrange
	values := []float64{2, 3, 5, 7, 11, 13}
	expected := 6.833333333333333

	// act
	result := main.AverageFunc(values)

	//Nativo
	if expected != result {
		t.Errorf("NO ok")
	}
	//testify
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")

}

func TestMarinaSalary_OK(t *testing.T) {
	//arrange
	minutes := 10
	cat := "A"
	expected := float64(minutes) * 3000 * 1.5

	// act
	result := main.MarinaSalary(minutes, cat)

	//Nativo
	if expected != result {
		t.Errorf("NO ok")
	}
	//testify
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")

}
