
package main

import "fmt" 

func main() {
	var rows int = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(fmt.Sprintf("%d ", j))
		}
		fmt.Println("")
	}
}
