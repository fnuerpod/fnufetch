package builders

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/fnuerpod/fnufetch/package_mgr"

	"github.com/capnm/sysinfo"
	"github.com/jaypipes/ghw"

	"github.com/fnuerpod/fnufetch/kernel"
	"github.com/fnuerpod/fnufetch/os_release"

	"math"
	"time"
)

func humanizeDuration(duration time.Duration) string {
	if duration.Seconds() < 60.0 {
		return fmt.Sprintf("%d seconds", int64(duration.Seconds()))
	}
	if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%dm %ds", int64(duration.Minutes()), int64(remainingSeconds))
	}
	if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%dh %dm %ds",
			int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds))
	}
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)
	return fmt.Sprintf("%dd %dh %dm %ds",
		int64(duration.Hours()/24), int64(remainingHours),
		int64(remainingMinutes), int64(remainingSeconds))
}

func BuildHeader() string {
	// get current user and hostname
	hostname, err := os.Hostname()

	if err != nil {
		hostname = "UNKNOWN"
	}

	username := ""

	user, err := user.Current()

	if err != nil {
		username = "UNKNOWN"
	} else {
		username = user.Username
	}

	hostname_string := "\u001b[35;1m" + username + "\u001b[0m@" + "\u001b[35;1m" + hostname

	return hostname_string
}

func BuildOS(os_struct *os_release.OSRelease) string {
	return "${c4}os ${c0}    " + os_struct.PrettyName
}

func BuildHost() string {
	// get product name
	product_name := "UNKNOWN"
	product, err := ghw.Product(ghw.WithDisableWarnings())
	if err == nil {
		product_name = product.Name
		//fmt.Printf("Error getting product info: %v", err)
	}

	host_string := "${c4}host ${c0}  " + product_name

	return host_string
}

func BuildKernel() string {
	return "${c4}kernel ${c0}" + kernel.GetKernelVersion()
}

func BuildUptime() string {
	return "${c4}uptime ${c0}" + humanizeDuration(sysinfo.Get().Uptime)
}

func BuildPackages(os_struct *os_release.OSRelease) string {
	return "${c4}pkgs ${c0}  " + strconv.Itoa(package_mgr.GetPackageCount(os_struct))
}

func BuildMemory() string {
	inf := sysinfo.Get()

	return "${c4}memory ${c0}" + strconv.FormatUint((inf.TotalRam-inf.FreeRam)/1000, 10) + "M / " + strconv.FormatUint((inf.TotalRam)/1000, 10) + "M"
}
