package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

/* TODO
O: 't' for showing current time
O: show percent of time
?: use type times []time.Time
*/

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer keyboard.Close()

	var exit bool
	var distracted []time.Time
	working := []time.Time{time.Now()}

	fmt.Println("Space for flipping status, enter to get times")

	printStatus := func() {
		fmt.Print("\rStatus: ")
		if len(working)&1 == 1 {
			fmt.Print("Working")
		} else {
			fmt.Print("Distracted")
		}
		fmt.Print("          ")
	}
	printStatus()

	for !exit {
		_, key, _ := keyboard.GetKey()
		now := time.Now()
		switch key {
		case keyboard.KeySpace:
			working = append(working, now)
			distracted = append(distracted, now)
			printStatus()
		case keyboard.KeyEnter:
			exit = true
			if len(working)&1 == 1 { // odd
				working = append(working, now)
			} else {
				distracted = append(distracted, now)
			}
		}
	}

	var workingDuration time.Duration
	var distractedDuration time.Duration

	for i := range working {
		if i&1 == 0 { // even
			workingDuration += working[i+1].Sub(working[i])
		}
	}
	for i := range distracted {
		if i&1 == 0 { // even
			distractedDuration += distracted[i+1].Sub(distracted[i])
		}
	}

	fmt.Printf("\nZorche's Working Time %v\n", workingDuration)
	fmt.Printf("Zorche's Distracted Time %v\n", distractedDuration)
}
