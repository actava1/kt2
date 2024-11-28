package main

// Элементы
type Shape interface {
	Accept(visitor Visitor)
}

type Circle struct{}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

type Square struct{}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Посетитель
type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

type AreaCalculator struct{}
