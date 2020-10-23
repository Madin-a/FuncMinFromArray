package main

import "fmt"

/*func MaxFromArray(numbers [5]int )  {
	max := numbers[0]
	for _, i:=range numbers{
		if i > max {
			max = i
		}
	}
	fmt.Println("max = ", max)
}*/
func MinFromArray(numbers [5]int )  {
	min := numbers[0]
	for _, i:=range numbers{
		if i < min {
			min = i
		}
	}
	fmt.Println("min = ",min)
}
func main() {
	MinFromArray([5]int{3,2,-8,7,9})
	//MaxFromArray([5]int{3,2,-8,7,9})
}
