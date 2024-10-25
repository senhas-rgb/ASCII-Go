package main

import (
	"os"
	"os/exec"
	"time"
	"fmt"
)

func ScreenClear(period int) {
	time.Sleep(time.Duration(period) * time.Millisecond)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func frame(character string) {
	bottom := [32]string{}
	middle := [32]string{}
	top := [32]string{}
	i := 0
	for {
		top[i] = "-"
		bottom[i] = "-"
		middle[i] = " "
		i++
		if i == 32 {
			break
		}
	}
	i = 0
	middle[0] = character
	position := 1
	for {
		ScreenClear(0)
		fmt.Println(top)
		fmt.Println(middle)
		fmt.Println(bottom)
		ScreenClear(200)
		i++
		if i == 32 {
			break
		}
		middle[position] = character
		middle[position-1] = " "
		position++
	}

}

func main() {
	character := "o"
	frame(character)
}
