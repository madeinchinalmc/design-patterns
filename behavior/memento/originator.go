package memento

import "fmt"

// 原发器
type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{
		state: o.state,
	}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(s string) {
	o.state = s
}

func (o *originator) getState() string {
	return o.state
}

// 备忘录
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

//负责人
type caretaker struct {
	mementos []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementos = append(c.mementos, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementos[index]
}

// 客户端代码

func RunApplication() {
	myCaretaker := &caretaker{
		mementos: make([]*memento, 0),
	}
	myOriginator := &originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", myOriginator.getState())
	myCaretaker.addMemento(myOriginator.createMemento())

	myOriginator.setState("B")
	fmt.Printf("Originator Current State: %s\n", myOriginator.getState())
	myCaretaker.addMemento(myOriginator.createMemento())

	myOriginator.setState("C")
	fmt.Printf("Originator Current State: %s\n", myOriginator.getState())
	myCaretaker.addMemento(myOriginator.createMemento())

	myOriginator.restoreMemento(myCaretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", myOriginator.getState())

	myOriginator.restoreMemento(myCaretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", myOriginator.getState())
}
