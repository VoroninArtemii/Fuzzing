FROM registry.altlinux.org/alt/alt:c10f2

RUN echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64 classic" > /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64-i586 classic" >> /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/noarch classic" >> /etc/apt/sources.list.d/altsp.list
RUN apt-get update -y
RUN apt-get install git golang -y
RUN git config --global http.postBuffer 524288000
RUN git clone --single-branch --branch=release-1.31 --depth=10 https://github.com/kubernetes/apimachinery.git
WORKDIR /apimachinery
RUN git switch --detach v0.31.2
WORKDIR /
RUN git clone --single-branch --branch=release-1.31 --depth=10 https://github.com/kubernetes/kubernetes.git
WORKDIR /kubernetes
RUN git switch --detach v1.31.2
WORKDIR /

COPY parse_test.go /kubernetes/pkg/util/parsers/parse_test.go
VOLUME /kubernetes/pkg/util/parsers/

COPY decode_test.go /apimachinery/pkg/runtime/serializer/protobuf/decode_test.go
VOLUME /apimachinery/pkg/runtime/serializer/protobuf/
