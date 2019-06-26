FROM sabayon/base-amd64:latest

ENV ACCEPT_LICENSE=*

RUN equo up && equo install enman && \
    enman add https://dispatcher.sabayon.org/sbi/namespace/devel/devel && \
    equo up && equo u && equo i mottainai-agent && equo cleanup

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && chmod +x kubectl && mv kubectl /usr/bin

VOLUME ["/etc/mottainai", "/srv/mottainai"]

ENTRYPOINT [ "/usr/bin/mottainai-agent" ]
