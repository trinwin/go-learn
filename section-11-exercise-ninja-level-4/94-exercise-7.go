package main

import "fmt"

func main() {
	for by := 1997; by <= 2021; by++ {
		if by == 2021 {
			fmt.Println("Current year", by)
		} else if by == 2020{
			fmt.Println("Last year", by)
		} else {
			fmt.Println(by)
		}
	}
}

