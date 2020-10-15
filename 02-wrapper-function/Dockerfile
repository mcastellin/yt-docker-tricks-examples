FROM debian:buster-slim
LABEL maintainer="Manuel Castellin"

# installing UUID generator utility
RUN apt-get update \
    && apt-get install --no-install-recommends -qqy \
    uuid-runtime \
    gosu \
    && rm -rf /var/lib/apt/lists

WORKDIR /myapp

# create a non-root user to run the script generator
RUN useradd -l -M appuser
USER appuser

# copy scripts and set exec permissions
COPY --chown=appuser *.sh ./
RUN chmod +x *.sh

VOLUME [ "/data" ]

ENTRYPOINT ["./entrypoint.sh"]
CMD ["./generator.sh"]
