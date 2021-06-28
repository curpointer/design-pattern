package chain

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("医生已经给病人检查完毕.")
		d.next.execute(p)
		return
	}
	fmt.Println("医生正在给病人检查.")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}
