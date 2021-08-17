package chain

import "fmt"

// 病人

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

//处理者接口
type department interface {
	execute(*patient)
	setNext(department)
}

//具体处理者 -> 前台
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		if r.next != nil {
			r.next.execute(p)
		}
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *reception) setNext(d department) {
	r.next = d
}

//具体处理类 -> 医生
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *doctor) setNext(p department) {
	d.next = p
}

//具体处理类 -> 药房
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medical already given to patient")
		if m.next != nil {
			m.next.execute(p)
		}
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	if m.next != nil {
		m.next.execute(p)
	}
}

func (m *medical) setNext(p department) {
	m.next = p
}

// 具体处理类 —> 收银台
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
		if c.next != nil {
			c.next.execute(p)
		}
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	p.paymentDone = true
	if c.next != nil {
		c.next.execute(p)
	}
}

func (c *cashier) setNext(p department) {
	c.next = p
}

func RunApplication() {
	c := &cashier{}

	m := &medical{}
	m.setNext(c)

	d := &doctor{}
	d.setNext(m)

	r := &reception{}
	r.setNext(d)

	p := &patient{name: "tom"}
	r.execute(p)

	fmt.Println("------------------")
	d.execute(p)
}
