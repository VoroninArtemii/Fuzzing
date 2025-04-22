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
go test -fuzz=FuzzParseImageName -test.fuzzcachedir=./corpus -parallel=8
```

Сбор покрытия _ParseImageName_:
```
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```