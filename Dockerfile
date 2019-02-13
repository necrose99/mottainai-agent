FROM sabayon/base-amd64:latest

ENV ACCEPT_LICENSE=*

RUN equo up && equo install enman && \
    enman add devel && \
    equo up && equo u && equo i mottainai-agent

VOLUME ["/etc/mottainai", "/srv/mottainai"]

ENTRYPOINT [ "/usr/bin/mottainai-agent" ]
