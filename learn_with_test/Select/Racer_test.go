package Select

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.ul"
	want := fastURL
	got := Racer(slowURL, fastURL)
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
