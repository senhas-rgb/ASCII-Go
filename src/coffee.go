package main

import (
	"fmt"
)


func coffee() {
	smoke := [2]string{"&", "%"}
	partitcle := [3]string{")", "(", "*"}
	pointer_particle := 0
	pointer_smoke := 0
	partitcles1 := [12]string{}
	partitcles2 := [12]string{}
	partitcles3 := [12]string{}
	smoke1 := [12]string{}
	smoke2 := [12]string{}
	i := 0
	s := 0
	for {
		smoke1[s] = " "
		smoke2[s] = " "
		s++
		if s == 12 {
			break
		}
	}

	for {
		partitcles1[i] = " "
		partitcles2[i] = " "
		partitcles3[i] = " "
		i++
		if i == 12 {
			break
		}
	}
	i = 2
	j := 2
	z := 11
	smoking := 1
	for {
		ScreenClear(0)
		fmt.Println("###########################")
		fmt.Println(smoke1)
		fmt.Println(smoke2)
		fmt.Println(partitcles1)
		fmt.Println(partitcles2)
		fmt.Println(partitcles3)
		fmt.Println("     _____________         ")
		fmt.Println("    <_____________> ___    ")
		fmt.Println("    |             |/ _ /   ")
		fmt.Println("    |               | | |  ")
		fmt.Println("    |               |_| |  ")
		fmt.Println(" ___|             |/___/   ")
		fmt.Println("/    /___________/    /    ")
		fmt.Println("|_____________________/    ")
		fmt.Println("###########################")
		ScreenClear(500)
		if i == 8 || j == 8 {
			partitcles1[i-1] = " "
			partitcles2[j-1] = " "
			partitcles3[j-1] = " "
			i = 2
			j = 2
		} 
		if z == 1 {
			z = 11
		}
		if pointer_particle == 3 {
			pointer_particle = 0
		}
		partitcles1[i] = partitcle[pointer_particle]
		partitcles2[j] = partitcle[pointer_particle]
		partitcles3[i] = partitcle[pointer_particle]
		partitcles1[i-1] = " "
		partitcles2[j-1] = " "
		partitcles3[j-1] = " "
		partitcles1[z] = partitcle[pointer_particle]

		if pointer_smoke == 2 {
			pointer_smoke = 0
		}

		smoke1[smoking] = smoke[pointer_smoke]
		smoke2[smoking] = smoke[pointer_smoke]

		if smoking == 10 {
			smoking = 3
		}


		j++
		i++
		z--
		pointer_particle++
		pointer_smoke++
		smoking++

	}

}
func coffees() {
	coffee()
}
