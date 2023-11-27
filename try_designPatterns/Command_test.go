package try_designPatterns

import (
	"fmt"
	"testing"
)

type command interface {
	Execute()
}

type Light struct {
}

func (l *Light) on() {
	fmt.Println("点燃一盏灯")
}

type LightOnCommand struct {
	light *Light
}

func (l LightOnCommand) Execute() {
	l.light.on()
}

type RemoteControl struct {
	command command
}

func (r RemoteControl) pressButton() {
	r.command.Execute()
}

func TestCommand(t *testing.T) {
	control := RemoteControl{
		command: LightOnCommand{light: &Light{}},
	}
	control.pressButton()
}
