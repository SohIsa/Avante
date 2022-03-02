package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var tellMe string
	fmt.Printf("enter a string: ")
	//fmt.Scan(&tellMe)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tellMe = scanner.Text()
		break
	}

	tellMe = strings.ToLower(tellMe)
	//fmt.Printf(strings.ReplaceAll(tellMe, " ", ""))

	if strings.HasPrefix(tellMe, "i") &&
		strings.HasSuffix(tellMe, "n") &&
		strings.Contains(tellMe, "a") {
		fmt.Printf("FOUND!")
	} else {
		fmt.Printf("Not Found!")
	}
}
