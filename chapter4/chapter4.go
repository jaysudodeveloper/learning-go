package chapter4

import "fmt"

var outerScope string = "this variable is available to all functions and inside main"

func Chapter4()  {
	/*Practicing and learning Variable assignment*/

	var innerScope string = "this variable is available to main but not outside of main"
	fmt.Println(innerScope)

	fmt.Println(outerScope)

	//There are two ways to assign variables
	//var <variable name> <type> = <data>
	var var1 string = "Long format assignment: var var1 string"
	fmt.Println(var1)

	//<variable name> := <data>
	var2 := "Short format assignment: var2 :="
	fmt.Println(var2)

	//constants allow us to assign variables that cannot change
	const imConstant string = "I cannot change this variable"
	fmt.Println(imConstant)

	fmt.Println("Multiple variable assignment format: var a,b,c int = 1,2,3")
	//multiple variable assignment
	var (
		a = 5
		b = 6
		c = 7
	)
	//if all variables have same type we can save time by only using it once at the end of the last variable
	var aa, bb, cc int = 1, 2, 3

	fmt.Println(a,b,c, aa,bb,cc)
}
