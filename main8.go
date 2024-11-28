package main

import "fmt"

// Команда
type Command interface {
	Execute()
}

// Получатель
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is On")
}

func (l *Light) Off() {
	fmt.Println("Light is Off")
}

// Конкретные команды
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Инвокер
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Клиентский код
func main() {
	light := &Light{}
	onCommand := &LightOnCommand{light: light}
	offCommand := &LightOffCommand{light: light}

	remote := &RemoteControl{}
	remote.SetCommand(onCommand)
	remote.PressButton()

	remote.SetCommand(offCommand)
	remote.PressButton()
}
