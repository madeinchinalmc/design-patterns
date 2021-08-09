package decoration

import "fmt"

// 组件接口

type pizza interface {
	GetPrice() int
}

// 具体零件

type veggeMania struct {
}

func (p *veggeMania) GetPrice() int {
	return 15
}

// 具体装饰

type tomatoTopping struct {
	pizza pizza
}

func (p *tomatoTopping) GetPrice() int {
	concurrentPrice := p.pizza.GetPrice()
	fmt.Println("add tomato")
	return concurrentPrice + 10
}

type cheeseTopping struct {
	pizza pizza
}

func (p *cheeseTopping) GetPrice() int {
	concurrentPrice := p.pizza.GetPrice()
	fmt.Println("add cheese")
	return concurrentPrice + 8
}

type Customer struct{}

func (c *Customer) Buy() {
	vepizza := &veggeMania{}

	//add cheese
	pizzaWithCheese := &cheeseTopping{vepizza}

	//add tomato
	pizzaWithTomoto := &tomatoTopping{pizza: pizzaWithCheese}

	fmt.Println(pizzaWithTomoto.GetPrice())

}
