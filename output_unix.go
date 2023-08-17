//go:build linux || darwin || openbsd || freebsd || netbsd
// +build linux darwin openbsd freebsd netbsd

package liner

import (
	"os"
	"syscall"
	"unsafe"
)

func (s *State) getColumns() bool {

	var ws winSize
	mc := syscall.Stdout
	if s.writer == os.Stderr {
		mc = syscall.Stderr
	}
	ok, _, _ := syscall.Syscall(syscall.SYS_IOCTL, uintptr(mc),
		syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws)))
	if int(ok) < 0 {
		return false
	}
	s.columns = int(ws.col)
	return true
}
