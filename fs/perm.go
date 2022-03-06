// +build !windows

package fs

import (
	"fmt"
	"os"
	"syscall"
)

const rootEUID = uint32(0)

// IsWritable test if directory path is writable or not
func IsWritable(path string) (isWritable bool, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("not a directory")
	}

	// Check if the user bit is enabled in file permission
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, nil
	}

	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return false, err
	}

	euid := uint32(os.Geteuid())
	if euid == rootEUID {
		return true, nil
	}
	return euid == stat.Uid, nil
}
