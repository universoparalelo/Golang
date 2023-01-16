package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi %v, how are you?", name)
	return message
}
