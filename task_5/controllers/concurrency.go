package controllers

import (
	"fmt"
	"keraton_task_5/helpers"
	"keraton_task_5/services"

	"github.com/gofiber/fiber/v2"
)

// var numbers = make(chan int)
var stopGenerating = make(chan struct{})

func ShutDownRequest(c *fiber.Ctx) error {
	//Initiate Stop Channel Signal
	stopChan := make(chan struct{})

	fmt.Println("Shutdown requested via API")
	go services.ShutDownRequest(stopChan)
	close(stopChan)
	return helpers.SuccessResponse(c, 200, "Shutdown request recieved. Shutting down...", "Back to terminal")
}

func StartRandomNumber(c *fiber.Ctx) error {
	fmt.Println("Generating...")
	go services.StartRandomNumber(stopGenerating)

	return helpers.SuccessResponse(c, 200, "Generating number in background", "Check your terminal")

}

func StopRandomNumber(c *fiber.Ctx) error {
	close(stopGenerating)
	return helpers.SuccessResponse(c, 200, "Stopped Succesfully", "Check your terminal")

	//Issues : if random number is not generated, then it will panic because channel is closed.
	//I cannot find the solution at this time

}
