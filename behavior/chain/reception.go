package chain

import "fmt"

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("病人已经在前台完成登记.")
		r.next.execute(p)
		return
	}
	fmt.Println("病人在前台登记中.")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}
