//go:build !windows
// +build !windows

package sip

import "syscall"

func reusePortControl(network, address string, c syscall.RawConn) error {
	var opErr error
	if err := c.Control(func(fd uintptr) {
		opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, 0xf, 1)
	}); err != nil {
		return err
	}
	return opErr
}
