//go:build windows
// +build windows

package hfolder

import (
	"os"
	"syscall"
)

func HideFile(foldername string) error {

	if err := os.MkdirAll(foldername, 0755); err != nil {
		return err
	}

	filenameW, err := syscall.UTF16PtrFromString(foldername)
	if err != nil {
		return err
	}
	if err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN); err != nil {
		return err
	}

	return nil
}
