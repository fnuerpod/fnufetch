package os_release

import (
	"fmt"
	"testing"
)

var OSReleaseStruct *OSRelease

func Test_CreateOSR(t *testing.T) {
	OSReleaseStruct = NewOSR()
}

func Test_OSRContent(t *testing.T) {
	fmt.Println(OSReleaseStruct.String())
}
