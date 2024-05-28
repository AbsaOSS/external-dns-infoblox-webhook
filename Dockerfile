FROM gcr.io/distroless/static-debian12:nonroot

USER 20000:20000
COPY --chmod=555 external-dns-infoblox-webhook /opt/external-dns-infoblox-webhook/app

ENTRYPOINT ["/opt/external-dns-infoblox-webhook/app"]
