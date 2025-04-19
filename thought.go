package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GenerateThought(prompt string) (string, error) {
	cmd := exec.Command("ollama", "run", "qwen2.5-coder:14b", prompt)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("ollama error: %v\nstderr: %s", err, stderr.String())
	}

	return out.String(), nil
}
