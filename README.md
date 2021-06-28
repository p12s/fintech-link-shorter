![License](https://img.shields.io/github/license/p12s/fintech-link-shorter?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/fintech-link-shorter?style=plastic)
<img src="https://github.com/p12s/fintech-link-shorter/workflows/lint-test-build/badge.svg?branch=master">

**Внимание:** *Тестовое задание найдено на просторах github-а. Для обучения и тренировки, попробовал решить ее в меру своего понимания. На ревью не отправлял, за оптимальность не ручаюсь.*

# Сервис сокращения ссылок

## Задача
Создать сервис, который будет как укорачивать ссылку, так и по-короткой ссылке возвращать исходную.   
Подробнее [здесь](task.md)

## Нефункциональные требования
- ❌ В качестве хранилица использовать РСУБД(postgresql, sqllite)
  postgresql можно запустить в docker:
  docker run --rm -p 5432:5432 postgres:10.5
- ❌ В качестве структуры веб сервиса - https://github.com/golang-standards/project-layout
- ❌ Сервис можно реализовать как стандартной библиотекой(net/http), так и фреймворками gin, echo
- ❌ Запросы в БД на pure sql, либо https://github.com/Masterminds/squirrel
- ✅ Короткие ссылки должны основываться на id записи(sequence) в БД, переведённой в систему счисления с алфавитом [A-Za-z0-9]

docker run --name=link-shorter-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm nouchka/sqlite3
docker exec -it 5a9d643bdaf3 /bin/bash

