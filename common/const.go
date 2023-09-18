package common

import "fmt"

const (
	DbTypeRestaurant = 1
	DbTypeUser = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if r := recover(); r != nil {
		fmt.Println("<------------------------------------------------")
		fmt.Println("Recovered: ", r)
		fmt.Println("------------------------------------------------>")
	}
}
