FROM gcr.io/distroless/static-debian12:nonroot@sha256:a9329520abc449e3b14d5bc3a6ffae065bdde0f02667fa10880c49b35c109fd1
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
