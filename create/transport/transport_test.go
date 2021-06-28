package transport

import "testing"

func TestTransport(t *testing.T) {
	ship, _ := GetTransport(TRANSPORT_SHIPPING)
	ship.SetName("sasa")
	PrintTransport(ship)
	truck, _ := GetTransport(TRANSPORT_LAND)
	truck.SetName("tata")
	PrintTransport(truck)
}
