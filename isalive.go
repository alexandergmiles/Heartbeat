package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	currentTime := time.Now();
	fmt.Println("Heartbeat has been built and running at: ", currentTime.String())
    
	for {
		loopTime := time.Now();
		
		fmt.Printf("[%v] Program is still alive and running under the process %v \n", loopTime.Format("2006-01-02 15:04:05"), os.Getpid())
		time.Sleep(5*time.Second);
	}
}