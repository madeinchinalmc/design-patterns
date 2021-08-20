package command

import "fmt"

// 基础命令接口
type command interface {
	execute()
}

// 请求者
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// 具体命令接口1
type onCommand struct {
	drive drive
}

func (b *onCommand) execute() {
	b.drive.on()
}

// 具体命令接口2
type offCommand struct {
	drive drive
}

func (b *offCommand) execute() {
	b.drive.off()
}

//接收者接口
type drive interface {
	on()
	off()
}

// 具体接收者
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// 客户端
func RunApplication() {
	myTv := &tv{}
	on := &onCommand{
		drive: myTv,
	}
	off := &offCommand{
		drive: myTv,
	}
	onButton := &button{
		command: on,
	}
	onButton.press()

	offButton := &button{
		command: off,
	}
	offButton.press()
}
