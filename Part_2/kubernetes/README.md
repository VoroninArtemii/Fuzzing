Сборка образа:
```
docker build -t alt-kubernetes .
```

Запуск контейнера:
```
docker run --name alt-kubernetes -it alt-kubernetes
```

Запуск фаззинга _ParseImageName_:
```
cd /kubernetes/pkg/util/parsers
go test -fuzz=Fuzz -test.fuzzcachedir=./corpus -parallel=1
```

Сбор покрытия _ParseImageName_:
```
cd /kubernetes/pkg/util/parsers
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```

Запуск фаззинга _Decode_:
```
cd /apimachinery/pkg/runtime/serializer/protobuf
go test -fuzz=Fuzz -test.fuzzcachedir=./corpus -parallel=1
```

Сбор покрытия _Decode_:
```
cd /apimachinery/pkg/runtime/serializer/protobuf
go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```