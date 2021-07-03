![License](https://img.shields.io/github/license/p12s/fintech-link-shorter?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/fintech-link-shorter?style=plastic)
[![Coverage Status](https://img.shields.io/codecov/c/github/p12s/fintech-link-shorter/master.svg)](https://codecov.io/gh/p12s/fintech-link-shorter)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/fintech-link-shorter)](https://goreportcard.com/report/github.com/p12s/fintech-link-shorter)
<img src="https://github.com/p12s/fintech-link-shorter/workflows/lint-build/badge.svg?branch=master">

**Внимание:** *Тестовое задание найдено на просторах github-а. Для обучения и тренировки, попробовал решить ее в меру своего понимания. На ревью не отправлял, за оптимальность не ручаюсь.*

# Сервис сокращения ссылок

## Задача
Создать сервис, который будет как укорачивать ссылку, так и по-короткой ссылке возвращать исходную.   
Подробнее [здесь](task.md)

## Нефункциональные требования
- ✅ В качестве хранилица использовать РСУБД(postgresql, sqllite)  
  postgresql можно запустить в docker:  
  docker run --rm -p 5432:5432 postgres:10.5  
  **Выбран Sqlite3, файл создается в корне проекта и при достижении определенного в конфиге размера, пересоздается (чтобы не хостить БД и не перегружать тестовый стенд)**
- ✅ В качестве структуры веб сервиса - https://github.com/golang-standards/project-layout
- ✅ Сервис можно реализовать как стандартной библиотекой(net/http), так и фреймворками gin, echo  
  **Выбран пакет net/http с нативной реализацией роутинга**
- ✅ Запросы в БД на pure sql, либо https://github.com/Masterminds/squirrel
- ✅ Короткие ссылки должны основываться на id записи(sequence) в БД, переведённой в систему счисления с алфавитом [A-Za-z0-9]
  **Id ссылки в БД переводится в 62-ичную систему счисления 0-9a-zA-Z: https://p12s.ru/1N (p12s.ru - короткий домен из .env, 1N - перекодированный ID)**

## Дополнительные улучшения
- Не используются фреймворки, роутинг на net/http  
- Добавлена документация с помощью http-swagger  

## Что можно усовершенствовать
- вместо стандартного net/http-пакета попробовать [fasthttp](https://github.com/valyala/fasthttp)    
  у него меньше аллокаций памяти и быстрее скорость за счет использования кодогенерации вместо рефлекции (текущий этап моих знаний)  

