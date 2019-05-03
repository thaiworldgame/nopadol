package main

import (
	"fmt"
	"math"
)

const englishPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
		fmt.Println(1)
	case spanish:
		prefix = spanishHelloPrefix
		fmt.Println(2)
	default:
		prefix = englishPrefix
	}
	return
}

func Add(x, y int) int {
	return x + y
}

type Shape interface {
	Area() float64
}


type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64  {
	return r.Width * r.Height
}

func (c Circle) Area() float64  {
	return math.Pi * c.Radius * c.Radius
}