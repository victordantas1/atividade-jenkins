# Usa a imagem oficial do Jenkins LTS (Long Term Support)
FROM jenkins/jenkins:lts

# Muda para o usuário root para ter permissão de instalar pacotes
USER root

# Baixa e instala o Golang 1.22
RUN curl -OL https://go.dev/dl/go1.22.1.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz \
    && rm go1.22.1.linux-amd64.tar.gz

# Adiciona o Go ao PATH do container
ENV PATH=$PATH:/usr/local/go/bin

# Volta para o usuário jenkins por segurança
USER jenkins
