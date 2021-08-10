package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a Huipu Printer")
}

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Windows struct {
	Printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}

type Mac struct {
	Printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}

func ApplicationRun() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}
	macComputer := &Mac{}
	winComputer := &Windows{}

	macComputer.Printer = hpPrinter
	macComputer.Print()
	fmt.Println()

	macComputer.Printer = epsonPrinter
	macComputer.Print()
	fmt.Println()

	winComputer.Printer = hpPrinter
	winComputer.Print()
	fmt.Println()

	winComputer.Printer = epsonPrinter
	winComputer.Print()
	fmt.Println()

}
