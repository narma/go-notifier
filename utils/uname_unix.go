// +build linux darwin dragonfly freebsd openbsd solaris

package utils

// #include <sys/utsname.h>
import "C"

func OSRelease() string {
	var utsname C.struct_utsname
	C.uname(&utsname)
	return C.GoString(&utsname.release[0])
}
