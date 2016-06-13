package droptest

import "testing"

func TestGreet(t *testing.T) {
	actual := Greet("Mat")
	expected := "Hello Mat"
	if actual != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, actual)
	}
}
