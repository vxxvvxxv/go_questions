# Вопросы на собеседование Golang

## Links:


### Теория:

- https://habr.com/ru/post/654569/
- https://habr.com/ru/post/670974/
- https://parshinpn.ru/ru/blog/golang-interview
- Good: https://habr.com/ru/post/325468/

### Задачи:

- https://habr.com/ru/company/rebrainme/blog/540354/

#### SQL

Напишите запрос:

```
Есть две таблицы 
1- users (id, first_name, last_name) , 
2- transactions(id, user_id, amount, currency)

Требуется построить отчет user_id, first_name, last_name, количество транзакций.
```

Online:

- Postgres: https://sqlize.online/s/1t
- MySQL: https://sqlize.online/s/vt


Ответ (PostgreSQL):

```sql
SELECT 
  u.id as user_id,
  u.first_name,
  u.last_name,
  COUNT(t.id) as counts
FROM users as u
LEFT JOIN transactions as t ON u.id = t.user_id
GROUP BY u.id, u.first_name, u.last_name
ORDER BY u.id
;

```
Answer:

- MySQL URL: https://sqlize.online/s/Lt
- Postgres URL: https://sqlize.online/s/It

#### Bites

Url: https://play.golang.org/p/-O5CF8y5FZ0

```golang
package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	dir1 = append(dir1, "suffix"...)
	fmt.Printf("%s\n", dir1)
	fmt.Printf("%s\n", dir2)
}
```

Output:
```
AAAAsuffix
uffixBBBB

Program exited.
```

Почему так? https://www.golangforall.com/ru/post/golang-slice.html

Slice - это всегда ссылка на массив, и пока нет изменения по cap, slice будет ссылаться и меняться у всех, кто ссылается на массив-исходник.

```golang
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

#### Types

Url: https://play.golang.org/p/VElWNaDQ3bV

Почему так: https://habr.com/ru/post/325468/

> В обоих случаях мы как бы видим nil, но есть большая разница между "интерфейс с переменной внутри, чьё значение равно nil" и "интерфейс без переменной внутри". Теперь, понимая эту разницу, попробуйте спутать вот эти два примера:

![Alt text](https://hsto.org/r/w1560/files/ac9/905/1fb/ac99051fbb924e6d8912f479915b7a78.png "a title")

#### Merge

Url: https://go.dev/play/p/Ga2YbYCr9m1

```
a=[], b=[b], n=1 => res=[b], err=<nil>
a=[], b=[b], n=2 => res=[], err=error process msg
a=[apple orange kiwi plum], b=[apple banana orange plum], n=5 => res=[apple orange kiwi plum banana], err=<nil>
a=[a], b=[b], n=2 => res=[a b], err=<nil>
a=[a], b=[a b], n=2 => res=[a b], err=<nil>
a=[a b], b=[a], n=2 => res=[a b], err=<nil>
```

- Answer (not correct): https://go.dev/play/p/CQYqicDwMSO
- Answer (can be correct): https://go.dev/play/p/VwezksQvpN3

---

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

### Ответы для рекрутера

- Сколько хотите зарабатывать?
  - Не отвечаем на вопрос, стараемся уйти от ответа.

Мне нравится ваша компания и я давно хотел в ней поработать.
Я пользуюсь продуктами компании и хотел бы быть частью команды, которая помогаем людям.


## БД

### Индексы

#### B-tree (B-дерево)

Бинарный поиск имеет алгоритмическую сложность O(log(n)).

Принцип работы одного из важнейших индексов в базе данных (индекс на основе B-дерева) основан именно на рассмотренном нами выше принципе — возможности хранить данные в отсортированном виде.

С помощью B-дерева можно проиндексировать любые данные, которые могут быть отсортированы, т. е. для которых применимы операции сравнения больше/меньше/равно. Сюда можно отнести числа, строки, даты и время, логический тип и любые данные, которые можно ими закодировать.

Кроме того, индекс на основе B-дерева ускоряет сортировку результатов, если в ORDER BY указано проиндексированное поле.

--

Предикат P может вычисляться от значения нескольких колонок. В этом случае для ускорения запроса используется индекс, построенный не для одной колонки, а для нескольких. Такие индексы называют составными.

Или мы можем создать индекс не по полю таблицы, а по функции или скалярному выражению с одной или несколькими колонками таблицы (такие индексы называют функциональными или индексами по выражению). Это позволяет быстро находить данные в таблице по результатам вычислений.

#### GiST и SP-GiST

GiST — сокращение от «generalized search tree». Это сбалансированное дерево поиска.

Для тех операций, в которых упорядочивания не имеет смысла, например, геоданные и геометрические объекты.

Он позволяет распределить данные любого типа по сбалансированному дереву и использовать это дерево для поиска по самым разным условиям. Если при построении B-дерева мы сортируем все множество объектов и делим его на части по принципу больше-меньше, при построении GiST индексов можно реализовать любой принцип разбиения любого множества объектов.

Например, в GiST-индекс можно уложить R-дерево для пространственных данных с поддержкой операторов взаимного расположения (находится слева, справа; содержит и т. д.). Такой индекс доступен в PostgreSQL и может быть полезен при разработке геоинформационных систем, в которых возникают запросы вида «получить множество объектов на карте, находящихся от заданной точки на расстоянии не более 1 км».

#### GIN

GIN индексы полезны для организации полнотекстового поиска и для индексации таких типов данных, как массивы или jsonb.

