Сборка образа:
```
docker build -t auth .
```

Запуск контейнера:
```
docker run --name auth -it auth
```

Запуск фаззинга _ParseSingleAuthHeader_:
```
go test -fuzz=FuzzParseSingleAuthHeader -test.fuzzcachedir=./corpus -parallel=8
```

Запуск фаззинга _ParseMultiAuthHeader_:
```
go test -fuzz=FuzzParseMultiAuthHeader -test.fuzzcachedir=./corpus -parallel=8
```

Сбор покрытия _ParseSingleAuthHeader_ и _ParseMultiAuthHeader_:
```
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```