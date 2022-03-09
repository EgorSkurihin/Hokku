# Echo REST API

REST API сервис с использованием языка Golang и фреймворка Echo

## Использованные библиотеки:
1. `github.com/labstack/echo/v4` - Основной фреймворк
1. `github.com/go-ozzo/ozzo-validation` - Валидация
1. `github.com/BurntSushi/toml` - Парсинг и использование toml-конфига
1. `github.com/gorilla/sessions` - Аутентификация при помощи сессий
1. `github.com/stretchr/testify` - Тестирование
1. `github.com/swaggo/swag` - Документация

## База данных MySQL:
1. `github.com/go-sql-driver/mysql` - MySQL-драйвер
1. `github.com/golang-migrate/migrate` - Миграции БД

## Документация
1. `github.com/swaggo/swag`, `github.com/swaggo/echo-swagger` - Swagger Документация

## Запуск
1. `git clone github.com/EgorSkurihin/Hokku`
1. `cd Hokku`
1. `docker-compose build`
1. `docker-compose up`
1. Проверить работу сервиса на http://localhost:1323/health
1. Посмотреть Swagger документацию на http://localhost:1323/swagger/index.html
