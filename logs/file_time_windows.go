//go:build windows
// +build windows

package logs

import (
	"syscall"
)

func CreateFileAgainstSystemTunneling(filename string) error {
	pathp, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}
	h, err := syscall.CreateFile(pathp,
		syscall.FILE_WRITE_ATTRIBUTES, syscall.FILE_SHARE_WRITE, nil,
		syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(h)
	var c syscall.Filetime
	syscall.GetSystemTimeAsFileTime(&c)
	return syscall.SetFileTime(h, &c, nil, nil)
}
