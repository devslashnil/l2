package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Применимость:
	-

	Плюсы:
	-

	Минусы:
	-

	Реальные примеры:
	-
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
