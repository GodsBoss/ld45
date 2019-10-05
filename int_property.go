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
	if prop.current < prop.minimum {
		prop.current = prop.minimum
	}
}

func (prop *intProperty) Inc(amount int) {
	prop.current += amount
	if prop.current > prop.maximum {
		prop.current = prop.maximum
	}
}
