//go:build !windows
// +build !windows

package logs

import "os"

func CreateFileAgainstSystemTunneling(filename string) error {
	// not implement
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return f.Close()
}
