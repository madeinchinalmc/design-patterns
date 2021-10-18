package visitor

import "fmt"

type IGraphicsVisitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Rectangle)
}
type IShape interface {
	getType() string
	accept(visitor IGraphicsVisitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v IGraphicsVisitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v IGraphicsVisitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) accept(v IGraphicsVisitor) {
	v.visitForTriangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForTriangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates area for square")
}
func (a *middleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForTriangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

// 客户端代码

func RunApplication() {

	s := &Square{side: 2}
	c := &Circle{radius: 3}
	r := &Rectangle{l: 4, b: 5}

	ac := &areaCalculator{}
	s.accept(ac)
	c.accept(ac)
	r.accept(ac)

	m := &middleCoordinates{}
	s.accept(m)
	c.accept(m)
	r.accept(m)

}
