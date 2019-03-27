// +build linux,arm64

package panicwrap

import (
	//"syscall"
)

func dup2(oldfd, newfd int) error {
	return errors.New("dup3")
}
