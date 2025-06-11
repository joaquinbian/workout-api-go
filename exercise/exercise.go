package exercise

import (
	"fmt"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(i Item) {
	p.Inventory = append(p.Inventory, i)
	fmt.Printf("El jugador %s agarro un item de tipo %s: %s", p.Name, i.Type, i.Name)
}

func (p *Player) DropItem(name string) {
	var itemDropped Item
	for idx, item := range p.Inventory {
		if item.Name == name {
			itemDropped = p.Inventory[idx]
			if idx == len(p.Inventory)-1 {
				p.Inventory = p.Inventory[:idx]
			} else {
				p.Inventory = append(p.Inventory[:idx], p.Inventory[:idx+1]...)
			}

			return
		}
	}
	fmt.Printf("El jugador %s dejo un item de tipo %s: %s\n", p.Name, itemDropped.Type, itemDropped.Name)

}

func (p Player) UseItem(name string) {

	if p.hasItem(name) {
		fmt.Printf("El jugador %s decidio usar %s!\n", p.Name, name)
		if name == "poison" {
			p.DropItem(name)
		}
	} else {
		fmt.Println("No tienes el item ", name)
	}
}

func (p Player) hasItem(name string) bool {
	found := false
	for _, item := range p.Inventory {
		if item.Name == name {
			found = true
			break
		}

	}
	return found
}
