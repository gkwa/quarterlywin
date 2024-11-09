package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRedisCommand(t *testing.T) {
	expectedURL := "redis://localhost:6379"
	os.Setenv("REDIS_URL", expectedURL)
	defer os.Unsetenv("REDIS_URL")

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd := rootCmd
	cmd.SetArgs([]string{"redis"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	expectedOutputs := []string{
		"Our redis url is: " + expectedURL,
		"the first character of our variable is: r",
		"the last character is: 9",
		"the value length is: " + fmt.Sprintf("%d", len(expectedURL)),
	}

	for _, expected := range expectedOutputs {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but got: %s", expected, output)
		}
	}
}
