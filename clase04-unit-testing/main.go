// go mod init julio.com/clase03-funciones
// go run main.go
// go build main.go && ./main
// Funciones en Go
// Práctica clase 2 - Go Bases

package main

import "fmt"

func main() {

	var salary float64 = 40000
	tax := RetrieveTax(salary)
	fmt.Println("The tax of", salary, "is", tax)

	values := []float64{2, 3, 5, 7, 11, 13}
	avg := AverageFunc(values)
	fmt.Printf("Average of %v is %v \n", values, avg)

	var minutes int = 20
	var category string = "B"
	mSalary := MarinaSalary(minutes, category)
	fmt.Printf("Marina Salary for %v minutes ang class %v  is %v \n", minutes, category, mSalary)

	minFunc, err := operation(minimum)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	AverageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	averageValue := AverageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Min %v , err: %v \n", minValue, err)
	fmt.Printf("Avg %v , err: %v \n", averageValue, err)
	fmt.Printf("Max %v , err: %v \n", maxValue, err)

}

func RetrieveTax(salary float64) float64 {
	percentage := 0.17
	if salary > 50000 {
		percentage = 0.27
	}
	return salary * percentage
}

func AverageFunc(values []float64) float64 {
	var sum float64 = 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func MarinaSalary(minutes int, category string) float64 {

	var salary float64 = 0.0
	switch category {
	case "C":
		salary = float64(minutes) * 1000 * 1
	case "B":
		salary = float64(minutes) * 1500 * 1.2
	case "A":
		salary = float64(minutes) * 3000 * 1.5
	}

	return salary
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type MathOperation func(...int) float64

func operation(operator string) (MathOperation, string) {
	ok := "OK"
	switch operator {
	case average:
		return averageFunc1, ok
	case minimum:
		return minFunc1, ok
	case maximum:
		return maxFunc1, ok
	default:
		return nil, "NO OK"

	}
}

func averageFunc1(values ...int) float64 {
	var sum int = 0
	for _, v := range values {
		sum += v
	}
	return float64(sum / len(values))
}
func minFunc1(values ...int) float64 {
	var min int = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return float64(min)
}
func maxFunc1(values ...int) float64 {
	var max int = values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return float64(max)
}
