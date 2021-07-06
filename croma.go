package chengcroma

// Builder for croma
type Builder struct {
	c Croma
}

func (b *Builder) Build() Croma {
	return b.c
}

// Usage c :=&Builder{}
// croma := b.
// 			  Gas("NG").
// 			  Build()
//

func (b *Builder) Gas(s string) *Builder {
	if s == "NG" {
		b.c.MolFrac["C1"] = 0.95
		b.c.MolFrac["C2"] = 0.03
		b.c.MolFrac["C3"] = 0.01
		b.c.MolFrac["N2"] = 0.01
		b.c.MolFrac["CO2"] = 0.01
	}

	if s == "Air" {
		b.c.MolFrac["N2"] = 0.79
		b.c.MolFrac["O2"] = 0.21
	}
	return b
}

type Croma struct {
	// Mol Frac User input
	MolFrac map[string]float64
	// Mol Frac normalized
	MolFracNorm map[string]float64
	// Pressure in kpa
	P float64
	// Temperature in K
	T float64
	// Properties Calculated with norm Values this must be populated with Calc
	// HHV kj/kmol
	HHV float64
	// LHV kj/kmol
	LHV float64
	// Mol W kg/kmol
	MW float64
	// Comp Factor
	Z float64
	// Critical Temp in K
	TC float64
	// Reduced parameter
	TR float64
	// Critical Pressure in kpa
	PC float64
	// Reduced parameter
	PR float64
	// Acentricity factor
	AcFacT float64
	// Boiling point temp in K
	TB float64
	// Freezing point Temp in K
	TF float64
	// Latent heat vap kj/Kg
	Lambda float64
	// Cp kj/kgK
	Cp float64
	// Cv kj/KgK
	Cv float64
	// K = Cp/Cv
	K float64
	// High Exp Limit %
	HEL float64
	// Lower Exp Limit %
	LEL float64
	// Next sections is for internal flags
	// norm if has already adjusted
	norm bool
}

// method
func (c *Croma) Calc(P, T float64) error {
	c.setP(P)
	c.setP(T)
	err := c.setK()
	if err != nil {
		return err
	}

	err = c.setMW()
	if err != nil {
		return err
	}

	err = c.setK()
	if err != nil {
		return err
	}
	err = c.setPc()
	if err != nil {
		return err
	}
	err = c.setTc()
	if err != nil {
		return err
	}

	return nil
}

// method to set HHV
func (c *Croma) setHHV() error {
	val, err := c.GetProp(HHV)
	if err != nil {
		return err
	}
	c.HHV = val
	return nil
}

// method to set LHV
func (c *Croma) setLHV() error {
	val, err := c.GetProp(LHV)
	if err != nil {
		return err
	}
	c.LHV = val
	return nil
}

// method to set Cp
func (c *Croma) setCp() error {
	val, err := c.GetProp(Cp)
	if err != nil {
		return err
	}
	c.Cp = val
	return nil
}

// method to set Cv
func (c *Croma) setCv() error {
	val, err := c.GetProp(Cv)
	if err != nil {
		return err
	}
	c.Cv = val
	return nil
}

// method to set K
func (c *Croma) setK() error {
	err := c.setCp()
	if err != nil {
		return err
	}
	err = c.setCv()
	if err != nil {
		return err
	}
	c.K = c.Cp / c.Cv
	return nil
}

// method to set Z
func (c *Croma) setZ() error {
	val, err := c.GetProp(Z)
	if err != nil {
		return err
	}
	c.Z = val
	return nil
}

// Return the mean MW according to GP S23 p23
func (c *Croma) setMW() error {
	val, err := c.GetProp(MW)
	if err != nil {
		return err
	}
	c.MW = val
	return nil
}

// Return and set the Pc
func (c *Croma) setPc() error {
	val, err := c.GetProp(PC)
	if err != nil {
		return err
	}
	c.PC = val
	return nil
}

// Return and set the Tc
func (c *Croma) setTc() error {
	val, err := c.GetProp(TC)
	if err != nil {
		return err
	}
	c.TC = val
	return nil
}

// Set a new value of pressure
// p in kpa
func (c *Croma) setP(p float64) {
	c.PR = 0.0
	c.P = p
}

// Set
func (c *Croma) setPr() error {
	val, err := ReducedParameter(c.P, c.PC)
	if err != nil {
		return err
	}
	c.PR = val
	return nil
}

// Set a new value of temp
// t in K
func (c *Croma) setT(t float64) {
	c.PR = 0.0
	c.P = t
}

// Set
func (c *Croma) setTr() error {
	val, err := ReducedParameter(c.T, c.TC)
	if err != nil {
		return err
	}
	c.TR = val
	return nil
}

// Normalize the fractions or percenrts
func (c *Croma) Normalize() {
	sum := 0.0
	for _, cmp := range comp {

		sum = sum + c.MolFrac[cmp]
	}
	if sum != 0 {
		c.norm = true
		for _, cmp := range comp {
			c.MolFracNorm[cmp] = c.MolFrac[cmp] / sum
		}
	}
}

// Get some prop from mol fractions
func (c *Croma) GetProp(prop map[string]float64) (float64, error) {
	if c.norm != true {
		c.Normalize()
	}
	val, err := MeanProp(c.MolFracNorm, prop)
	return val, err
}
