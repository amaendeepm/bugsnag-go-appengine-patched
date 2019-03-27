// +build darwin dragonfly freebsd linux,!arm64 netbsd openbsd

package panicwrap

import (
	//"syscall"
	"errors"
)

func dup2(oldfd, newfd int) error {
	return errors.New("dup2")
}
