package main

import (
	"fmt"
	"math"
)

var pi float64 = 3.14159
var choice int64
var formulas = [7]string{"Area of a circle", "Radius of a circle", "Pythagorean theorem", "Area of a Parallelogram", "Area of a triangle", "Sum of angles of a polygon", "Amount of diagonals of a polygon"}
var inp float64
var inp2 float64
var given string

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
	switch choice {
	case 1:
		calc1()
	case 2:
		calc2()
	case 3:
		calc3()
	case 4:
		calc4()
	case 5:
		calc5()
	case 6:
		calc6()
	case 7:
		calc7()
	}
	// highly inneficient :D
}

func calc1() {
	fmt.Print("Enter the Radius:")
	fmt.Scan(&inp)
	result := pi * math.Pow(inp, 2)
	fmt.Println("Area of the circle is:", roundFloat(result, 3))
}
func calc2() {
	fmt.Print("Enter the Circumference: ")
	fmt.Scan(&inp)
	result := inp / (pi * 2)
	fmt.Println("Radius of the circle is:", roundFloat(result, 3))
}
func calc3() {
	fmt.Print("Solve for hypotenuse or leg? (h/l) ")
	fmt.Scan(&given)
	if given == "h" {
		fmt.Print("What is the lenght of the legs?")
		fmt.Scan(&inp, &inp2)
		result := math.Sqrt(math.Pow(inp, 2) + math.Pow(inp2, 2))
		fmt.Println("Lenght of the hypotenuse squared is:", roundFloat(result, 3))
	} else if given == "l" {
		fmt.Println("What is the lenght of the hypotenuse and the given leg? (input lenghts by the given order) ")
		fmt.Scan(&inp, &inp2)
		result := math.Sqrt(math.Pow(inp, 2) - math.Pow(inp2, 2))
		fmt.Println("Lenght of the second leg is: ", roundFloat(result, 3))
	}
}
func calc4() {
	fmt.Println("Enter the base and the lenght (input lenghts by the given order)")
	fmt.Scan(&inp, &inp2)
	result := inp * inp2
	fmt.Println("Area of the parallelogram is: ", roundFloat(result, 3))
}
func calc5() {
	fmt.Println("Enter the base and the height (input lenghts by the given order)")
	fmt.Scan(&inp, &inp2)
	result := (inp * inp2) / 2
	fmt.Println("Area of the triangle is: ", roundFloat(result, 3))
}
func calc6() {
	fmt.Println("Enter the amount of angles")
	fmt.Scan(&inp)
	result := (inp - 2) * 180
	fmt.Print("Sum of angles of the polygon is: ", roundFloat(result, 3), "°\n")
}
func calc7() {
	fmt.Println("Enter the amount of angles")
	fmt.Scan(&inp)
	result := (inp * (inp - 3)) / 2
	fmt.Println("Amount of diagonals of the polygon is: ", roundFloat(result, 3))
}
