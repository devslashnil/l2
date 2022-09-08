Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

Переменная ```err``` содержит интерфейсный тип значения ```nil``` и дескриптором типа ```*customError```, а
интерфейсный тип равен ```nil```, если и дескриптор типа равен ```nil```, для ликвидации этой ошибки 
```func test()``` должен возвращать интерфейс ```error```