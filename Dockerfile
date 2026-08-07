© Copyright Hewlett Packard Enterprise Development LP

# builder image
FROM --platform=$BUILDPLATFORM registry.access.redhat.com/ubi9/ubi:9.8-1785807559 AS build
# install prereqs
RUN dnf install -y make golang
ENV GOVERSION="go1.25.0"
RUN go install golang.org/dl/$GOVERSION@latest
ENV PATH="/root/go/bin:$PATH"
RUN $GOVERSION download
ENV GOROOT="/root/sdk"
ENV GOPATH="/root/gocache/go"
RUN ln -s /root/go/bin/$GOVERSION /root/go/bin/go

# build binary
WORKDIR /usr/src/hpe-cosi-driver
ADD main.go .
ADD go.sum .
ADD go.mod .
ADD Makefile .
COPY servers /usr/src/hpe-cosi-driver/servers
COPY alletraMPX10000_sdk /usr/src/hpe-cosi-driver/alletraMPX10000_sdk
COPY glcp_tasks_sdk /usr/src/hpe-cosi-driver/glcp_tasks_sdk
COPY iam /usr/src/hpe-cosi-driver/iam
ARG TARGETARCH
RUN make clean
RUN make ARCH=$TARGETARCH build
ARG RUN_TESTS
RUN if [ "$RUN_TESTS" = "true" ]; then make test; fi

FROM registry.access.redhat.com/ubi9/ubi-micro:9.8-1784702951

LABEL name="HPE COSI Driver for Kubernetes" \
    maintainer="HPE Storage" \
    vendor="HPE" \
    version="2.0.0" \
    summary="HPE COSI Driver for Kubernetes" \
    description="The HPE COSI Driver for Kubernetes enables container orchestrators to manage the life-cycle of object storage resources." \
    io.k8s.display-name="HPE COSI Driver for Kubernetes" \
    io.k8s.description="The HPE COSI Driver for Kubernetes enables container orchestrators to manage the life-cycle of object storage resources."

COPY LICENSE /licenses/

# Copy CA bundle from builder stage (ubi-micro does not include ca-certificates)
COPY --from=build /etc/pki/tls/certs/ca-bundle.crt /etc/pki/tls/certs/ca-bundle.crt

RUN echo "cosi-driver:x:1000:1000::/home/cosi-driver:/sbin/nologin" >> /etc/passwd && \
    echo "cosi-driver:x:1000:" >> /etc/group

COPY --from=build /usr/src/hpe-cosi-driver/bin/cosi-driver cosi-driver

USER 1000:1000

ENTRYPOINT ["/cosi-driver"]
