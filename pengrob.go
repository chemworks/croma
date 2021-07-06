package chengcroma

import (
	"fmt"
	"math"

	"github.com/chemworks/gocubicsolver"
)

const (
	R = 8.314 // Comment
)

// TODO Doc
func PengP(P, Pc, T, Tc, acent float64) float64 {
	// https://en.wikipedia.org/wiki/Equation_of_state#Peng%E2%80%93Robinson_equation_of_state
	// split terms in Ta Tb ...
	// we make
	// ax^3 + bx^2 + cx + d = 0
	A := helpA(P, Pc, T, Tc, acent)
	B := helpB(P, Pc, T, Tc)
	a := 1.0
	b := (1 - B)
	c := (A - 2*B - 3*B*B)
	d := (A*B - B*B - B*B*B)
	r, err := gocubicsolver.Solve(a, b, c, d)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(r)
	return r[0]
}

func helpa(Pc, Tc float64) float64 {
	return 0.45724 * (R * R) * (Tc * Tc) / Pc
}

func helpb(Pc, Tc float64) float64 {
	return 0.07780 * (R * Tc) / Pc
}

func alpha(Tr, k float64) float64 {
	T := math.Pow(Tr, 0.5)
	return math.Pow(1+k*(1-T), 2)
}

func kpar(acent float64) float64 {
	return 0.37464 + 1.54226*acent - 0.26992*acent*acent
}

func helpB(P, Pc, T, Tc float64) float64 {
	b := helpb(Pc, Tc)
	return (b * P) / (R * T)
}

func helpA(P, Pc, T, Tc, acent float64) float64 {
	Tr := T / Tc
	k := kpar(acent)
	alf := alpha(Tr, k)
	a := helpa(Pc, Tc)
	return (alf * a * P) / (R * R * T * T)

}
