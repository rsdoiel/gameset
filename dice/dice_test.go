package dice

import (
	"bytes"
	"testing"
)

func TestDice(t *testing.T) {
	for i := 0; i < 10000; i++ {
		j := Dice(4)
		if j == 0 {
			t.Errorf("zero returned by dice roll")
			t.FailNow()
		}
		if j < 0 {
			t.Errorf("negative one returned by dice roll for Dice(4)")
			t.FailNow()
		}
		if j > 4 {
			t.Errorf("Dice(4) returned large than 4")
			t.FailNow()
		}
	}
}

func TestParseRoll(t *testing.T) {
	testRolls := map[string]bool{
		"2d4":         true,
		"2q3":         false,
		"d4":          false,
		"1d4+5-10":    true,
		"1d4+1 2d4+2": false,
	}
	buf := []byte{}
	out := bytes.NewBuffer(buf)
	for roll, expected := range testRolls {
		_, _, _, err := ParseRoll(roll)
		if err != nil && expected == true {
			t.Errorf("expected success for ParseRoll(%q), got error %s", roll, err)
		}
		if err == nil && expected == false {
			t.Errorf("expected failure, for ParseRoll(%q), no error found", roll)
		}
		_, err = RollDice(out, []string{roll}, false)
		if err != nil && expected == true {
			t.Errorf("expected success for DiceRoll(%q), got error %s", roll, err)
		}
		if err == nil && expected == false {
			t.Errorf("expected failure for DiceRoll(%q), no error found", roll)
		}
	}
}
