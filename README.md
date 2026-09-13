# executor

Реализация выполнения внешних команд для Sudzekai Web OS.

## Установка

```bash
go get github.com/sudzekai-web-os/executor@v1.0.1
```

## Использование

`NewExecutor` принимает фабрику логгеров и возвращает `abstractions.IExecutor`. Результат содержит стандартный вывод, стандартный поток ошибок и ошибку выполнения.

```go
factory := logging.NewLoggerFactory(os.Stdout)
runner := executor.NewExecutor(factory)

result := runner.Execute("go", "version")
if result.Error != nil {
	log.Println(result.Error)
}
fmt.Print(result.Stdout)
```

Для запуска команды из определенной директории используется `ExecuteInDirectory`:

```go
result := runner.ExecuteInDirectory("./workspace", "go", "test", "./...")
```

Команда запускается через `os/exec`; stdout и stderr собираются в строки после завершения процесса.