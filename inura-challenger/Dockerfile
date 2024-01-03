ARG OP_STACK_GO_BUILDER=us-docker.pkg.dev/oplabs-tools-artifacts/images/op-stack-go:latest
FROM $OP_STACK_GO_BUILDER as builder
# See "make golang-docker" and /ops/docker/op-stack-go

FROM alpine:3.18

# Make the bundled inura-program the default cannon server
COPY --from=builder /usr/local/bin/inura-program /usr/local/bin/inura-program
ENV OP_CHALLENGER_CANNON_SERVER /usr/local/bin/inura-program

# Make the bundled cannon the default cannon executable
COPY --from=builder /usr/local/bin/cannon /usr/local/bin/cannon
ENV OP_CHALLENGER_CANNON_BIN /usr/local/bin/cannon

COPY --from=builder /usr/local/bin/inura-challenger /usr/local/bin/inura-challenger

CMD ["inura-challenger"]
