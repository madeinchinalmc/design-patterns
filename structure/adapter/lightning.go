package adapter

import "fmt"

// 客户端对象

type client struct {
}

// 客户端函数执行

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

// 客户端接口

type computer interface {
	insertIntoLightningPort()
}

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct {
}

func (m *windows) insertIntoLightningPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// 系统适配器

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoLightningPort()
}

func RunApplication() {
	myClient := &client{}
	myClient.insertLightningConnectorIntoComputer(&mac{})

	myWin := &windows{}
	winAdap := &windowsAdapter{
		myWin,
	}
	myClient.insertLightningConnectorIntoComputer(winAdap)
}
