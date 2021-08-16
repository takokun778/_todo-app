
FROM alpine:3.14

ARG UID=${UID}

ARG GID=${GID}

ARG UNAME=${UNAME}

ARG VERSION=${VERSION}

RUN apk update && apk add curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v${VERSION}/migrate.linux-amd64.tar.gz | tar xvz -C ./bin

RUN mv ./bin/migrate.linux-amd64 ./bin/migrate

RUN addgroup -g ${GID} -S ${UNAME} && adduser -u ${UID} -S ${UNAME} -G ${UNAME}

USER ${UID}

WORKDIR /home/${UNAME}/migration/