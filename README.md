# Тестовое задание Vortex

<!--Установка-->
## Установка
```git clone https://github.com/layzy-wolf/StatisticTest```

Перед началом работы необходимо установить ClickHouse

1. Настрока config

```cd config```

### необходимо заменить данные base.yaml на актуальные

```yaml
port: <PORT> 
db_socket: <CLICKHOUSE_SOCKET>
db_name: <CLICKHOUSE_DB_NAME>
db_user: <CLICKHOUSE_USER>
db_password: <CLICKHOUSE_PASSWORD>
```

2. Выйти в корневую директорию проекта

```cd ..```

3. Запустить приложение 

```go run cmd/main.go```

<!-- Тестирование -->

### Тесты можно запустить через скомпелированный файл

```./tests/tests.test```