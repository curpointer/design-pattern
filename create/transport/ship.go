package transport

import "fmt"

type ship struct {
	name string
}

func (s *ship) SetName(name string) {
	s.name = name
}

func (s *ship) GetName() string {
	return s.name
}

func (s *ship) Deliver() {
	fmt.Printf("A ship delivers goods, the name of ship is %s\n", s.GetName())
}

func NewShip() iTransport {
	return &ship{
		name: "bird",
	}
}
