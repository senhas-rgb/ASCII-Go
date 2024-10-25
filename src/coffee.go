package main

import (
	"fmt"
)


func coffee() {
	partitcle := [3]string{")", "(", "*"}
	pointer := 0
	partitcles1 := [12]string{}
	partitcles2 := [12]string{}
	i := 0
	for {
		partitcles1[i] = " "
		partitcles2[i] = " "
		i++
		if i == 12 {
			break
		}
	}
	i = 4
	j := 5
	for {
		ScreenClear(0)
		fmt.Println("###########################")
		fmt.Println(partitcles1)
		fmt.Println(partitcles2)
		fmt.Println("     _____________         ")
		fmt.Println("    <_____________> ___    ")
		fmt.Println("    |             |/ _ /   ")
		fmt.Println("    |               | | |  ")
		fmt.Println("    |               |_| |  ")
		fmt.Println(" ___|             |/___/   ")
		fmt.Println("/    /___________/    /    ")
		fmt.Println("|_____________________/    ")
		fmt.Println("###########################")
		ScreenClear(200)
		if i == 12 || j == 12 {
			partitcles1[i-1] = " "
			partitcles2[j-1] = " "
			i = 2
			j = 3
		}
		if pointer == 3 {
			pointer = 0
		}
		partitcles1[i] = partitcle[pointer]
		partitcles2[j] = partitcle[pointer]
		partitcles1[i-1] = " "
		partitcles2[j-1] = " "
		j++
		i++
		pointer++
	}

}
func main() {
	coffee()
}
