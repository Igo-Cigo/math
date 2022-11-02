package main

import (
	"fmt"
	"math"
)

var pi float64 = 3.14159
var choice int64
var formulas = [7]string{"Area of a circle", "Radius of a circle", "Pythagorean theorem", "Area of a Parallelogram", "Area of a triangle", "Sum of angles", "Amount of angles"}

func main() {
	fmt.Print("Which formula do you want? (Enter the number of the formula)\n\n")
	pick()
	calculation()
}

// https://gosamples.dev/round-float/
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func pick() {
	for i := 0; i < len(formulas); i++ {
		fmt.Print(i+1, ". ", formulas[i], "\n")
	}
	fmt.Print("\n> ")
	fmt.Scan(&choice)
}

func calculation() {
	var inp float64
	var inp2 float64
	var given string
	if choice == 1 {
		fmt.Print("Enter the Radius:")
		fmt.Scan(&inp)
		result := pi * math.Pow(inp, 2)
		fmt.Println("The area of the circle is:", roundFloat(result, 3))
	} else if choice == 2 {
		fmt.Print("Enter the Circumference: ")
		fmt.Scan(&inp)
		result := inp / (pi * 2)
		fmt.Println("The radius of the circle is:", roundFloat(result, 3))
	} else if choice == 3 {
		fmt.Print("Solve for hypotenuse or leg? (h/l) ")
		fmt.Scan(&given)
		if given == "h" {
			fmt.Print("What is the lenght of the legs?")
			fmt.Scan(&inp, &inp2)
			result := math.Sqrt(math.Pow(inp, 2) + math.Pow(inp2, 2))
			fmt.Println("The lenght of the hypotenuse squared is:", roundFloat(result, 3))
		} else if given == "l" {
			fmt.Println("What is the lenght of the hypotenuse and the given leg? (input lenghts by the given order) ")
			fmt.Scan(&inp, &inp2)
			result := math.Sqrt(math.Pow(inp, 2) - math.Pow(inp2, 2))
			fmt.Println("The lenght of the second leg is: ", roundFloat(result, 3))
		}
	} else if choice == 4 {
		//todo: add the area of Parallelogram to calculator
	}

}
