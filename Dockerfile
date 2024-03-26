FROM golang:1.22.0-bookworm as build
LABEL MAINTAINER="shiny.gift6738@fastmail.com"

WORKDIR /app

RUN apt-get update && apt-get -y upgrade
# ADD file changes when new commits have been made - this forces a new checkout instead of using cache
ADD https://api.github.com/repos/skeletonkey/coldStorage/git/refs/heads/$BRANCH version.json
RUN git clone -b initialApp --single-branch --depth 1 https://github.com/skeletonkey/coldStorage.git coldStorage
RUN cd coldStorage && CGO_ENABLED=1 go build -o bin/cold-storage app/*.go

FROM golang:1.22.0-bookworm as prod
LABEL MAINTAINER="shiny.gift6738@fastmail.com"

ENV RACHIO_CONFIG_FILE="/app/config.json"

WORKDIR /app

COPY config.json .
COPY --from=build /app/coldStorage/bin/cold-storage cold-storage

RUN chmod 0755 /app/cold-storage

# ENTRYPOINT [ "sleep 100000" ]
ENTRYPOINT [ "/app/cold-storage" ]