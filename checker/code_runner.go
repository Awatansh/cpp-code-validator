package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runner() {
	cppfile := "./source/code.cpp"
	inputFile := "input.txt"
	outputFile := "output.txt"
	executable := "code.out"
	errorFile := "error.txt"
	errfile, err := os.Create(errorFile)
	if err != nil {
		fmt.Printf("Error file cannot be opened: %v\n", err)
	}
	defer errfile.Close()
	cmd := exec.Command("g++", cppfile, "-o", executable)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Compilation failed : %v\n", err)
		os.Exit(1)
	}
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Failed to open input.txt\n")
		os.Exit(1)
	}
	defer input.Close()
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to Create output.txt\n")
		os.Exit(1)
	}
	defer output.Close()
	cmd = exec.Command("./" + executable)
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = errfile
	if err := cmd.Run(); err != nil {
		fmt.Printf("Executation failed: %v,\n", err)
		os.Exit(1)
	}
	if err := os.Remove(executable); err != nil {
		fmt.Printf("Failed to remove executable: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Program executed successfully. Output written to output.txt. All errors saved to error.txt\n")
}
