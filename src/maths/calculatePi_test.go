package maths

import (
	"math/big"
	"testing"
)

func TestPi(t *testing.T) {
	// t.Error("TestPI")
	type args struct {
		accuracy uint
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{"TestPI", args{60}, new(big.Float).SetPrec(60)},
		{"TestPI", args{600}, new(big.Float).SetPrec(600)},
	}
	wanted, status := tests[0].want.SetString("3.141592653589793258")
	if status {
		tests[0].want = wanted
	}

	wanted, status = tests[1].want.SetString("3.141592653589793256960399361738762404019183156248573243493179283571046450248913467118511784317615354282017929416292809050813937875283435610586313363548602436768047706489838924381929")
	if status {
		tests[1].want = wanted
	}
	// fmt.Sscan("3.141592653589793258", tests[0].want)
	t.Log("Test PI")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pi(tt.args.accuracy)
			t.Log("Expected", tt.want, "Result:", got)
			if got.Cmp(wanted) != 0 {
				t.Errorf("Pi() = %v, want %v", got, tt.want)
			}
		})
	}
}
