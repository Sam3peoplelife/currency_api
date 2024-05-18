1. Запуск додатку:

    ```bash
    docker-compose up --build
    ```
    
    Адреса - [http://localhost:8080](http://localhost:8080).

2. Зупинка:

    ```bash
    docker-compose down
    ```

## Структура проекту

```
api/
|-- main.go
|-- handlers.go
|-- migrations/
|   |-- migrations.go
|-- email/
|   |-- email.go
|-- models/
|   |-- user.go
|-- templates/
|   |-- index.html
|-- static/
|   |-- style.css
|-- utils/
|   |-- fetch.go
|-- go.mod
|-- go.sum
|-- handlers_test.go
|-- Dockerfile
|-- docker-compose.yml
|-- README.md
```
