FROM registry.altlinux.org/alt/alt:c10f2

RUN apt-get update -y
RUN apt-get dist-upgrade -y
RUN apt-get install curl git golang -y
RUN git clone git://git.altlinux.org/gears/k/kubernetes1.31.git
RUN cd kubernetes1.31
RUN git checkout v1.31.1 && go mod vendor
RUN make