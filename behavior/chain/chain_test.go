package chain

import "testing"

func TestChain(t *testing.T) {
	// 前台—>医生->收银->药房取药
	patient := &patient{
		name: "John",
	}
	reception := &reception{}
	doctor := &doctor{}
	reception.setNext(doctor)

	cashier := &cashier{}
	doctor.setNext(cashier)

	medicinal := &medicinal{}
	cashier.setNext(medicinal)

	reception.execute(patient)

}
