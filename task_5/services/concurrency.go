package services

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ShutDownRequest(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Program Stopped")
			time.Sleep(time.Second)
			os.Exit(0)
			return
		default:
			fmt.Println("Program is running")
			time.Sleep(time.Second)

		}
	}
}

func StartRandomNumber(stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Stopping")
			time.Sleep(time.Second)
			fmt.Println("Generating stopped")
			return
		default:
			number := rand.Int()
			fmt.Println("Printing", number)
			time.Sleep(time.Second)
		}
	}
}
