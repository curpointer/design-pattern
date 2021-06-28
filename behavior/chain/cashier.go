package chain

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("病人在收银台支付完成.")
		c.next.execute(p)
		return
	}
	fmt.Println("病人正在收银台支付中.")
	p.paymentDone = true
	c.next.execute(p)
}

func (c *cashier) setNext(next department) {
	c.next = next
}
