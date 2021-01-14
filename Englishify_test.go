package englishify

import "testing"

func Test(t *testing.T) {
	check("1", "1st", t)
	check(1, "1st", t)
	check("232", "232nd", t)
	check(934855, "934855th", t)
	check(23, "23rd", t)
}

func check(in interface{}, expected string, t *testing.T) {
	input := FmtNumber(in)
	if input != expected {
		t.Error("Check doesn't match: ", input, "; ", expected)
	}
}
