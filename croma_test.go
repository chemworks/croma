package chengcroma

import (
	"fmt"
	"testing"
)

func TestCroma(t *testing.T) {
	c := &Croma{
		MolFrac: map[string]float64{
			"C1":  98,
			"C2":  2,
			"C3":  0,
			"iC4": 0,
			"nC4": 0,
			"iC5": 0,
			"nC5": 0,
			"nC6": 00,
			"nC7": 00,
			"nC8": 00,
			"N2":  00,
			"CO2": 0,
			"H2O": 0,
			"SH2": 0,
		},
	}
	ma := c.CromaNorm()
	fmt.Print(ma)
}

func BenchmarkCroma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := &Croma{
			MolFrac: map[string]float64{
				"C1":  98,
				"C2":  2,
				"C3":  0,
				"iC4": 0,
				"nC4": 0,
				"iC5": 0,
				"nC5": 0,
				"nC6": 00,
				"nC7": 00,
				"nC8": 00,
				"N2":  00,
				"CO2": 0,
				"H2O": 0,
				"SH2": 0,
			},
		}
		c.CromaNorm()
	}

}
