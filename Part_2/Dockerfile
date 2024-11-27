FROM registry.altlinux.org/alt/alt:c10f2

RUN apt-get update -y
RUN apt-get dist-upgrade -y
RUN apt-get install git golang make nano -y
RUN git config --global http.postBuffer 524288000
RUN git clone https://github.com/kubernetes/apimachinery.git
RUN echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64 classic" > /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64-i586 classic" >> /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/noarch classic" >> /etc/apt/sources.list.d/altsp.list
RUN apt-get update -y
RUN apt-get install golang -y

COPY fuzz/fuzz_test.go /apimachinery/pkg/runtime/serializer/protobuf/fuzz_test.go
WORKDIR /apimachinery/pkg/runtime/serializer/protobuf/
VOLUME /apimachinery/pkg/runtime/serializer/protobuf/

RUN go test -fuzz=Fuzz -test.fuzzcachedir=./corpus -fuzztime=180s -parallel=1
RUN go test -test.fuzzcachedir=./corpus -coverprofile=cover.out
RUN go tool cover -html=cover.out -o cover.html