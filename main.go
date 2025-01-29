package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var messages = []string{
	"Goodbye, cruel world...",
	"You can't fire me, I quit!",
	"This message will self-destruct in 3...2...1...",
	"I was never here.",
	"Oops... I slipped into the void.",
	"Did you really think I’d stay forever?",
}

func main() {
	message := messages[rand.Intn(len(messages))]
	fmt.Println("🛑", message)

	time.Sleep(2 * time.Second)

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("⚠️ Failed to locate the executable.")
		return
	}

	err = os.Remove(exePath)
	if err != nil {
		fmt.Println("⚠️ Uh oh... I tried to disappear, but something stopped me.")
	}
}
