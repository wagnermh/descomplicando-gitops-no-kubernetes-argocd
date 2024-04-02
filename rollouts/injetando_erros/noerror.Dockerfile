FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static" -X main.shouldBeGeneratingError=false' -tags timetzdata -o /app/app

FROM scratch
# the test program:
COPY --from=app-builder /app/app /app
# the tls certificates:
# NB: this pulls directly from the upstream image, which already has ca-certificates:
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 3000
ENTRYPOINT ["/app"]