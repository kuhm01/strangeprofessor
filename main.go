package main

import (
	"fmt"
	"os"
	"professorc/routeoption"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("이상한 나라의 교수님은 학점을 짜게 주는거야!\n")
		fmt.Printf("그래서 외쳤어! A 줘\n")
		fmt.Printf("version 0.7.5\n")
	} else if len(args) > 1 {
		routeoption.RouteOption(args)
	} else {
		fmt.Printf("Critical Error!\n")
	}
}
