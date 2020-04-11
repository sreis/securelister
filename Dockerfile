FROM gcr.io/distroless/static@sha256:08322afd57db6c2fd7a4264bf0edd9913176835585493144ee9ffe0c8b576a76

WORKDIR /

COPY bin/apiserver /apiserver
COPY bin/controller-manager /controller-manager

ENTRYPOINT [ "/apiserver" ]
