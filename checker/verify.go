package main

import (
	"fmt"
	"os"
)

func judge() bool {
	contestant := "output.txt"
	judge := "judge.txt"

	content1, err1 := os.ReadFile(contestant)
	if err1 != nil {
		saveError(fmt.Sprintf("Error reading %s : %v", contestant, err1))
		os.Exit(1)
	}
	content2, err2 := os.ReadFile(judge)
	if err2 != nil {
		saveError(fmt.Sprintf("Error reading %s : %v", judge, err2))
		os.Exit(1)
	}
	saveError(string(content1))
	saveError(string(content2))
	if string(content1) == string(content2) {
		return true
	} else {
		return false
	}
}
func saveError(errMsg string) {
	errfilepath := "error.txt"
	f, err := os.OpenFile(errfilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to write to error.txt: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(errMsg + "\n"); err != nil {
		fmt.Printf("Failed to write to error.txt: %v\n", err)
	}
}
