package common

import "fmt"

func AppRecover() {
	if r := recover(); r != nil {
		fmt.Println("<------------------------------------------------")
		fmt.Println("Recovered: ", r)
		fmt.Println("------------------------------------------------>")
	}
}
