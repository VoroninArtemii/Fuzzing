FROM registry.altlinux.org/alt/alt:c10f2

RUN echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64 classic" > /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/x86_64-i586 classic" >> /etc/apt/sources.list.d/altsp.list && \
    echo "rpm http://ftp.altlinux.org/pub/distributions/ALTLinux c10f2/branch/noarch classic" >> /etc/apt/sources.list.d/altsp.list
RUN apt-get update -y
RUN apt-get install git golang -y
RUN git clone --single-branch --branch=v4.9 --depth=10 https://github.com/containers/podman
WORKDIR /podman
RUN git switch --detach v4.9.5

COPY auth_test.go /podman/pkg/auth/auth_test.go
VOLUME /podman/pkg/auth/
WORKDIR /podman/pkg/auth/
