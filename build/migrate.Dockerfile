FROM alpine:3.22.2
WORKDIR /backend
RUN apk add --no-cache curl \
  && curl -sSf https://atlasgo.sh | sh
COPY migration/ migration/
COPY atlas.hcl atlas.hcl
ENTRYPOINT ["atlas"]
