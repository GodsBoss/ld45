package ld45

type playResult struct {
	victory bool
}

func (result *playResult) IsVictory() bool {
	return result.victory
}

func (result *playResult) SetVictory() {
	result.victory = true
}

func (result *playResult) SetDefeat() {
	result.victory = false
}
