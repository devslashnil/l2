package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Применимость:
	- Система должна быть легко расширяемой для новых типов
	- Необходимость унифицировать создание объектов некоторого интерфейса
	- Необходим дополнительный контроль за созданием структур

	Плюсы:
	- Гибкость при интеграции объектов новых типов в существующую систему
	- Возможность обобщенной работы со структурами порождённой фабричным методов

	Минусы:
	- Порождающие объекты должны придерживаться одного интерфейса в отличие от
	объектов порождаемых Abstract Factory

	Реальные примеры:
	- Создание подходящей структуры в зависимости от поступивших данных
	- Создания ответа на пришедший запрос
*/

// action возможные параметры на вход фабрики
type action string

// A, B константы action
const (
	A action = "A"
	B action = "B"
)

// FactoryProduct интерфейс продукта, порождаемого Creator
type FactoryProduct interface {
	Foo()
}

// Creator интерфейс фабрики
type Creator interface {
	FactoryMethod(a action) FactoryProduct
}

// NewCreator конструктор фабрики
func NewCreator() Creator {
	return &Creator1{}
}

// Creator1 имплементация фабрики Creator
type Creator1 struct {
}

// FactoryMethod фабричный метод Creator1
func (_ *Creator1) FactoryMethod(a action) (p FactoryProduct) {
	switch a {
	case A:
		p = &ProductA{}
	case B:
		p = &ProductB{}
	}
	return
}

// ProductA структура создаваемого фабрикой интерфейса FactoryProduct
type ProductA struct {
}

// Foo метод ProductA
func (_ *ProductA) Foo() {}

// ProductB структура создаваемого фабрикой интерфейса FactoryProduct
type ProductB struct {
}

// Foo метод ProductB
func (_ *ProductB) Foo() {}
