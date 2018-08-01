// Copyright 2018 Makoto Tokano.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package systime

import "time"

type systm interface {
	SetLocalTime(*time.Time) error
}

var (
	st systm
)

func chooseSysTM(choose systm) {
	st = choose
}

// SetLocalTime is SetTimeOfDay() or SetLocalTime() systemcall.
func SetLocalTime(t *time.Time) error {
	return st.SetLocalTime(t)
}
