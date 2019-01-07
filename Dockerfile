############################
# STEP 1 build executable binary
############################

FROM golang:1.11-alpine as builder
COPY . /go/src/flaviojmendes/triangle-determiner
ENV GO111MODULE=on
WORKDIR /go/src/flaviojmendes/triangle-determiner
RUN apk -U add git build-base && \
    mkdir -p /build && \
    go build  -ldflags '-extldflags "-static"' -o /build/triangle-determiner


############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /build/triangle-determiner /build/triangle-determiner
# Run the hello binary.
ENTRYPOINT ["/build/triangle-determiner"]
