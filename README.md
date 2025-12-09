---
# Практическое задание 5

## ЭФМО-02-25 

## Алиев Каяхан Командар оглы
---
# Информация о проекте
Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)

## Цели занятия
1.	Установить и настроить PostgreSQL локально.
2.	Подключиться к БД из Go с помощью database/sql и драйвера PostgreSQL.
3.	Выполнить параметризованные запросы INSERT и SELECT.
4.	Корректно работать с context, пулом соединений и обработкой ошибок

## Файловая структура проекта:

<img width="1094" height="293" alt="image" src="https://github.com/user-attachments/assets/6a651772-5e45-419c-ba1b-d01e0d5ea80e" />


## Ключевые компоненты

main.go - точка входа приложения: инициализация логгера, создание и запуск экземпляра приложения через internal/app/app.go, обработка graceful shutdown

db.go - ядро приложения: инициализация подключения к базе данных

repository.go - интерфейс репозитория задач: контракты для методов работы с данными

go.mod - файл модуля Go: описание зависимостей (chi v5), версия Go, имя модуля

# Примечания по конфигурации и требования

Для запуска требуется:

Go: версия 1.25.1

PostgreSQL: версия не ниже 14

<img width="841" height="232" alt="Установка Git и Go" src="https://github.com/user-attachments/assets/8e01d831-5a7f-4376-8348-9052b240aec9" />


# Команды запуска/сборки
Для запуска http нужно выполнить 4 шага:
## 1) Клонировать данный репозиторий в удобную для вас папку:
```Powershell
git clone https://github.com/kayahan81/PZ5-baza
```
## 2) Перейти в папку http:
```Powershell
cd PZ5-baza
```
## 3) Загрузка зависимостей:
```Powershell
go mod tidy
```
## 4) Команда запуска
```Powershell
go run .
```

# Команда сборки
Для сборки бинарника и запуска .exe файла используются данные программы

```Powershell
go build -o server.exe .
server.exe
```
# Проверка работоспособности

## Базовый функционал

Проверка создания таблицы, добавления задач, подключения к бд и вывода задач

<img width="1165" height="273" alt="4 запуск и проверка" src="https://github.com/user-attachments/assets/5189cec0-f0c4-47be-9b51-58c0b6b7eca2" />

## Проверочные задания

Функция ListDone

<img width="720" height="470" alt="image" src="https://github.com/user-attachments/assets/31000924-ed13-4138-acbd-17f4ecfa2c6f" />

Функция FindByID

<img width="506" height="133" alt="image" src="https://github.com/user-attachments/assets/dc837bd3-a40c-4d0a-b64a-83e205ee478c" />

Функция CreateMany

<img width="397" height="133" alt="image" src="https://github.com/user-attachments/assets/8d184a2b-632c-4295-9b81-542e9fd5569e" />

Вывод и подсчёт текущих задач 

<img width="440" height="429" alt="image" src="https://github.com/user-attachments/assets/98beaed1-1cce-452b-93f8-c54e36fdafe9" />

#Ответы на вопросы

db.SetMaxOpenConns(1) - одно одновременное подключение, для данной практической работы хватает. Не ноль, потому что ноль означает что ограничений нет 

db.SetMaxIdleConns(1) // одно простаивающее соединение, для данной практической работы хватает

db.SetConnMaxLifetime(1 * time.Minute) // 1 минута ожидания. Если я буду ждать 30 минут чего-то, это что-то скорее всего не работает. Я не марсоходом управляю, чтобы пинг 30 минут был

$1, $2 используем чтобы защитить базу от SQL-инъекций (в файле пратики так написано)

db.QueryRow возвращает одну строку, доступ к данным через Scan.

db.Query используется для нескольких строк: нужно пройти по rows.Next() и сканировать каждую запись.

db.Exec команда, которая позволяет выполнять динамически сформированные SQL-инструкции 
