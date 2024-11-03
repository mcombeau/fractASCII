package main

import (
	"os"
	"syscall"
	"unsafe"
)

// Enable raw TTY mode where key presses are not interpreted by the terminal.
// Instead, they are passed to the program directly.
func enableRawModeTTY() {
	fd := int(os.Stdin.Fd())
	// Save original terminal state
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termiosBackup)))
	// Create new rawTermios mode termios
	rawTermios := termiosBackup
	// Disable echo and canonical modes, and signal generation (ISIG)
	rawTermios.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG
	// Disable CR-to-NL, NL-to-CR, and IGNCR to prevent newline issues
	rawTermios.Iflag &^= syscall.ICRNL | syscall.INLCR | syscall.IGNCR
	// Apply new raw mode termios using TCSETS
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&rawTermios)))
}

// Restore the original TTY state
func disableRawModeTTY() {
	fd := int(os.Stdin.Fd())
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&termiosBackup)))
}
