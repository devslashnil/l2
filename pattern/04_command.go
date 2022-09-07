package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Применимость:
	- Необходимость инкапсулирование запроса в виде структуры

	Плюсы:
	- Возможность работать с запросом как со структурой т.е. хранить и модифицировать

	Минусы:
	- Появление новой абстракции над вызовом запроса, меньший контроль над вызовом запроса

	Реальные примеры:
	- Протоколирование запросов
	- Отправка или хранение запросов как структур для параллельного выполнения,
	либо отправки по сети
*/

// Command интерфейс команды
type Command interface {
	Execute()
}

// Command1 реализует интерфейс Command
type Command1 struct {
	receiver *Receiver
}

// Execute выполняет команду
func (c *Command1) Execute() {
	c.receiver.Action1()
}

// Command2 реализует интерфейс Command
type Command2 struct {
	receiver *Receiver
}

// Execute выполняет команду
func (c *Command2) Execute() {
	c.receiver.Action2()
}

// Receiver реализация ресивера
type Receiver struct {
}

// Action1 действие Receiver
func (r *Receiver) Action1() {}

// Action2 действие Receiver
func (r *Receiver) Action2() {}

// Invoker реализация инвокера
type Invoker struct {
	commands []Command
}

// StoreCommand добавляет команду в инвокер
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand убирает команду в инвокере
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute выполняет все команды инвокера
func (i *Invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}
