package view

import "fmt"

func PrintError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("[ERROR] %s\n", err.Error())
}
