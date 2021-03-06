# Вопросы на собеседование Golang

## TODO

[ ] Расширить кругозов вопросов

#### Знакомы ли вы с понятиями Stateless и Stateful? Можете объяснить что они значат и чем различаются?

> https://youtu.be/WPCz_U7D8PI

#### Какие приемущества есть у Golang?

> Компилируемость
> Кокнкурентность
> Паралелизм
> 

#### Что будет в качестве значения?

> - `var i int`
> - `var s struct{}`
> - `var s2 *struct{}`

Example: https://play.golang.org/p/YzxAGf3m4Lc

#### Как скопировать одну map в другую?

> range

#### Что такое гонка в Go?

> Начальные вопросы

#### Что можете сказать про пакет sync? Какие примитивы из данного пакета вы знаете?

> Знание работы о синхранизации:
> 
> - WaitGroup
> - Mutex
> - Cond
> - Map
> - Once
> - Pool

#### В чем разница между Mutex и RWMutex?

> Понимание отличия работы

#### Знакомы ли вы с пакетом Atomic? Что он делает?

> Понимание того, для чего данный пакет и когда его использовать

#### Знакомы ли вы с понятием Greceful Shutdown? Для чего он? Как можно организовать?

> Понимание, как работает остановщик и для чего он нужен. Здесь начинаются каналы и частично пакет os (os.SIGTERM).

#### Если в приложении множество горутин, как можно передавать данные между ними? 

> Понимание того, как принято работать с goroutine-ми и есть ли понимании практики работы с каналами.

#### А как можно завершить множество goroutine при использовании Graceful shutdown?

> Тот же вопрос что и выше, только совмещаем два понятия и смотрим, что человек не теряется.

#### 1. Что выведет код?

```golang
package main

import "fmt"

type example struct {
    X, Y string
}

var (
    a = example{"Aus","Can"} 
    b = &example{"Jap","Kor"}
    c = example{X:"US",Y:"UK"}
    d = example{}
)

func main() {
    e := b
    b.X = "Rus"
    f := *b
    fmt.Println("a:\t",a)
	fmt.Println("b:\t",b)
	fmt.Println("c:\t",c)
	fmt.Println("d:\t",d)
	fmt.Println("e:\t",e)
	fmt.Println("d:\t",f)	
}
```

https://play.golang.org/p/02MOucvqU9G

Вывод:

```
a:       {Aus Can}
b:       &{Rus Kor}
c:       {US UK}
d:       { }
e:       &{Rus Kor}
d:       {Rus Kor}
```

#### 2. Что выведет код?

```
package main

var x = 3
var n int = 5.0 / 3 * 3
var m int = 5.0 / x * 3

func main() {
	println(n, m)
}
```

https://play.golang.org/p/dJNe-KX-Srg

Вывод:

```
5 3
```

#### 3. Что выведет код?

```
package main
 
import (
    "fmt"
)
 
type Temp struct {
}
 
func main() {
    var pnt *Temp       // pointer
    var inf interface{} // interface declaration
    inf = pnt           // inf is a non-nil interface holding a nil pointer (pnt)
 
    fmt.Printf("pnt is a nil pointer: %v\n", pnt == nil)
    fmt.Printf("inf is a nil interface: %v\n", inf == nil)
}
```

https://play.golang.org/p/SjpSv4QemTS

А так?

```
fmt.Printf("inf is a interface holding a nil pointer: %v\n", inf == (*Temp)(nil))
```

https://play.golang.org/p/2Z2D4SD7owu

Output:

```
pnt is a nil pointer: true
inf is a nil interface: false
inf is a interface holding a nil pointer: true
```

Почему так?

Source: https://www.golangprograms.com/how-to-check-pointer-or-interface-is-nil.html

#### Если программа работает не так, как предполагается, какими способами можно начать проверять наличие проблем? Как проводите Research проблемы?

> Понимание, на что в первую очередь обращяет внимание, есть ли опыт в таких ситуациях. Обычно смотрят на:
>
> - логи
> - метрики
> - использование памяти или CPU.
> - просмотр зависимостей (может «зависать» другой сервис или база).

#### С какими БД вы работали?

#### Какой функцией можно проанализировать запрос к базе данных (например MySQL)?

>  Понимание, как можно оптимизировать запрос. Команда EXPLAIN. 
>
> - https://habr.com/ru/post/31129/ 
> - https://habr.com/ru/post/211022/

#### Чем строковая от колоночной БД отличается?

#### Если приходят две задачи: одна с багой, другая с фичой. Какую выберете?


---

## Вопросы для компании

- Расскажите о проекте?
- Размер команды и сколько команд?
- График работы? Можно ли его двигать?
- Какие планы на ближайшие 6 месяцев?
- Какой стэк?
- Какие задачи предстоит решать?
- Отпуск?
- Командировки?
- Обучение за счет компании?
- Повышение квалификации?
- В какой валюте ЗП – RUB/USD/EUR/GBP?
- Какие сроки по фидбэку по вакансии?
- Испытательный срок?
  - На какой период? 
  - Какая зарплата на этот срок? 
  - Возможен ли пересмотр по результатам испытательного срока. 

### Если рекрутер

- Сколько этапов собеседования?
- Есть ли техническое задание?
