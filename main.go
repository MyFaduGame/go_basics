package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Go")
	// constants()
	// for_loop()
	// if_else()
	// switch_case()
	arrays_example()
}

func constants() {
	fmt.Println("Constant declarations")
	const s string = "Constant"
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func for_loop() {
	i := 1
	fmt.Println("Only Conditional For Loop")
	for i <= 3{  // only condition for loop
		fmt.Println(i)
		i = i+1
	}

	fmt.Println("Inline Condition For Loop")
	for j:=3; j>=1; j-- { // inline declaration for loop
		fmt.Println(j)
	}

	fmt.Println("range for loop like python")
	for i:= range 3 { // direct range for loop, range is always 1 less then given ex 3=2
		fmt.Println(i)
	}

	fmt.Println("without condition for loop and break statement")
	for { //break we can delclare for loop without condition also in go
		fmt.Println("break")
		break
	}

	fmt.Println("Conditional for loop and continue statement")
	for n:= range 6 { //for loop with condition
		if n%2==0 {
			continue
		}
		fmt.Println(n)
	}

}

func if_else() {
	
	fmt.Println("normal if else statement")
	if 7%2 == 0{
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	fmt.Println("single if statement")
	if 8%4==0 {
		fmt.Println(true)
	}

	fmt.Println("Or Condition")
	if 8%4 == 0 || 7%2 == 0 {
		fmt.Println("Inside the If Statement")
	}

	fmt.Println("And Condition")
	if 8%4 == 0 && 7%2 == 0 {
		fmt.Println("this statement never gonna print")
	}

	fmt.Println("If else if statements")
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

}

func switch_case() {
	i:= 2
	fmt.Println("write",i,"as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t:=i.(type) {
		case bool:
			fmt.Println("Its bool")
		case int:
			fmt.Println("Its an int")
		default:
			fmt.Printf("maybe %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("new")

}

func arrays_example() {
	var a [5]int
	fmt.Println(a)
}
