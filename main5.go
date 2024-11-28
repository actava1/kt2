package main

import "fmt"

// Подсистемы
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU starting...")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory loading...")
}

type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Hard drive reading...")
}

// Фасад
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) Start() {
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.Read()
}

// Клиентский код
func main() {
	computer := NewComputerFacade()
	computer.Start()
}
