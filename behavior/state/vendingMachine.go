package state

import "fmt"

type vendingMachine struct {
	hasItem       state //有商品
	itemRequested state //请求商品
	hasMoney      state //收到纸币
	noItem        state //无商品

	currentState state //当前状态

	itemCount int //商品数量
	itemPrice int //商品价格
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &hasItemState{
		vendingMachine: v,
	}

	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}

	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}

	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)

	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
