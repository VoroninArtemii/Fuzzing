Сборка образа:
```
docker build -t decode .
```

Запуск контейнера:
```
docker run --name decode -it decode
```

Запуск фаззинга _Decode_:
```
go test -fuzz=FuzzDecode -test.fuzzcachedir=./corpus -parallel=8
```

Сбор покрытия _Decode_:
```
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```