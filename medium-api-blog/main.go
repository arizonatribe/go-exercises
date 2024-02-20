package main

import (
	"fmt"
	"log"
)

func main() {
	service := NewFacebookService()
	userProfile, apiError, err := service.GetUserDetails()
	if err != nil {
		log.Fatalln(err)
	}
	if apiError != nil {
		log.Fatalln(apiError.Error.Message)
	}
	fmt.Printf("User:\n\t%v\n", userProfile)
}
