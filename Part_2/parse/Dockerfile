FROM registry.altlinux.org/alt/alt:c10f2

RUN echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64 classic" > /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64-i586 classic" >> /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/noarch classic" >> /etc/apt/sources.list.d/altsp.list
RUN apt-get update -y
RUN apt-get install git golang -y
RUN git config --global http.postBuffer 524288000
RUN git clone --branch=v1.31.3 --depth=1 https://github.com/kubernetes/kubernetes.git
WORKDIR /kubernetes

COPY parse_test.go /kubernetes/pkg/util/parsers/parse_test.go
VOLUME /kubernetes/pkg/util/parsers/
WORKDIR /kubernetes/pkg/util/parsers/