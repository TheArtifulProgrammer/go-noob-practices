package main

import "fmt"

// Features
// - Strong and statistically typed
// - Fast Compile time
// - Automatic garbage collection
// - Goroutines and channels for concurrency
// - Multiple programming paradigms supported (OOP, FP, etc.)

// Limitations
// - No support for generics

// := declaring and initializing variables in a single statement, can only be used inside functions, not for package-level variables
// %v for actual value
// %T for type

// func main()  {
// 	a := "Hello World"
// 	fmt.Printf("%v %T", a, a)
// }

// func main()  {
// 	var (
// 		movie_name string = "Interstellar"
// 		movie_rating float32 = 8.5
// 		movie_year int = 2014
// 	)

// 	fmt.Printf("%v %T\n", movie_name, movie_name)
// 	fmt.Printf("%v %T\n", movie_rating, movie_rating)
// 	fmt.Printf("%v %T\n", movie_year, movie_year)

// }
func main() {
	// const (
	// 	a int = iota
	// 	b int = iota
	// 	c int = iota
	// )
	// fmt.Println(a, b, c)

	// var marks [8]int

	// for i := 0; i < 8; i++ {
	// 	marks[i] = i
	// }

	// fmt.Println(marks[6])
	// fruits := [3]string{"apple", "banana", "orange"}
	// fruits[2] = "grape"
	// fmt.Println(fruits)
	// maps
	// fruit_basket := map[string]int{"apple": 5, "banana": 10, "orange": 15}
	// fmt.Println(fruit_basket)
	getGreeting := greetings

	fmt.Println(getGreeting("John", getAge))

}

func greetings(name string, getAge func() int) string {
	return fmt.Sprintf("Hi, my name is %v and I am %v years old", name, getAge())
}

func getAge() int {
	return 24
}
