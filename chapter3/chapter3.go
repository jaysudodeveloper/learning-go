package chapter3

import "fmt"

func Chapter3()  {
	/*Practicing And learning Go Types*/
	fmt.Println("Go Integer types: ")
	fmt.Println("uint8, uint16, uint32, uint64, int8, int16, int32 and int64")
	fmt.Println("Go Float types: ")
	fmt.Println("float32, float64, complex64, complex128")

	fmt.Println("\nExamples:")
	fmt.Println("addition ints   1+1 = ", 1+1)
	fmt.Println("addition floats 1.0+1.0 = ", 1.0+1.0)
	fmt.Println()

	fmt.Println("Strings:")
	//spaces are considered characters
	fmt.Println("Using len() on: Hello world - ", len("Hello world"))
	fmt.Println("Using [] on: Hello world - ", "Hello world"[1]) //using the index access shows 101 as its byte representation
	fmt.Println("Using + on: Hello world - ", "Hello" + " world")

	fmt.Println("Booleans:")
	fmt.Println("&& and")
	fmt.Println("|| or")
	fmt.Println("! not")

}