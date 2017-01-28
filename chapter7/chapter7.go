package chapter7

import "fmt"

func Chapter7() {
	//create a float slice for our function to use
	xs := []float64{98,93,77,82,83}

	//print the results
	fmt.Println(average(xs))

	fmt.Println("This used a named return value", namedReturn())

	a, b := multipleReturn() // we can also use the _ syntax to discard an un needed value
	fmt.Println("These were returned from a multi return function", a,b)
}

//our first function
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//our functions can also name the return type for an implicit return of that variable
func namedReturn() (name float64) {
	name = 5 //assign our named return variable a value
	return
}

//multiple return values
func multipleReturn() (int, int) {
	return 5,6
}

