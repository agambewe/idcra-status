package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mgutz/ansi"
)

func printAlert(color, service, status string) string {
	return fmt.Sprintf("%s%s \t ----=> %s", color, service, status)
}

func main() {
	reset := ansi.ColorCode("reset")
	services := []string{
		"nginx",
		"idcra-api",
		"idcra-reminder",
	}

	for _, service := range services {

		syscmd := exec.Command("systemctl", "status", service)
		grepcmd := exec.Command("grep", "Active:")

		grepcmd.Stdin, _ = syscmd.StdoutPipe()
		var byteB bytes.Buffer
		grepcmd.Stdout = &byteB

		_ = grepcmd.Start()
		_ = syscmd.Run()
		_ = grepcmd.Wait()

		status := fmt.Sprintf("%s", &byteB)

		if strings.Contains(status, "active (running)") {
			color := ansi.ColorCode("green+bh:black")

			fmt.Println(printAlert(color, service, status), reset)
		} else if strings.Contains(status, "active (mounted)") {
			color := ansi.ColorCode("yellow:black")

			fmt.Println(printAlert(color, service, status), reset)
		} else {
			color := ansi.ColorCode("red+bh:black")

			fmt.Println(printAlert(color, service, status), reset)
		}
	}
}
