package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	channel := make(chan bool)
	evilNinjas := []string{"Michael", "Johnny", "Tommy", "Timmy", "Toto"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, channel)
	}
	checkDone := <- channel
	fmt.Println(checkDone)

}


func attack(target string, attacked chan bool) {
	fmt.Println("throwing stars at", target)
	time.Sleep(time.Second *2)
	attacked  <- true
}
