package greetutil

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"male", "Mr"},
		{"female", "Mrs"},
		{"", "Mrs"},
	}
	for _, c := range cases {
		got := Greet(c.in)
		if got != c.want {
			t.Errorf("Greet(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
