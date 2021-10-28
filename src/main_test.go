package owoifyx

import "testing"

func TestOwoifxy(t *testing.T) {
	teststruc := []struct {
		in  string
		out string
	}{
		{"hello world", "hewwo wowwd"},
		{"Nane", "Nyanye"},
		{"yesoveno", "yesuvnyo"},
		{"Hello Worls", "Hewwo Wowws"},
		{"nd", "ndo"},
		{"nd ", "ndo"},
	}

	for _, x := range teststruc {
		if out := Owoifxy(x.in); out != x.out {
			t.Errorf("Have gotten %v,wanted %v for %v,", out, x.out, x.in)
		}
	}
}
