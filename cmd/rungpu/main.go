package main

import (
	"GPUContainerRunnerGo/pkg/cui_input"
	"fmt"
)

func main() {
	//var test string
	//fmt.Scan(&test)
	//fmt.Println(test)
	scriptInfo := cui_input.InputUser()
	fmt.Println(scriptInfo)
}
