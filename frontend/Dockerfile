FROM ubuntu:latest

RUN apt update && \
    apt install python3-venv -y && \
    apt install python3-pip -y && \
    pip3 install --upgrade pip && \
    pip3 install --upgrade virtualenv && \
    apt install nginx -y && \
    apt install vim -y && \
    apt install curl -y && \
    apt install npm -y && \
    apt install ufw -y && \
    ufw allow 'Nginx HTTP' && \
    apt install systemctl -y && \
    systemctl start nginx && \
    systemctl enable nginx 

WORKDIR /

ENV ENV "dev"

EXPOSE 80

CMD ["tail", "-f", "/dev/null"]
