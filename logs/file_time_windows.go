//go:build windows
// +build windows

package logs

import (
	"fmt"
	"syscall"
)

func CreateFileAgainstSystemTunneling(filename string) error {
	path, err := syscall.FullPath(filename)
	if err != nil {
		return err
	}
	pPath, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return fmt.Errorf("UTF16PtrFromString: %w", err)
	}
	h, err := syscall.CreateFile(pPath,
		syscall.FILE_WRITE_ATTRIBUTES, syscall.FILE_SHARE_WRITE, nil,
		syscall.CREATE_NEW, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return fmt.Errorf("CreateFile: %w", err)
	}
	defer syscall.Close(h)
	var c syscall.Filetime
	syscall.GetSystemTimeAsFileTime(&c)
	if err:= syscall.SetFileTime(h, &c, nil, nil); err != nil {
		return fmt.Errorf("SetFileTime: %w", err)
	}
	return nil
}
