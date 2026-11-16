FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab
ARG TARGETPLATFORM
LABEL org.opencontainers.image.title="twitch_exporter" \
    org.opencontainers.image.description="Export Twitch metrics to Prometheus" \
    org.opencontainers.image.source="https://github.com/esamarathon/twitch_exporter" \
    org.opencontainers.image.documentation="https://github.com/esamarathon/twitch_exporter/blob/main/README.md" \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.vendor="ESAMarathon"
COPY $TARGETPLATFORM/twitch_exporter /bin/twitch_exporter
EXPOSE 9184
USER nonroot
ENTRYPOINT ["/bin/twitch_exporter"]
