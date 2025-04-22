Сборка образа:
```
docker build -t parse .
```

Запуск контейнера:
```
docker run --name parse -it parse
```

Запуск фаззинга _ParseImageName_:
```
cd /kubernetes/pkg/util/parsers
go test -fuzz=FuzzParseImageName -test.fuzzcachedir=./corpus -parallel=1
```

Сбор покрытия _ParseImageName_:
```
cd /kubernetes/pkg/util/parsers
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```