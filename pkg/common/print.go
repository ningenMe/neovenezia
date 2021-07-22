package common

import "fmt"

func Print(messages []string) {
	for _, value := range messages {
		fmt.Println(value)
	}
}

func PrintOne(message string)  {
	Print([]string{message})
}

