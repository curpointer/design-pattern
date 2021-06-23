package transport

import "fmt"

type iTransport interface {
	SetName(name string)
	GetName() string
	Deliver()
}

func GetTransport(transportType string) (iTransport, error) {
	switch transportType {
	case TRANSPORT_LAND:
		return NewTruck(), nil
	case TRANSPORT_SHIPPING:
		return NewShip(), nil
	default:
		return nil, fmt.Errorf("not matched transport")
	}
}

func PrintTransport(i iTransport) {
	fmt.Printf("Transport is:%s\n", i.GetName())
	i.Deliver()
}
