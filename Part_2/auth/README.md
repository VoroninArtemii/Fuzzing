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
cd /podman/pkg/auth/
go test -fuzz=FuzzParseSingleAuthHeader -test.fuzzcachedir=./corpus -parallel=1
```

Запуск фаззинга _ParseMultiAuthHeader_:
```
cd /podman/pkg/auth/
go test -fuzz=FuzzParseMultiAuthHeader -test.fuzzcachedir=./corpus -parallel=1
```

Сбор покрытия _ParseSingleAuthHeader_ и _ParseMultiAuthHeader_:
```
cd /podman/pkg/auth/
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```
