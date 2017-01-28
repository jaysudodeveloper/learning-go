package chapter6

import "fmt"

func Chapter6()  {
	/*Practicing and learning Arrays, slices and Maps*/

	fmt.Println()
	fmt.Println("Lets play with some arrays and slices and maps:")

	fmt.Println("Some arrays:")
	var arr1 [5]int // here we create an array if exactly 5 elements
	fmt.Println("emp:", arr1)

	//use this to declare and initialize an array with data
	var arr2 = [5]int{1,2,3,4,5}
	arr3 := [5]int{1,2,3,4,5} //short hand declaration
	var arrCount = [...]int{1,2,3,4,5,6,7,8} // we can have the compiler count the items for us

	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("compiler counted the elements for us in this array: ",  arrCount)

	var arr4 [5][5]int
	fmt.Println(arr4)

	fmt.Println("Some slices:")
	//unlike arrays, slices are typed only by the elements they contain
	//make a slice like so:
	slc1 := make([]int, 3) // here we make a slice with non-zero length of integers
	fmt.Println("emp:", slc1)

	slc1[0] = 1
	slc1[1] = 2
	slc1[2] = 3

	fmt.Println("len() returns length of slice:", len(slc1))

	fmt.Println("we can append to a slice using append(<slice>, <data>)")
	slc1 = append(slc1, 4)
	fmt.Println(slc1)

	//two dimensional slices are declared similar to arrays
	fmt.Println("two dimensional slice")
	slc2 := make([][]int, 5)
	fmt.Println(slc2)

	fmt.Println("Some maps:")
	//we make a map using the builtin make() function
	m1 := make(map[string]int)

	//then we can access and add data to it using a map[key] = value syntax
	m1["1"] = 1
	m1["2"] = 2

	fmt.Println("map: ", m1)
	fmt.Println("we can also call len on a map to get number of items inside: ", len(m1))

	//built in delete method removes the key:val pairs from the map
	delete(m1, "1")
	fmt.Println("we delete the key:val pair \" 1\" ", m1)

	//we can also use the optional return value to indicate if a key was present or not
	_, prs := m1["1"] // here we use _ to ignore the value in favour of is present indicator
	fmt.Println("is key 1 present: ", prs)

	//we can also declare maps and initialize them similar to arrays/slices
	m2 := map[string]int{"1": 1, "2": 2}
	fmt.Println(m2)

	//a map of maps is also possible
	m3 := make(map[string]map[string]int)
	fmt.Println("this is a map of maps: ", m3)
}