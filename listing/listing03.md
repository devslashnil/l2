Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```
- ```fmt.Println(err)``` вывел значение интерфейсного типа ```error```
- ```go fmt.Println(err == nil)``` вывел ```false```, т.к. интерфейсный тип равен ```nil```, когда
дескриптор интерфейсного типа тоже равен ```nil```, а в данном случае он равен ```*os.PathError```
- Интерфейс представляет собой пару конкретного значения ```unsafe.Pointer``` и дескриптора типа ```*itab```,
содержащего информацию о типе интерфейса и конкретного значения. Информация в этих подструктурах содержит указатели на
функции которые вызываются данным интерфейсом.
