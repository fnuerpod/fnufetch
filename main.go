// main.go
package main

import (
	"fmt"

	"flag"
	"strings"

	"github.com/fnuerpod/fnufetch/builders"
	"github.com/fnuerpod/fnufetch/logos"
	"github.com/fnuerpod/fnufetch/os_release"
)

func main() {
	os_struct := os_release.NewOSR()

	forceDistroFlag := flag.String("force_distro", os_struct.PrettyName, "Forces distribution name - overrides /etc/os-release PrettyName variable.")

	flag.Parse()

	if *forceDistroFlag != "" {
		// force flag set.
		os_struct.PrettyName = *forceDistroFlag
	}
	sys_logo := logos.GetLogo(os_struct.PrettyName)

	// make strings for injection
	inject_strings := []string{
		builders.BuildHeader(),
		builders.BuildOS(os_struct),
		builders.BuildHost(),
		builders.BuildKernel(),
		builders.BuildUptime(),
		builders.BuildPackages(os_struct),
		builders.BuildMemory(),
	}

	inject_strings = logos.ColorReplace(inject_strings)

	for line, val := range inject_strings {
		sys_logo[line] = sys_logo[line] + "\u001b[0m" + "\t  " + val
	}

	fmt.Println(strings.Join(sys_logo, "\n"))

}
