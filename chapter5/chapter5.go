package chapter5

import "fmt"

func Chapter5()  {
	/*Practicing and learning Control structures*/

	//The for statement allows us to repeat many operations repeatedly
	//First create a variable that we want to print and increment
	i := 0

	//Then create the for block, also known as a while loop
	fmt.Println("This is a for loop in go:")
	for i <= 10 {
		fmt.Println(i)
		i = i + 1 //increment the value of i by 1
	}

	//The above loop can be written more concisely
	fmt.Println("This is a more concisely written for loop")
	for i := 0; i <= 10; i++ { // notice how we use i++ to increment i for us
		fmt.Println(i)
	}

	//Go does not have while or do while loops but the for loop block can be rewritten to emulate these
	//While loop in go is just for <evaluated boolean expression> { //block content }

	fmt.Println("This is just a for loop without a conditional but usese a break to exit early")
	for {
		fmt.Println("Breaking out")
		break
	}

	fmt.Println("Lets look at some switch and if examples")
	fmt.Println("Lets use if/else to print odd evens")

	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	fmt.Println("If blocks can also contain assignment expresions before the conditional")
	for i := 0; i <= 10; i++ {
		if mod := i % 2 ; mod == 0 {//we assign the result of mod operation to mod and use it to evaluate an expression
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	fmt.Println("Lets use switch to print odd evens")

	for i := 0; i <= 10; i++ {
		switch i % 2 {
		case 0: fmt.Println("Even")
		default: fmt.Println("Odd")
		}
	}

	fmt.Println("Switch blocks can also have coma seperated cases:")
	for i := 0; i <= 10; i++ {
		switch i {
		case 0, 1, 2, 3: fmt.Println("0,1,2,3 cases caught")
		default: fmt.Println("The rest")
		}
	}

	fmt.Println("Switch without an expression is another way to express if/else logic")
	switch {
	case 1 < 2:
		fmt.Println("1 < 2")
	default:
		fmt.Println("Default case")
	}

	fmt.Println("We can also use switch to compare types instead of values")
	it := 1
	b := true
	s := "string"
	whatAmI(it) //we will come back to declaring functions and function expressions
	whatAmI(b)
	whatAmI(s)
	whatAmI(5.5)
}

func whatAmI(i interface{})  {
	switch t := i.(type) {
	case bool:
		fmt.Println("im a bool")
	case int:
		fmt.Println("im a int")
	case string:
		fmt.Println("im a string")
	default:
		fmt.Printf("i dont know this type: %T", t)
	}
}