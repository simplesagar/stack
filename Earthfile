VERSION 0.7

ARG --global ALPINE_VERSION=3.17
ARG --global VERSION=develop
ARG --global GOOS=linux
ARG --global GOARCH=arm64
ARG --global BUILD_DATE=-
ARG --global REPOSITORY=ghcr.io
ARG --global DOCKER_VERSION=23.0.1
ARG --global GOLANG_VERSION=1.20
ARG --global NODE_VERSION=19
ARG --global GOLANGCI_LINT_VERSION=v1.51.2
ARG --global OPENAPI_GENERATOR_VERSION=v6.4.0

FROM alpine:${ALPINE_VERSION}
WORKDIR /src

DOWNLOAD_DEPENDENCIES:
    COMMAND
    ARG DEPENDENCY_PATH
    WORKDIR /src/$DEPENDENCY_PATH
    COPY (+sources/$DEPENDENCY_PATH/go.* --SOURCE_PATH=$DEPENDENCY_PATH) .
    ENV GOPROXY http://172.26.0.12:8080
    ENV GOSUMDB=off
    RUN go mod download -x
    WORKDIR /src

LOAD_SOURCES:
    COMMAND
    ARG DEPENDENCY_PATH
    COPY (+sources/$DEPENDENCY_PATH/ --SOURCE_PATH=$DEPENDENCY_PATH) /src/$DEPENDENCY_PATH/

LOAD_SOURCE_FILE:
    COMMAND
    ARG DEPENDENCY_PATH
    COPY (+sources/$DEPENDENCY_PATH --SOURCE_PATH=$DEPENDENCY_PATH) /src/$DEPENDENCY_PATH

SAVE_IMAGE:
    COMMAND
    ARG COMPONENT
    SAVE IMAGE ${REPOSITORY}/formancehq/${COMPONENT}:${VERSION}

STD_BUILD:
    COMMAND
    ARG COMPONENT
    ARG EARTHLY_GIT_HASH
    ARG CGO_ENABLED=0
    RUN go build -o $COMPONENT \
        -ldflags="-X github.com/formancehq/$COMPONENT/cmd.Version=${VERSION} \
        -X github.com/formancehq/$COMPONENT/cmd.BuildDate=$BUILD_DATE \
        -X github.com/formancehq/$COMPONENT/cmd.Commit=$EARTHLY_GIT_HASH" ./
    SAVE ARTIFACT ./$COMPONENT

GO_STD_TEST_COMMAND:
    COMMAND
    RUN go test -coverpkg ./... -coverprofile coverage.out -covermode atomic ./...
    SAVE ARTIFACT ./coverage.out AS LOCAL ./coverage.out

TESTS:
    COMMAND
    ARG LOCATION
    ARG CGO_ENABLED=0
    ENV GOPROXY http://172.26.0.5:8080
    ENV GOSUMDB=off
    DO +GO_STD_TEST_COMMAND

LINT:
    COMMAND
    DO +LOAD_SOURCE_FILE --DEPENDENCY_PATH=.golangci.yml
    RUN go mod tidy
    RUN --mount=type=cache,target=/root/.cache/go-build \
        golangci-lint -v run --fix --timeout=10m --verbose -c /src/.golangci.yml
    SAVE ARTIFACT ./* AS LOCAL ./

# Share common images to avoid version drift
benthos:
    FROM jeffail/benthos:4.11

postgres:
    FROM postgres:15-alpine

base-build-image:
    FROM docker:23.0.1-dind-alpine3.17
    RUN apk update && apk add --virtual build-dependencies gcc musl-dev jq go nodejs make bash curl
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /bin ${GOLANGCI_LINT_VERSION}
    ARG GOPROXY
    ARG GOSUMDB

    WORKDIR /src

base-final-image:
    RUN apk update && apk add ca-certificates curl
    ARG COMPONENT
    ENV OTEL_SERVICE_NAME $COMPONENT

sources:
    ARG --required SOURCE_PATH
    COPY $SOURCE_PATH /src/$SOURCE_PATH
    SAVE ARTIFACT /src/$SOURCE_PATH $SOURCE_PATH

build-image:
    ARG --required COMPONENT
    BUILD ./components/$COMPONENT+build-image

tests:
    ARG --required LOCATION
    BUILD ./$LOCATION+tests

lint:
    ARG --required LOCATION
    BUILD ./$LOCATION+lint

build-all:
    DO +LOAD_SOURCES --DEPENDENCY_PATH=components
    FOR component IN $(ls components)
        BUILD +build-image --COMPONENT=$component
    END

tests-all:
    BUILD +tests --LOCATION=libs/go-libs
    DO +LOAD_SOURCES --DEPENDENCY_PATH=components
    FOR component IN $(ls components)
        BUILD +tests --LOCATION=components/$component
    END

lint-all:
    BUILD +lint --LOCATION=libs/go-libs
    DO +LOAD_SOURCES --DEPENDENCY_PATH=components
    FOR component IN $(ls components)
        BUILD +lint --LOCATION=components/$component
    END

sdk-generation:
    FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION}
    WORKDIR /src/openapi
    COPY openapi/package.* .
    RUN npm install
    SAVE ARTIFACT ./package.*
    SAVE ARTIFACT ./node_modules

build-openapi-spec:
    FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION}
    WORKDIR /src
    COPY . .
    COPY (+sdk-generation/*) ./openapi
    WORKDIR /src/openapi
    RUN npm run build
    RUN sed -i -e "s/SDK_VERSION/$VERSION/g" build/generate.json
    SAVE ARTIFACT ./build/generate.json

generate-sdk:
    ARG --required LANG
    FROM openapitools/openapi-generator-cli:${OPENAPI_GENERATION_VERSION}
    WORKDIR /src
    COPY components components
    COPY openapi openapi
    COPY (+build-openapi-spec/generate.json) ./openapi/build/generate.json
    RUN docker-entrypoint.sh generate \
        -i ./openapi/build/generate.json \
        -g $(echo $LANG | cut -d - -f1) \
        -c ./openapi/configs/$LANG.yaml \
        -o ./sdks/$LANG \
        --git-user-id=formancehq \
        --git-repo-id=formance-sdk-$LANG \
        -p packageVersion=$VERSION \
        -p apiVersion=$VERSION
    SAVE ARTIFACT ./sdks/$LANG AS LOCAL ./sdks/$LANG

generate-all-sdk:
    BUILD +generate-sdk --LANG=go
    BUILD +generate-sdk --LANG=typescript-node
    BUILD +generate-sdk --LANG=python
    BUILD +generate-sdk --LANG=java
    BUILD +generate-sdk --LANG=php

tests-integrations:
    FROM +base-build-image
    DO +LOAD_SOURCES --DEPENDENCY_PATH=sdks/go
    DO +LOAD_SOURCES --DEPENDENCY_PATH=libs/go-libs
    DO +LOAD_SOURCES --DEPENDENCY_PATH=libs/events
    DO +LOAD_SOURCES --DEPENDENCY_PATH=components
    DO +LOAD_SOURCE_FILE --DEPENDENCY_PATH=go.mod
    DO +LOAD_SOURCE_FILE --DEPENDENCY_PATH=go.sum
    RUN go mod download -x
    DO +DOWNLOAD_DEPENDENCIES --DEPENDENCY_PATH=tests/integration
    DO +LOAD_SOURCES --DEPENDENCY_PATH=tests/integration

    WORKDIR /src/tests/integration
    # Keep here and don't use a fixed version. Ginkgo is using go.mod version.
    RUN go install github.com/onsi/ginkgo/v2/ginkgo
    RUN cp $(go env GOPATH)/bin/ginkgo /bin

    ARG STRIPE_API_KEY
    WITH DOCKER --compose docker-compose.yml \
        --allow-privileged \
        --load jeffail/benthos:4.11=+benthos
        RUN export DOCKER_HOSTNAME=$(ip addr show eth0 | grep brd | grep inet | cut -d \  -f 6 | cut -d / -f 1) && \
            echo "Found gateway IP: $DOCKER_HOSTNAME" && \
            ginkgo -v -p ./suite
    END

all:
    BUILD +build-all
    BUILD +lint-all
    BUILD +tests-all
    BUILD +tests-integrations
