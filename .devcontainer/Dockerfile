# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.238.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends redis

# https://github.com/dominikh/go-tools
USER vscode
RUN curl -L -O https://github.com/dominikh/go-tools/releases/download/v0.3.3/staticcheck_linux_arm64.tar.gz
RUN tar -xzf staticcheck_linux_arm64.tar.gz && mv staticcheck/staticcheck $(go env GOPATH)/bin
RUN rm -rf staticcheck_linux_arm64.tar.gz staticcheck

# https://github.com/vektra/mockery
USER vscode
RUN curl -L -O https://github.com/vektra/mockery/releases/download/v2.14.0/mockery_2.14.0_Linux_arm64.tar.gz
RUN tar -xzf mockery_2.14.0_Linux_arm64.tar.gz && mv mockery $(go env GOPATH)/bin
RUN rm -rf mockery_2.14.0_Linux_arm64.tar.gz

# https://github.com/cespare/reflex
USER vscode
RUN curl -L -O https://github.com/cespare/reflex/releases/download/v0.3.1/reflex_linux_arm64.tar.gz
RUN tar -xzf reflex_linux_arm64.tar.gz && mv reflex_linux_arm64/reflex $(go env GOPATH)/bin
RUN rm -rf reflex_linux_arm64.tar.gz reflex_linux_arm64