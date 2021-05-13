package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	
	}
	//range 
	names := []string{"ashutosh" , "chrisal", "tahir"}
	for x := 0; x <len(names); x++ {
		fmt.Println(names[x])
	}
	for k , v := range names{
		fmt.Println(k,v)
	}
	// fallthrough keyword
	j := 45

    switch {
    case j < 10:
        fmt.Println("i is less than 10")
    case j < 50:
        fmt.Println("i is less than 50")
    case j < 100:
        fmt.Println("i is less than 100")
    }


}
