FROM sabayon/base-amd64:latest

ENV ACCEPT_LICENSE=*

RUN equo up && equo install enman && \
    enman add https://dispatcher.sabayon.org/sbi/namespace/devel/devel && \
    equo up && equo u && equo i mottainai-agent && equo cleanup

VOLUME ["/etc/mottainai", "/srv/mottainai"]

ENTRYPOINT [ "/usr/bin/mottainai-agent" ]
