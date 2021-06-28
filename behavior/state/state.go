package state

type state interface {
	addItem(int) error           //添加商品
	requestItem() error          //选择商品
	insertMoney(money int) error //插入纸币
	dispenseItem() error         //提供商品
}
