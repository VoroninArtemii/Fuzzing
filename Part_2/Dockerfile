FROM registry.altlinux.org/alt/alt:c10f2

RUN apt-get update -y
RUN apt-get dist-upgrade -y
RUN apt-get install git golang make -y
RUN git clone --branch=release-1.31 --depth=10 https://github.com/kubernetes/kubernetes.git

WORKDIR /kubernetes
RUN git switch --detach v1.31.2
RUN make

COPY new/ test/fuzz/new/
WORKDIR /kubernetes/test/fuzz/new/
VOLUME /kubernetes/test/fuzz/new/
RUN go test -fuzz=Fuzz -fuzztime=7200s -parallel=1