package main

import "fmt"

type state interface {
	requestItem() error
	dispenseItem() error
}

type selectItemMachineState struct {
	vendingMachine *vendingMachine
}

func (si *selectItemMachineState) requestItem() error {
	fmt.Println("Preparing selected item")
	si.vendingMachine.currentState = si.vendingMachine.itemRequested
	return nil
}

func (si *selectItemMachineState) dispenseItem() error {
	return fmt.Errorf("First select an item")
}

type dispenseItemMachineState struct {
	vendingMachine *vendingMachine
}

func (di *dispenseItemMachineState) requestItem() error {
	return fmt.Errorf("Dispensing item, please wait")
}

func (di *dispenseItemMachineState) dispenseItem() error {
	fmt.Println("Item dispensed")
	di.vendingMachine.currentState = di.vendingMachine.selectionAvailable
	return nil
}

type vendingMachine struct {
	selectionAvailable state
	itemRequested      state

	currentState state
}

func (vm *vendingMachine) requestItem() error {
	return vm.currentState.requestItem()
}

func (vm *vendingMachine) dispenseItem() error {
	return vm.currentState.dispenseItem()
}

func newVendingMachine() *vendingMachine {

	v := &vendingMachine{}

	selectItemState := &selectItemMachineState{
		vendingMachine: v,
	}

	dispenseItemState := &dispenseItemMachineState{
		vendingMachine: v,
	}

	v.selectionAvailable = selectItemState
	v.itemRequested = dispenseItemState
	v.currentState = selectItemState

	return v
}

func main() {

	vm := newVendingMachine()

	fmt.Println("Please select an item")

	err := vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("---")

	err = vm.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("---")

	fmt.Println("Please select an item")

	err = vm.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("---")

	err = vm.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("---")

}
