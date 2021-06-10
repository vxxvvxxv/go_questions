# Вопросы на собеседование Golang

#### Что будет в качестве значения?

> `var i int`
> `var s struct{}`
> `var s2 *struct{}`

#### Что под капотом функции Pipe?

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


#### Если приходят две задачи: одна с багой, другая с фичой. Какую выберете?
