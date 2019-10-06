package ld45

func possibleAll(fs ...func(*player) bool) func(*player) bool {
	return func(p *player) bool {
		for i := range fs {
			if !fs[i](p) {
				return false
			}
		}
		return true
	}
}

func possibleSome(fs ...func(*player) bool) func(*player) bool {
	return func(p *player) bool {
		for i := range fs {
			if fs[i](p) {
				return true
			}
		}
		return false
	}
}

func possibleNot(f func(*player) bool) func(*player) bool {
	return func(p *player) bool {
		return !f(p)
	}
}

func possibleAlways(_ *player) bool {
	return true
}

func possibleNever(_ *player) bool {
	return false
}

func minimalInventory(id itemID, minAmount int) func(*player) bool {
	return func(p *player) bool {
		return p.inventory.has(id, minAmount)
	}
}

func minimalToolQuality(id toolID, minQuality toolQuality) func(*player) bool {
	return func(p *player) bool {
		return p.equipment[id] >= minQuality
	}
}
