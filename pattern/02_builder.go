package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Применимость:
	- Применяется когда нужно выделить конфигурацию объекта в отдельную структуру
	для упрощения создания структур т.к. код создания структуры разделяется со самой структурой

	Плюсы:
	- Предоставляет интерфейс для изменения создаваемой структуры
	- Разделяет код создания структуры и самой структуры, что в некоторых случаях
	позволяет добиться сохранения принципа единственной ответственности SOLID

	Минусы:
	- Структуры создаваемые строителем не могут быть имутабельными
	- Каждая структура требует собственный строитель

	Реальные примеры:
	- Необходимость конструировать кастомный XML для ответа на разные запросы
	- Необходимость собрать по частям однотипные HTML страницы
*/

// Builder интерфейс строителя
type Builder interface {
	Foo(str string)
	Bar(str string)
	Build() Product
}

// Director реализует менеджера по сборке
type Director struct {
	builder Builder
}

// Construct задает билдеру порядок и контент для сборки
func (d *Director) Construct() {
	d.builder.Foo("Foo")
	d.builder.Bar("Bar")
}

// Product структура, которую собирает билдер
type Product struct {
	foo string
	bar string
}

// ConcreteBuilder имплементирует Builder
type ConcreteBuilder struct {
	product *Product
}

// Foo задает foo в создаваемом Product
func (b *ConcreteBuilder) Foo(str string) {
	b.product.foo = str
}

// Bar задает bar в создаваемом Product
func (b *ConcreteBuilder) Bar(str string) {
	b.product.bar = str
}

// Build возвращает построенный Product
func (b *ConcreteBuilder) Build() *Product {
	return b.product
}
