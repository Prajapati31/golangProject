package main

import (
	"fmt"
	"net/http"
)

// Division function
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// Multiplication function
func multiply(a, b float64) float64 {
	return a * b
}

// Addition function
func add(a, b float64) float64 {
	return a + b
}

// Subtraction function
func subtract(a, b float64) float64 {
	return a - b
}

// Hello World function
func helloWorld(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello, World!</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
		</body>
		</html>
	`

	fmt.Fprint(w, html)
}

// FlyHigh function
func flyHigh(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Fly High!</title>
		</head>
		<body>
			<h1>Fly High!</h1>
		</body>
		</html>
	`

	fmt.Fprint(w, html)
}

// MyName function
func myName(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>My Name</title>
		</head>
		<body>
			<h1>My Name</h1>
		</body>
			</html>
	`

	fmt.Fprint(w, html)
}

func main() {
	// Division
	result := divide(10, 2)
	fmt.Println("Division:", result)

	// Multiplication
	result = multiply(5, 6)
	fmt.Println("Multiplication:", result)

	// Addition
	result = add(3, 4)
	fmt.Println("Addition:", result)

	// Subtraction
	result = subtract(8, 5)
	fmt.Println("Subtraction:", result)

	// HTTP server setup
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/flyhigh", flyHigh)
	http.HandleFunc("/myname", myName)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
