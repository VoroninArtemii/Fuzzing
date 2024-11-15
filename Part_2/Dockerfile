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

WORKDIR /apimachinery
COPY new/ pkg/test/new/
WORKDIR /apimachinery/pkg/test/new/
VOLUME /apimachinery/pkg/test/new/
RUN go test -fuzz=Fuzz -fuzztime=60s -parallel=1