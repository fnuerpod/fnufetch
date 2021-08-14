package package_mgr

import (
	"log"
	"os/exec"
	"strings"

	"github.com/fnuerpod/fnufetch/os_release"
)

func GetPackageCount(osr *os_release.OSRelease) int {
	if osr.Name == "Arch Linux" {
		out, err := exec.Command("pacman", "-Qq").Output()
		if err != nil {
			log.Fatal(err)
		}

		string_slice := strings.Split(string(out), "\n")

		// -1 removes the POSIX standard newline at end of file.
		return len(string_slice) - 1
	} else if osr.Name == "Ubuntu" {
		out, err := exec.Command("dpkg", "-l").Output()
		if err != nil {
			log.Fatal(err)
		}

		string_slice := strings.Split(string(out), "\n")

		// -1 removes the POSIX standard newline at end of file.
		return len(string_slice) - 1
	}

	// no conditions met... return 0.
	return 0

}
