package exercise

import "fmt"

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
	fmt.Printf("Player %s picked up an %s item: %s", p.Name, i.Type, i.Name)
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
	fmt.Printf("Player %s dropped an %s item: %s\n", p.Name, itemDropped.Type, itemDropped.Name)

}

func (p Player) UseItem(name string) {

}
