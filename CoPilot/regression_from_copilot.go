package main

import (
	"github.com/montanaflynn/stats"
	"golang.org/x/exp/errors/fmt"
)

func regress(data []stats.Coordinate) (float64, float64, error) {

	//run the regression
	regression, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, err
	}

	var x0, x1, y0, y1 float64

	//find two different pairs in regression that have different x values
	for i := 0; i < len(regression); i++ { //copilot correctly predicted this line
		if regression[i].X != regression[i+1].X { //copilot did not immediately understand the required if statement, but after typing "if regression[i].X" it correctly predicted the rest of the for statement
			x0 = regression[i].X
			y0 = regression[i].Y
			x1 = regression[i+1].X
			y1 = regression[i+1].Y
			break
		}
	}

	//calculate the slope
	m := (y1 - y0) / (x1 - x0) //copilot correctly predicted this line

	//calcuate the intercept
	b := y0 - m*x0 //copilot correctly predicted this line

	return m, b, nil //copilot correctly predicted this line

}

func main() {

	// Dataset 1
	data1 := []stats.Coordinate{
		{10, 8.04},
		{8, 6.95},
		{13, 7.58},
		{9, 8.81},
		{11, 8.33},
		{14, 9.96},
		{6, 7.24},
		{4, 4.26},
		{12, 10.84},
		{7, 4.82},
		{5, 5.68},
	}

	m1, b1, err := regress(data1) //copilot correctly predicted this line
	if err != nil {               //copilot correctly predicted this line
		fmt.Println(err)
	}
	fmt.Println("Dataset 1:")
	fmt.Printf("y = %.4fx + %.4f\n", m1, b1) //copilot predicted a println statement.  After typing fmt.Printf it correctly predicted the rest of the line

	// Dataset 2
	data2 := []stats.Coordinate{
		{10, 9.14},
		{8, 8.14},
		{13, 8.74},
		{9, 8.77},
		{11, 9.26},
		{14, 8.10},
		{6, 6.13},
		{4, 3.10},
		{12, 9.13},
		{7, 7.26},
		{5, 4.74},
	}

	m2, b2, err := regress(data2) //copilot correctly predicted this line
	if err != nil {               //copilot correctly predicted this line
		fmt.Println(err)
	}
	fmt.Println("Dataset 2:")
	fmt.Printf("y = %.4fx + %.4f\n", m2, b2) //copilot predicted a println statement.  After typing fmt.Printf it correctly predicted the rest of the line

	// Dataset 3
	data3 := []stats.Coordinate{
		{10, 7.46},
		{8, 6.77},
		{13, 12.74},
		{9, 7.11},
		{11, 7.81},
		{14, 8.84},
		{6, 6.08},
		{4, 5.39},
		{12, 8.15},
		{7, 6.42},
		{5, 5.73},
	}

	m3, b3, err := regress(data3) //copilot correctly predicted this line
	if err != nil {               //copilot correctly predicted this line
		fmt.Println(err)
	}
	fmt.Println("Dataset 3:")
	fmt.Printf("y = %.4fx + %.4f\n", m3, b3) //copilot predicted a println statement.  After typing fmt.Printf it correctly predicted the rest of the line

	// Dataset 4
	data4 := []stats.Coordinate{
		{8, 6.58},
		{8, 5.76},
		{8, 7.71},
		{8, 8.84},
		{8, 8.47},
		{8, 7.04},
		{8, 5.25},
		{19, 12.5},
		{8, 5.56},
		{8, 7.91},
		{8, 6.89},
	}

	m4, b4, err := regress(data4) //copilot correctly predicted this line
	if err != nil {               //copilot correctly predicted this line
		fmt.Println(err)
	}
	fmt.Println("Dataset 4:")
	fmt.Printf("y = %.4fx + %.4f\n", m4, b4) //copilot predicted a println statement.  After typing fmt.Printf it correctly predicted the rest of the line

}
