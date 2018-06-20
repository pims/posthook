package posthook

import "time"

// String is a convenience method for string ptrs
func String(s string) *string {
	return &s
}

// Time is a convenience method for time ptrs
func Time(t time.Time) *time.Time {
	return &t
}

// Int is a convenience method for int ptrs
func Int(i int) *int {
	return &i
}
