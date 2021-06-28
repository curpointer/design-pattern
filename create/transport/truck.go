package transport

import "fmt"

type truck struct {
	name string
}

func (t *truck) SetName(name string) {
	t.name = name
}

func (t *truck) GetName() string {
	return t.name
}

func (t *truck) Deliver() {
	fmt.Printf("A truck delivers goods, the name of truck is %s\n", t.GetName())
}

func NewTruck() iTransport {
	return &truck{
		name: "panda",
	}
}
