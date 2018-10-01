package main

import (
	"log"
	"os/exec"

	volume "github.com/itchyny/volume-go"
)

func execCmd(cmd string) string {
	out, err := exec.Command("cmd", "/c", cmd).CombinedOutput()

	if err != nil {
		log.Fatalf("Command finished with error: %v", err)
		return "ERROR"
	}

	return string(out)
}

func execCmdSilent(cmd string) {
	err := exec.Command("cmd", "/c", "START", "/MIN", cmd).Start()

	if err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}
}

func main() {
	volume.SetVolume(100)

	execCmdSilent("narrator")
}
