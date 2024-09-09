FROM ubuntu:latest

ENV TZ=Europe/Moscow \
    DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt install python3-venv -y && \
    apt install python3-pip -y && \
    pip3 install --upgrade pip && \
    pip3 install --upgrade virtualenv && \
    apt install postgresql -y && \
    service postgresql start && \ 
    apt install wget -y && \
    wget -c https://go.dev/dl/go1.21.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz  && \
    apt install vim -y && \
    apt install curl -y 
    

ENV PATH "$PATH:/usr/local/go/bin"

ENV ENV "dev"

WORKDIR /

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]
