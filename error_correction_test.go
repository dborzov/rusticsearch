package main

import "testing"

func TestRemoveDashesEdits(t *testing.T) {
	in := "Rolls-Royce"
	out := []string{"RollsRoyce", "Rolls Royce"}

	x := RemoveDashesEdits(in)

	if len(out) != len(x) {
		t.Errorf("RemoveDashesEdits(%#v) = %#v != %#v", in, x, out)
		return
	}

	for i := range out {
		if out[i] != x[i] {
			t.Errorf("RemoveDashesEdits(%#v) has %#v, expected %#v", in, x[i], out[i])
			return
		}
	}

}

func TestMultipleRemoveDashesEdits(t *testing.T) {
	in := "Rolls-Royce Counter-Intuitive"
	out := []string{
		"RollsRoyce Counter-Intuitive",
		"Rolls Royce Counter-Intuitive",
		"RollsRoyce CounterIntuitive",
		"Rolls Royce Counter Intuitive",
	}

	x := RemoveDashesEdits(in)

	if len(out) != len(x) {
		t.Errorf("RemoveDashesEdits(%#v) = %#v != %#v", in, x, out)
		return
	}

	for i := range out {
		if out[i] != x[i] {
			t.Errorf("RemoveDashesEdits(%#v) has %#v, expected %#v", in, x[i], out[i])
			return
		}
	}

}

func TestExtraLetterEdits(t *testing.T) {
	in := "ABBA"
	out := []string{"ABA", "ABA"}

	x := ExtraLetterEdits(in)

	if len(out) != len(x) {
		t.Errorf("ExtraLetterEdits(%#v) = %#v != %#v", in, x, out)
		return
	}

	for i := range out {
		if out[i] != x[i] {
			t.Errorf("ExtraLetterEdits(%#v) has %#v, expected %#v", in, x[i], out[i])
			return
		}
	}
}

func TestNullExtraLetterEdits(t *testing.T) {
	in := ""
	out := []string{}

	x := ExtraLetterEdits(in)

	if len(out) != len(x) {
		t.Errorf("ExtraLetterEdits(%#v) = %#v != %#v", in, x, out)
		return
	}

	for i := range out {
		if out[i] != x[i] {
			t.Errorf("ExtraLetterEdits(%#v) has %#v, expected %#v", in, x[i], out[i])
			return
		}
	}
}
