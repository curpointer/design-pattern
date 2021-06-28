package chain

import "fmt"

type medicinal struct {
	next department
}

func (m *medicinal) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("病人已经到药房取完药.")
		m.next.execute(p)
		return
	}
	fmt.Println("病人正在药房取药.")
	p.medicineDone = true
	// m.next.execute(p)
}

func (m *medicinal) setNext(next department) {
	m.next = next
}
