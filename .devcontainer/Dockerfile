FROM tinygo/tinygo:0.22.0

ENV GOPROXY=https://goproxy.cn,direct

ENV GOPRIVATE=codeup.aliyun.com

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list && sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN apt update && apt install -y git \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*
