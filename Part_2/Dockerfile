FROM registry.altlinux.org/alt/alt:c10f2

RUN apt-get update -y
RUN apt-get dist-upgrade -y
RUN apt-get install curl git golang -y
RUN git clone --branch=release-1.31 --depth=10 https://github.com/kubernetes/kubernetes.git
WORKDIR /kubernetes
RUN git switch --detach v1.31.2
#RUN go mod vendor
#RUN make