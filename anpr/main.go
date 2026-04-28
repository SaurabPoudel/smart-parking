package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sendInterval = time.Second

func genEvent() string {
	events := []string{"ENTRY", "EXIT"}
	return events[rand.Intn(len(events))]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println(genEvent())
		time.Sleep(sendInterval)
	}
}
