FROM golang:1.13.4-alpine3.10 as golang

VOLUME [ "/app" ]

# Environment
ARG environment=dev
ENV ENV=${environment}

# Installing git
RUN apk add --no-cache git

# Install watcher for development environment
RUN if [ ${ENV} = "dev" ] ; then \
      go get github.com/canthefason/go-watcher && \
      go install github.com/canthefason/go-watcher/cmd/watcher; \
      # apk add --update alpine-sdk; \
    fi
# Make go test -race work on alpine by building (patched) sanitizer manually
# as it is not built by default
# Ref: https://github.com/golang/go/issues/14481#issuecomment-281972886
# SHA: https://github.com/golang/go/blob/go1.13/src/runtime/race/README
# COPY ./build/go/0001-hack-to-make-Go-s-race-flag-work-on-Alpine.patch /race.patch;
# RUN cd / \
#   && apk add --no-cache --virtual .build-deps g++ git \
#   && mkdir -p compiler-rt \
#   && git clone https://llvm.org/git/compiler-rt.git \
#   && cd compiler-rt \
#   && git reset --hard fe2c72c59aa7f4afa45e3f65a5d16a374b6cce26 \
#   && patch -p1 -i /race.patch \
#   && cd lib/tsan/go \
#   && ./buildgo.sh 2>/dev/null \
#   && cp -v race_linux_amd64.syso /usr/local/go/src/runtime/race/ \
#   && rm -rf /compiler-rt /race.patch \
#   && apk del .build-deps;

# SERVER HOST
ENV API_ORIGIN 0.0.0.0:81
EXPOSE 81

# Working directory setup
WORKDIR /app

COPY ./main.go ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./vendor ./vendor
COPY ./pkg ./pkg
COPY ./go.mod ./
COPY ./go.sum ./

# Run app
CMD [ "go", "run", "main.go" ]