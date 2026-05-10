# zero-agency-test

[![wakatime](https://wakatime.com/badge/user/b2a0c08d-61f2-4144-ba78-aab13a59cb9f/project/86d33f83-a367-4535-80e9-9c5e306a02e4.svg)](https://wakatime.com/badge/user/b2a0c08d-61f2-4144-ba78-aab13a59cb9f/project/86d33f83-a367-4535-80e9-9c5e306a02e4)

## О проекте

В данном репозитории представлено выполненное тестовое задание для компании [Zero Agency](https://zeroagency.ru/).

Текст задания представлен в файле [task.md](./task.md).

### Структура

```bash
├── cmd
│   ├── bot-http    - Версия с использованием стандартных библиотек
│   └── bot-openai  - Версия с использованием официальной библиотеки openai/openai-go
├── configs         - Файлы конфигурации
├── internal
│   ├── app
│   │   └── bot     - Бот
│   ├── classifier  - Классификатор
│   ├── client      - Клиенты для взаимодействия с ИИ
│   │   ├── http    - С использованием стандартных библиотек
│   │   └── openai  - С использованием официальной библиотеки openai/openai-go
│   ├── config
│   ├── router      - Роутер (Диспетчер)
│   ├── shared
│   │   └── tags    - Теги
│   └── skills      - Навыки
├── pkg
│   └── logger
└── tests
    └── integration - Интеграционные тесты  
        ├── http
        └── openai
```

### Подготовка

Перед первым запуском создайте и заполните конфигурационный файл

Создание:

```bash
cp ./configs/config.yaml.example ./configs/config.yaml
```

Пример конфигурационного файла:

```yaml
env: local # Возможные варианты: local, dev, prod
open_ai:
  model: gpt-oss-20b
  url: <url>/v1
  api_key: <api-key>
```

### Запуск

#### Версия с использованием стандартных библиотек

Запуск:

```bash
make run-http
```

Билд:

```bash
make build-http
```

Запуск с помощью Docker `(в интерактивном режиме)`:

```bash
make docker-http
```

#### Версия с использованием официальной библиотеки [`openai/openai-go`](https://github.com/openai/openai-go)

Запуск:

```bash
make run-openai
```

Билд:

```bash
make build-openai
```

Запуск с помощью Docker `(в интерактивном режиме)`:

```bash
make docker-openai
```

## Тесты

### Unit-тесты

Простой прогон тестов:

```bash
make test
```

Прогон тестов с выводом покрытия:

```bash
make test-cover
```

### Интеграционные тесты

#### Версия с использованием стандартных библиотек

```bash
make test-integration-http
```

#### Версия с использованием официальной библиотеки [`openai/openai-go`](https://github.com/openai/openai-go)

```bash
make test-integration-openai
```
