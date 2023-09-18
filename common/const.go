package common

import "fmt"

const (
	DbTypeRestaurant = 1
	DbTypeUser = 2
)

const (
	CurrentUser = "user"
)

func AppRecover() {
	if r := recover(); r != nil {
		fmt.Println("<------------------------------------------------")
		fmt.Println("Recovered: ", r)
		fmt.Println("------------------------------------------------>")
	}
}
