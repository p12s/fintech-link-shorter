# Настройка swagger для net/http роутинга (без фреймворков) с использованием http-swagger

## 1. Документируем методы

([официальный репозиторий](https://github.com/swaggo/swag))
```go
// @title Link shorter API
// @version 0.0.1
// @description This is an API Server for link shorter

// @host localhost:80
// @BasePath /
// @query.collection.format multi

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	...
}
```
```go
// Short @Summary Getting a short link
// @Tags short
// @Description Getting a short link by a long one
// @ID get-short-link
// @Accept  json
// @Produce  json
// @Param url body shorter.UserLink true "long link"
// @Success 200 {object} shorter.UserLink
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /short [post]
func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {
	...
}
```

## 2. Генерируем папку с документацией  

Скачиваем пакет:  
```
go get -u github.com/swaggo/swag/cmd/swag
```
запускаем инициализацию:
```
/Library/go/go1.16.5/bin/bin/swag init -g ./cmd/main.go
```
в корне должна быдет появиться директория 'docs'   

## 3. Открываем документацию в браузере

В главном хендлере (в примере - ./cmd/main.go) подключаем модули и добавляем адрес для открытия документации в браузере:
```go
import (
	...
    httpSwagger "github.com/swaggo/http-swagger"
    
    _ "github.com/p12s/fintech-link-shorter/docs"
)
...

func (s *Server) Run(port string, handler handler.Handler) error {
    ...
    http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
    ...
}
```
Документация должна будет открыться по адресу:  
http://localhost:80/swagger/index.html