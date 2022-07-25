package main

import (
	"fmt"
)

//code related with computers
type computer interface {
	print()
	setPrinter(printer)
}

//desktop computer, struct definition and interface methods implementation
type desktop struct {
	printer printer
}

func (d *desktop) print() {
	fmt.Println("Print document for desktop PC")
	d.printer.printFile()
}

func (d *desktop) setPrinter(p printer) {
	d.printer = p
}

type laptop struct {
	printer printer
}

func (l *laptop) print() {
	fmt.Println("Print document for laptop")
	l.printer.printFile()
}

func (l *laptop) setPrinter(p printer) {
	l.printer = p
}

//code related to printers
type printer interface {
	printFile()
}

type printerOne struct{}

func (p printerOne) printFile() {
	fmt.Println("Printing with the printer ONE")
}

type printerTwo struct{}

func (p printerTwo) printFile() {
	fmt.Println("Printing with the printer TWO")
}

func main() {
	p1 := printerOne{}
	p2 := printerTwo{}

	desktopPC := desktop{}
	desktopPC.setPrinter(p1)
	desktopPC.print()
	fmt.Println()
	desktopPC.setPrinter(p2)
	desktopPC.print()

	fmt.Println("----------------------")

	laptopPC := laptop{}
	laptopPC.setPrinter(p1)
	laptopPC.print()
	fmt.Println()
	laptopPC.setPrinter(p2)
	laptopPC.print()
}
