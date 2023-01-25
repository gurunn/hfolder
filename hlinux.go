//go:build !windows
// +build !windows

package hfolder

import "os"

func HideFile(foldername string) error {

	if err := os.MkdirAll(foldername, 0755); err != nil {
		return err
	}

	if err := os.Rename(foldername, "."+foldername); err != nil {
		return err
	}

	return nil
}
