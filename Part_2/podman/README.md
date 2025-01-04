Сборка образа:
```
docker build -t alt-podman .
```

Запуск контейнера:
```
docker run --name alt-podman -it alt-podman
```

Запуск фаззинга _ParseSingleAuthHeader_:
```
go test -fuzz=FuzzParseSingleAuthHeader -test.fuzzcachedir=./corpus -parallel=1
```

Запуск фаззинга _ParseMultiAuthHeader_:
```
go test -fuzz=FuzzParseMultiAuthHeader -test.fuzzcachedir=./corpus -parallel=1
```

Сбор покрытия _ParseSingleAuthHeader_ и _ParseMultiAuthHeader_:
```
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```
