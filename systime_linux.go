// Copyright 2018 Makoto Tokano.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package systime

import (
	"syscall"
	"time"
)

type linuxSysTime struct{}

func (linuxSysTime) SetLocalTime(t *time.Time) error {

	tv := syscall.NsecToTimeval(t.UnixNano())
	return syscall.Settimeofday(&tv)
}

func init() {
	chooseSysTM(linuxSysTime{})
}
