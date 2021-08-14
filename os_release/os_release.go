// os_release.go
package os_release

import (
	"log"
	"os"
	"strings"
)

type OSRelease struct {
	Name             string
	PrettyName       string
	Id               string
	BuildId          string
	ANSIColor        string
	HomeURL          string
	DocumentationURL string
	SupportURL       string
	BugReportURL     string
	Logo             string
}

func (OSR *OSRelease) String() string {
	presend := []string{
		"Name >> ", OSR.Name, "\n",
		"Pretty Name >> ", OSR.PrettyName, "\n",
		"ID >> ", OSR.Id, "\n",
		"Build ID >> ", OSR.BuildId, "\n",
		"ANSI Color >> ", OSR.ANSIColor, "\n",
		"Home URL >> ", OSR.HomeURL, "\n",
		"Documentation URL >> ", OSR.DocumentationURL, "\n",
		"Support URL >> ", OSR.SupportURL, "\n",
		"Bug Report URL >> ", OSR.BugReportURL, "\n",
		"Logo >> ", OSR.Logo,
	}

	return strings.Join(presend, "")
}

func NewOSR() *OSRelease {
	osr, err := os.ReadFile("/etc/os-release")

	if err != nil {
		// probably doesn't exist...
		log.Fatal(err)
	}

	// convert the byte slice into a string.
	osr_string := string(osr)

	// split new os-release string by newline.
	// this gives us each config line.
	osr_split := strings.Split(osr_string, "\n")

	osr_struct := new(OSRelease)

	// traverse through new string slice to generate configuration var.
	for _, config_line := range osr_split {
		// split by the first = sign, this gives us a config key/value set.
		line_split := strings.SplitN(config_line, "=", 2)

		// check length of split line, if not 1 then not a valid config, should skip.
		if len(line_split) < 1 {
			continue
		}

		// check key so we know what we have to add to.
		if line_split[0] == "NAME" {
			// NAME, will equal the name of the release e.g. "Arch Linux".

			// make sure to strip quotations if they exist.
			osr_struct.Name = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "PRETTY_NAME" {
			// PRETTY_NAME, will equal to the pretty name of the release e.g. "Arch Linux".

			// make sure to strip quotations if they exist.
			osr_struct.PrettyName = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "ID" {
			// ID, will equal to the ID of the release e.g. "arch".
			osr_struct.Id = line_split[1]
		} else if line_split[0] == "BUILD_ID" {
			// BUILD_ID, will equal to the build ID of the release e.g. "rolling".
			osr_struct.BuildId = line_split[1]
		} else if line_split[0] == "ANSI_COLOR" {
			// ANSI_COLOR, will equal to the ANSI color of the release name e.g. "38;2;23;147;209".

			// make sure to strip quotations if they exist.
			osr_struct.ANSIColor = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "HOME_URL" {
			// HOME_URL, will equal to the home website URL of the release e.g. "https://archlinux.org".

			// make sure to strip quotations if they exist.
			osr_struct.HomeURL = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "DOCUMENTATION_URL" {
			// DOCUMENTATION_URL, will equal to the documentation website URL of the release e.g. "https://wiki.archlinux.org".

			// make sure to strip quotations if they exist.
			osr_struct.DocumentationURL = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "SUPPORT_URL" {
			// SUPPORT_URL, will equal to the support website URL of the release e.g. "https://bbs.archlinux.org".

			// make sure to strip quotations if they exist.
			osr_struct.SupportURL = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "BUG_REPORT_URL" {
			// BUG_REPORT_URL, will equal to the bug reporting website URL of the release e.g. "https://bugs.archlinux.org".

			// make sure to strip quotations if they exist.
			osr_struct.BugReportURL = strings.ReplaceAll(line_split[1], "\"", "")
		} else if line_split[0] == "LOGO" {
			// LOGO, returns the name of the logo(?) of the release e.g. "archlinux".
			osr_struct.Logo = line_split[1]
		}
	}

	return osr_struct
}
