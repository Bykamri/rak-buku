package main

import "fmt"

func main() {
	fmt.Println("Hello Worlds!")	
	
    limit := 1000
    oddNumbers := generateOddNumbers(limit)

    fmt.Println(oddNumbers)	

}

func generateOddNumbers(limit int) []int {
    var odds []int
    for i := 1; i <= limit; i += 2 {
        odds = append(odds, i)
    }
    return odds
}