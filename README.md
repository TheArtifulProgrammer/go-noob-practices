# Learning Go

This repository documents my journey learning the Go programming language, including code examples and explanations of key concepts.

## Features of Go

- Strong and statically typed
- Fast compile time
- Automatic garbage collection
- Goroutines and channels for concurrency
- Multiple programming paradigms supported (OOP, FP, etc.)

## Concepts Covered

### Variables and Data Types

```go
func variablesExample() {
    // Using := for declaration and initialization
    a := "Hello World"
    fmt.Printf("%v %T\n", a, a)

    // Multiple variable declaration
    var (
        movieName string = "Interstellar"
        movieRating float32 = 8.5
        movieYear int = 2014
    )
    fmt.Printf("%v %T\n", movieName, movieName)
    fmt.Printf("%v %T\n", movieRating, movieRating)
    fmt.Printf("%v %T\n", movieYear, movieYear)
}
```
### Constants and Iota

```go
const (
    a = iota
    b
    c
)
fmt.Println(a, b, c) // Outputs: 0 1 2
```
### Arrays and Slices

```go
func arraysAndSlicesExample() {
    // Array
    var marks [8]int
    for i := 0; i < 8; i++ {
        marks[i] = i
    }
    fmt.Println(marks[6])

    // Slice
    fruits := []string{"apple", "banana", "orange"}
    fruits[2] = "grape"
    fmt.Println(fruits)
}
```
### Maps

```go
fruitBasket := map[string]int{
    "apple":  5,
    "banana": 10,
    "orange": 15,
}
fmt.Println(fruitBasket)
```
### Functions and Higher-Order Functions

```go
func greetings(name string, getAge func() int) string {
    return fmt.Sprintf("Hi, my name is %v and I am %v years old", name, getAge())
}

func getAge() int {
    return 24
}

func main() {
    fmt.Println(greetings("John", getAge))
}
```
## Additional Important Concepts
### Pointers

```go
func pointersExample() {
    x := 10
    p := &x
    fmt.Println("Value:", x)
    fmt.Println("Pointer:", p)
    fmt.Println("Dereferenced value:", *p)
    *p = 20
    fmt.Println("New value:", x)
}
```
### Structs and Methods

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I'm %d years old\n", p.Name, p.Age)
}

func structsExample() {
    person := Person{Name: "Alice", Age: 30}
    person.SayHello()
}
```

### Interfaces

```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func interfacesExample() {
    var s Shape = Rectangle{Width: 5, Height: 10}
    fmt.Printf("Area: %.2f\n", s.Area())
}
```
### Error Handling

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func errorHandlingExample() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### Goroutines and Channels

```go
func goroutinesExample() {
    ch := make(chan string)
    go func() {
        ch <- "Hello from goroutine!"
    }()
    msg := <-ch
    fmt.Println(msg)
}
```

### Defer, Panic, and Recover

```go
func deferPanicRecoverExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Starting the function")
    panic("Something went wrong!")
    fmt.Println("This line will not be executed")
}
```
## Running the Code

To run any of the examples, use the following command:

```go
go run main.go
```
## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)

## Next Steps

- Build small projects to apply learned concepts
- Explore Go's standard library
- Learn about Go modules and dependency management
- Practice writing concurrent programs using goroutines and channels
- Contribute to open-source Go projects

