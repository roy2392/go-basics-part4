package main

import "fmt"

type Product struct {
	id    string
	title string
	price float32
}

func main() {
	//1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"read", "code", "swim"}
	// 		Output (print) that array in the command line.
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1:3])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	dynamicArray := []string{"creatre an API", "build SQL"}
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicArray[1] = "build stuff"
	dynamicArray = append(dynamicArray, "learn basics")
	fmt.Println(dynamicArray)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		Product{
			"first-product",
			"A First Product",
			12.99,
		},
		Product{

			"second-product",
			"A Secod Product",
			12.99,
		},
	}
	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

//

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
