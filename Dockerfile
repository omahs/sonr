
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"

# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Sonrd Builder                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux golang:1.19-alpine AS builder
ARG arch=x86_64


ENV SONR_VERSION=master

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

# Download go dependencies
WORKDIR /root
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download


# Cosmwasm - Download correct libwasmvm version
RUN set -eux; \
    export ARCH=$(uname -m); \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}'); \
    if [ ! -z "${WASM_VERSION}" ]; then \
      wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a; \
    fi; \
    go mod download;

# Copy the remaining files
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
        -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="sonr" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="sonrd" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o /root/sonr/build/sonrd ./cmd/sonrd/main.go

# ! ||--------------------------------------------------------------------------------||
# ! ||                               Sonr in Production                               ||
# ! ||--------------------------------------------------------------------------------||
FROM ${RUNNER_IMAGE}

# Copy sonrd binary and config
COPY --from=builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.docker.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8000
