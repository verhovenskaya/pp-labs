// Создать структуру Person с полями name и age. Реализовать метод для вывода информации о человеке
package main

import (
	"fmt"
	"math"
)

// Структура Person
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) Info() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d лет", p.name, p.age)
}

// Метод birthday, увеличивает возраст на 1 год
func (p *Person) birthday() {
	p.age++
}

// Структура Circle
type Circle struct {
	radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Интерфейс Shape
type Shape interface {
	Area() float64
}

// Структура Rectangle
type Rectangle struct {
	width, height float64
}

// Реализация метода Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Функция для вывода площади каждого объекта
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}

// Интерфейс Stringer
type Stringer interface {
	String() string
}

// Структура Book
type Book struct {
	title  string
	author string
}

// Реализация метода String для Book
func (b Book) String() string {
	return fmt.Sprintf("Название: %s, Автор: %s", b.title, b.author)
}

func main() {
	// Пример использования Person
	person := Person{name: "Иван", age: 30}
	fmt.Println(person.Info())
	person.birthday()
	fmt.Println("После дня рождения:", person.Info())

	// Пример использования Circle и Rectangle
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 3}

	shapes := []Shape{circle, rectangle}
	PrintAreas(shapes)

	// Пример использования Book и интерфейса Stringer
	book := Book{title: "1984", author: "Джордж Оруэлл"}
	fmt.Println(book.String())
}
