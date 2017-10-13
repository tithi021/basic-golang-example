/*
* Run this example from terminal by
* using this command - go run switch.go
*
* Output:
	Write 2 as Two
	It's a weekday
	It's after noon
	I'm a bool
	I'm a int
	Dont know type string
*/

package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")

	default:
		fmt.Println("It's after noon")
	}

	whatAmI :=
		func(i interface{}) {
			switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")

			case int:
				fmt.Println("I'm a int")

			default:
				fmt.Printf("Dont know type %T\n", t)
			}
		}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
