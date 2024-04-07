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
	// arrays_example()
	// slices_example()
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

	//seting value to the array
	a[4] = 100
	fmt.Println("set",a)

	//to print specific element
	fmt.Println("get",a[4])

	//smillar as rest of arrays
	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}


func slices_example() {

    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

	/*
	To create an empty slice with non-zero length, use the builtin make. 
	Here we make a slice of strings of length 3 (initially zero-valued). 
	By default a new slice’s capacity is equal to its length; 
	if we know the slice is going to grow ahead of time, 
	it’s possible to pass a capacity explicitly as an additional parameter to make.
	*/

    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
