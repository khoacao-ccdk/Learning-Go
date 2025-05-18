package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func loop() {
	sum := 1
	for {
		sum += sum
		if sum > 100 {
			break
		}
	}
	fmt.Println("Finished Looping, sum is:", sum)
}

func pow(x, n, lim float64) float64 {
	// Go accepts parameter init in the conditional statement
	// The variable will be scoped to the if statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // Else statement must be on the same line as the closing bracket
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func sqrt(x float64) float64 {
	if x < 0 {
		return 0
	}

	guess := float64(x / 2)

	for {
		guess -= (guess*guess - x) / (2 * guess)
		if (guess*guess - x) < 0.0000001 {
			break
		}
	}
	return guess
}

func getRuntimeOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return `Unknown OS: ${os}`
	}
}

func getTimeOfDay() string {
	hour := time.Now().Hour()

	// This essentially acts like a long if-else statement
	switch {
	case hour < 12:
		return "Good Morning"
	case hour < 17:
		return "Good Afternoon"
	case hour < 21:
		return "Good Evening"
	default:
		return "Good Night"
	}
}

func deferFunc() {
	// Defer statements are put on a stack and executed in LIFO order
	defer fmt.Println("Deferred function executed")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
}

func Module2() {
	fmt.Println("===MODULE 2===")
	loop()

	fmt.Println(
		"If Statement with init:",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		"Square Root of 149:",
		sqrt(149),
	)

	fmt.Println(
		"Current OS:",
		getRuntimeOS(),
	)

	fmt.Println(
		"Greeting:",
		getTimeOfDay(),
	)

	deferFunc()
}
