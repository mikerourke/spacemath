package spacemath

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	shell32           = syscall.NewLazyDLL("shell32.dll")
	shellExecuteWProc = shell32.NewProc("ShellExecuteW")
)

func open(filePath string) error {
	ret, _, _ := shellExecuteWProc.Call(
		0,
		stringToUintptr("open"),
		stringToUintptr(filePath),
		stringToUintptr(""),
		stringToUintptr(""),
		uintptr(syscall.SW_SHOWDEFAULT))

	// If the function succeeds, it returns a value greater than 32. If the
	// function fails, it returns an error value that indicates the cause of
	// the failure.
	if ret > 32 {
		return nil
	}

	return fmt.Errorf("could not open file: %s", filePath)
}

func stringToUintptr(value string) uintptr {
	if len(value) == 0 {
		return 0
	}

	ptr, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return 0
	}

	if ptr == nil {
		return 0
	}

	return uintptr(unsafe.Pointer(ptr))
}
