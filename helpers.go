package chengcroma

import (
	"errors"
	"fmt"
)

const (
	ptokpa = 1.033
)

// Return propertie calculated
func MeanProp(mf, prop map[string]float64) (float64, error) {
	p := 0.0 // property to be calculated
	keys := MapKeys(mf)
	for _, cmp := range keys { // using cmp has compound
		// Check if not component in prop
		if !MapHas(prop, cmp) {
			return 0, errors.New(fmt.Sprintf("Struct %+v not has %s", prop, cmp))
		}
		p = p + prop[cmp]*mf[cmp]
	}
	return p, nil

}

// Return and slice of string
// m: map with string has keys
func MapKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Check if a Map Has a string as key
func MapHas(m map[string]float64, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}
	return false
}

// Calculates the reduced parameter
func ReducedParameter(abs, cri float64) (float64, error) {
	if abs == 0 {
		return 0, errors.New(fmt.Sprintf("Abs prop is 0"))
	}
	if cri == 0 {

		return 0, errors.New(fmt.Sprintf("Crit prop is 0"))
	}

	return abs / cri, nil
}

// Return the pressure in kpa
func Pkgfcm2gToPkpa(p float64) float64 {
	return (p/ptokpa + 1) * 100

}
