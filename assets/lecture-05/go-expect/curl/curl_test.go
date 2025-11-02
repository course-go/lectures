package curl_test

import (
	"os"
	"os/exec"
	"testing"

	expect "github.com/Netflix/go-expect"
)

// START OMIT

func TestCURLGithub(t *testing.T) {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		t.Fatalf("failed creating pseudo-console: %v", err)
	}
	defer console.Close()
	command := exec.Command("curl", "-X", "HEAD", "-v", "github.com")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()
	err = command.Start()
	if err != nil {
		t.Fatalf("failed starting command: %v", err)
	}
	_, err = console.ExpectString("Location: https://github.com/")
	if err != nil {
		t.Fatalf("failed to read expected string: %v", err)
	}
	err = command.Wait()
	if err != nil {
		t.Fatalf("failed waiting for command to finish: %v", err)
	}
}

// END OMIT
