package ld45

type intProperty struct {
	current int
	minimum int
	maximum int
}

func (prop *intProperty) IsMinimum() bool {
	return prop.current == prop.minimum
}

func (prop *intProperty) IsMaximum() bool {
	return prop.current == prop.maximum
}

func (prop *intProperty) Dec(amount int) {
	prop.current -= amount
	prop.restrict()
}

func (prop *intProperty) Inc(amount int) {
	prop.current += amount
	prop.restrict()
}

func (prop *intProperty) restrict() {
	if prop.current < prop.minimum {
		prop.current = prop.minimum
	}
	if prop.current > prop.maximum {
		prop.current = prop.maximum
	}
}

func (prop *intProperty) Set(value int) {
	prop.current = value
	prop.restrict()
}
