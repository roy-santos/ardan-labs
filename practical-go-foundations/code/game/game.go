package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var i1 Item // empty struct
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2} // initialized struct w/o using attribute name, values initialized according to struct order
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{ // initialized struct using parameter name, can init in any order
		Y: 10,
		X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 5520))

	i3.Move(100, 200)
	fmt.Printf("i3: %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}

	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)           // parameters in the embeded type (item) are available at top level of the player struct
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X) // can also do this but its not common practice unless there are multiple parameters with same name and you need to be explicit
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1, // when using interface, you need to be more explicit about passing the reference unlike in regular move fxn (not using an interface)
		&p1,
		&i2,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)

	// time.Time implement json.Marshaler interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
}

/*
// Implement sortByDistance(players []Player, x, y int)
func sortByDistance(players []Player, x, y int) {

}
*/

// Implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

/* Exercise
- Add a "Keys" field to Player which is a slice of Key
- Add a "FoundKey(k Key) error" method to player which will add k to Key if it's not there
	- Err if k is not one of the known keys
*/

// Go's version of "enum"
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal, not exported
)

type Key byte

// Rule of thumb: Accept interfaces, return types

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k > invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	// if !containsKey(p.Keys, k) {
	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

type Player struct {
	Name string
	Item // Embed Item (this is not same as extending)
	Keys []Key
}

// i is called the receiver
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item {}
// func NewItem(x, y int) *Item {}
// func NewItem(x, y int) {Item, error) {}
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds of min 0/0 and max %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// The Go compiler does "escape analysis" and will allocate i on the heap (outlives the fxn)
	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
type Item struct {
	X int
	Y int
}
