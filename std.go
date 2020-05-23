package timemock

import "time"

var std = New()

func Now() time.Time {
	return std.Now()
}

func Since(t time.Time) time.Duration {
	return std.Since(t)
}

func Until(t time.Time) time.Duration {
	return std.Until(t)
}

func Freeze(t time.Time) {
	std.Freeze(t)
}

func Travel(t time.Time) {
	std.Travel(t)
}

func Scale(s float64) {
	std.Scale(s)
}

func Return() {
	std.Return()
}
