package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Применимость:
	- Необходимость определять новые операции над структурами не меняя самих структур
	- Требуется применить много не связанных друг с другом операций над объектов
	- Алгоритму нужна работа с набором структур
	- Требуется добавлять много функционала по работе со структурой, который выходит
	за зону ответственности структуры

	Плюсы:
	- Разделение зон ответственностей между визитором и элементами для соблюдения
	принципа единственной ответственности
	- Возможность хранить код по работе со структурами в одном месте, а не выносить
	повторяющуюся логику в каждую структуре элементу

	Минусы:
	- Каждый новый элемент требует расширение методов интерфейса Visitor

	Реальные примеры:
	- Необходимость работать с отрытым API, которе нет возможности изменить, но
	при этом имеется возможность "посетить"
*/

// Visitor интерфейс визитора
type Visitor interface {
	VisitConcreteElement1(e *ConcreteElement1)
	VisitConcreteElement2(e *ConcreteElement2)
}

// ConcreteVisitor реализация Visitor
type ConcreteVisitor struct {
}

// VisitConcreteElement1 имплементация метода интерфейса Visitor
func (_ *ConcreteVisitor) VisitConcreteElement1(e *ConcreteElement1) {
	e.Foo()
}

// VisitConcreteElement2 имплементация метода интерфейса Visitor
func (_ *ConcreteVisitor) VisitConcreteElement2(e *ConcreteElement2) {
	e.Bar()
}

// Client коллекция элементов
type Client struct {
	elements []Element
}

// Add добавляет элемент в коллекцию Client
func (c *Client) Add(e Element) {
	c.elements = append(c.elements, e)
}

// Accept посещает все элементы коллекции некоторым визитором
func (c *Client) Accept(v Visitor) {
	for _, e := range c.elements {
		e.Accept(v)
	}
}

// Element интерфейс, который визитор может посетить
type Element interface {
	Accept(v Visitor)
}

// ConcreteElement1 реализация Element
type ConcreteElement1 struct {
}

// Accept реализация посещение визитором элемента
func (e *ConcreteElement1) Accept(v Visitor) {
	v.VisitConcreteElement1(e)
}

// Foo метод элемента
func (_ *ConcreteElement1) Foo() {}

// ConcreteElement2 реализация Element
type ConcreteElement2 struct {
}

// Accept реализация посещение визитором элемента
func (e *ConcreteElement2) Accept(v Visitor) {
	v.VisitConcreteElement2(e)
}

// Bar метод элемента
func (_ *ConcreteElement2) Bar() {}
