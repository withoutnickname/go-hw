## Запуск тестов

Для запуска тестов необходимо выполнить команду `go test .`

Если тесты содержат benchmarks, можно выполнить их с помощью флагов `--bench` и `--benchmem`:

    go test . -v --bench . --benchmem

Обратите внимание, что benchmarks тесты будут запускаться на разных машинах с различной конфигурацией и результаты иногда могут отличаться серьезным образом.