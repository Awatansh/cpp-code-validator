package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func spinner(codecontent string) (bool, string) {
	tempDir, err := os.MkdirTemp("", "code-*")
	if err != nil {
		return false, fmt.Sprintf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)
	sourceDir := filepath.Join(tempDir, "source")
	err = os.Mkdir(sourceDir, 0755)
	if err != nil {
		return false, fmt.Sprintf("Error creating source directory: %v", err)
	}
	// Copy required files to tempDir
	files := []string{"input.txt", "output.txt", "error.txt", "judge.txt", "code_runner"}
	for _, file := range files {
		src := filepath.Join(".", file)
		dst := filepath.Join(tempDir, file)
		content, err := os.ReadFile(src)
		if err != nil {
			return false, fmt.Sprintf("Error reading %s: %v", file, err)
		}
		err = os.WriteFile(dst, content, 0644)
		if err != nil {
			return false, fmt.Sprintf("Error copying %s: %v", file, err)
		}
	}
	codePath := filepath.Join(sourceDir, "code.cpp")
	err = os.WriteFile(codePath, []byte(codecontent), 0644)
	if err != nil {
		return false, fmt.Sprintf("Error writing code.cpp: %v", err)
	}
	//Make executable
	cmd := exec.Command("chmod", "+x", filepath.Join(tempDir, "code_runner"))
	err = cmd.Run()
	if err != nil {
		return false, fmt.Sprintf("Error making code_runner executable: %v", err)
	}
	cmd = exec.Command("./code_runner")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running code_runner: %v\nOutput: %s", err, output)
		return false, fmt.Sprintf("Error running code_runner: %v\nOutput: %s", err, output)
	}
	fmt.Printf("output : %v", string(output))
	fmt.Printf("err : %v", err)
	// print("Here\n")
	errorContent, err := os.ReadFile(filepath.Join(tempDir, "output.txt"))
	if err != nil {
		return false, fmt.Sprintf("Error reading output.txt: %v", err)
	}
	fmt.Printf("Error content: %s\n", string(errorContent))
	outputStr := string(output)
	outputLines := strings.Split(outputStr, "\n")
	if len(outputLines) > 0 {
		lastLine := outputLines[len(outputLines)-2]
		fmt.Printf("Last line %v", lastLine)
		if strings.TrimSpace(lastLine) == "true" {
			return true, ""
		}
	}
	return false, outputStr
}
