package main

import "fmt"

// Продукт
type House struct {
	Windows string
	Doors   string
	Roof    string
}

// Строитель
type HouseBuilder struct {
	house *House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{house: &House{}}
}

func (b *HouseBuilder) SetWindows(windows string) *HouseBuilder {
	b.house.Windows = windows
	return b
}

func (b *HouseBuilder) SetDoors(doors string) *HouseBuilder {
	b.house.Doors = doors
	return b
}

func (b *HouseBuilder) SetRoof(roof string) *HouseBuilder {
	b.house.Roof = roof
	return b
}

func (b *HouseBuilder) Build() *House {
	return b.house
}

// Клиентский код
func main() {
	builder := NewHouseBuilder()
	house := builder.SetWindows("Double-glazed").
		SetDoors("Wooden").
		SetRoof("Tile").
		Build()

	fmt.Printf("House: %+v\n", house)
}
