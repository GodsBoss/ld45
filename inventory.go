package ld45

import "sort"

type inventory struct {
	possessions map[itemID]int
	objects     []Object
}

func (inv *inventory) add(id itemID, amount int) {
	inv.possessions[id] += amount
	inv.generateObjectList()
}

func (inv *inventory) get(id itemID) int {
	return inv.possessions[id]
}

func (inv *inventory) has(id itemID, amount int) bool {
	return inv.possessions[id] >= amount
}

func (inv *inventory) generateObjectList() {
	inv.objects = make([]Object, 0)

	itemIDs := make([]string, 0)
	for itemID := range inv.possessions {
		if inv.possessions[itemID] == 0 {
			continue
		}
		itemIDs = append(itemIDs, string(itemID))
	}
	sort.Strings(itemIDs)
	row := 0
	for _, id := range itemIDs {
		itemIconCount := inv.possessions[itemID(id)]
		if itemIconCount > maxItemIconCount {
			itemIconCount = maxItemIconCount
			inv.objects = append(
				inv.objects,
				Object{
					X:   400 - (itemIconCount+1)*itemIconWidth,
					Y:   2 + row*itemIconHeight,
					Key: "more_items_marker",
				},
			)
		}
		for i := 0; i < itemIconCount; i++ {
			inv.objects = append(
				inv.objects,
				Object{
					X:   400 - (i+1)*itemIconWidth,
					Y:   2 + row*itemIconHeight,
					Key: string(id),
				},
			)
		}
		row++
	}
}

const itemIconWidth = 8
const itemIconHeight = 6
const maxItemIconCount = 10

func (inv *inventory) Objects() []Object {
	return inv.objects
}
