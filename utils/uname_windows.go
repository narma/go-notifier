package utils

//#include <windows.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// Return windows release build.
// WARNING: this method is deprecated and without manifest returns maximum 6.2.9200 (Windows 8)
// hint: for version comparsion you can use VerifyVersionInfo instead
// Links:
// https://social.msdn.microsoft.com/Forums/windowsdesktop/en-US/c471de52-611f-435d-ab44-56064e5fd7d5/windows-81-preview-getversionex-reports-629200?forum=windowssdk
// http://blogs.msdn.com/b/chuckw/archive/2013/09/10/manifest-madness.aspx
func OSRelease() string {
	var info C.struct__OSVERSIONINFOW
	info.dwOSVersionInfoSize = C.DWORD(unsafe.Sizeof(info))
	C.GetVersionExW(&info)
	return fmt.Sprintf("%d.%d.%d", info.dwMajorVersion, info.dwMinorVersion, info.dwBuildNumber)
}
