// +build !windows

package term

import (
	"sync"

	"golang.org/x/sys/unix"
)

var (
	saveTermios     *unix.Termios
	saveTermiosFD   int
	saveTermiosOnce sync.Once
)

func getOriginalTermios(fd int) (*unix.Termios, error) {
	var err error
	saveTermiosOnce.Do(func() {
		saveTermiosFD = fd
		saveTermios, err = unix.IoctlGetTermios(fd, unix.TCGETS)
	})
	return saveTermios, err
}

// Restore terminal's mode.
func Restore() error {
	o, err := getOriginalTermios(saveTermiosFD)
	if err != nil {
		return err
	}
	return unix.IoctlSetTermios(saveTermiosFD, unix.TCSETS, o)
}
