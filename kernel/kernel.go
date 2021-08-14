package kernel

import (
	"log"
	"os/exec"
	"strings"
)

func GetKernelVersion() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(string(out), "\n")
}
