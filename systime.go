package systime

import "time"

type systm interface {
	SetLocalTime(time.Time) error
}

var (
	st systm
)

func chooseSysTM(choose systm) {
	st = choose
}

// SetLocalTime is SetTimeOfDay() or SetLocalTime() systemcall.
func SetLocalTime(t time.Time) error {
	return st.SetLocalTime(t)
}
