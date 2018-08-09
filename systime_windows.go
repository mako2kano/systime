// Copyright 2018 Makoto Tokano.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package systime

import (
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modkernel32       = windows.NewLazySystemDLL("kernel32.dll")
	procSetLocalTime  = modkernel32.NewProc("SetLocalTime")
	procSetSystemTime = modkernel32.NewProc("SetSystemTime")
)

type windowsSYSTEMTIME struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

type windowsSysTime struct{}

// SetLocalTime (timezone)
// BOOL SetLocalTime(CONST SYSTEMTIME *lpSystemTime)
func (windowsSysTime) SetLocalTime(t *time.Time) (err error) {

	st := windowsSYSTEMTIME{
		wYear:         uint16(t.Year()),
		wMonth:        uint16(t.Month()),
		wDay:          uint16(t.Day()),
		wHour:         uint16(t.Hour()),
		wMinute:       uint16(t.Minute()),
		wSecond:       uint16(t.Second()),
		wMilliseconds: uint16(t.Nanosecond() / 10e6),
	}

	return syscallSetLocalTime(&st)
}

func syscallSetLocalTime(st *windowsSYSTEMTIME) error {

	r0, _, e1 := syscall.Syscall(procSetLocalTime.Addr(), 1, uintptr(unsafe.Pointer(st)), 0, 0)

	// success
	if r0 != 0 {
		return nil
	}

	// fail
	return e1
}

// SetSystemTime (UTC)
// BOOL SetSystemTime(CONST SYSTEMTIME *lpSystemTime)
func (windowsSysTime) SetSystemTime(t *time.Time) (err error) {

	st := windowsSYSTEMTIME{
		wYear:         uint16(t.UTC().Year()),
		wMonth:        uint16(t.UTC().Month()),
		wDay:          uint16(t.UTC().Day()),
		wHour:         uint16(t.UTC().Hour()),
		wMinute:       uint16(t.UTC().Minute()),
		wSecond:       uint16(t.UTC().Second()),
		wMilliseconds: uint16(t.UTC().Nanosecond() / 10e6),
	}

	return syscallSetLocalTime(&st)
}

func syscallSetSystemTime(st *windowsSYSTEMTIME) error {

	r0, _, e1 := syscall.Syscall(procSetSystemTime.Addr(), 1, uintptr(unsafe.Pointer(st)), 0, 0)

	// success
	if r0 != 0 {
		return nil
	}

	// fail
	return e1
}

func init() {
	chooseSysTM(windowsSysTime{})
}
