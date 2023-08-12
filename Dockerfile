ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"

# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Cosmjs Faucet                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux node:18.7-alpine AS sonr-faucet

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

ENV COSMJS_VERSION=0.28.11

RUN npm install @cosmjs/faucet@${COSMJS_VERSION} --global --production

ENV FAUCET_CONCURRENCY=4
ENV FAUCET_PORT=4500
ENV FAUCET_GAS_PRICE=0.0000usnr
# Prepared keys for determinism
ENV FAUCET_MNEMONIC="decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"
ENV FAUCET_ADDRESS_PREFIX=idx
ENV FAUCET_TOKENS="usnr, snr"
ENV FAUCET_CREDIT_AMOUNT_STAKE=1000
ENV FAUCET_CREDIT_AMOUNT_TOKEN=100
ENV FAUCET_COOLDOWN_TIME=0

EXPOSE 4500

ENTRYPOINT [ "cosmos-faucet" ]

# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Sonrd Builder                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux golang:1.19-alpine AS sonr-builder
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



# ! ||-----------------------------------------------------------------------------||
# ! ||                               Sonr Base Image                               ||
# ! ||-----------------------------------------------------------------------------||
FROM alpine AS sonr-base

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Copy sonrd binary and config
COPY --from=sonr-builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080


# ! ||----------------------------------------------------------------------------------||
# ! ||                               Sonr Standalone Node                               ||
# ! ||----------------------------------------------------------------------------------||

FROM alpine AS sonr-node

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Copy sonrd binary and config
COPY --from=sonr-builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Download, extract, and install the toml-cli binary
RUN apk add --update curl
RUN curl -LO https://github.com/gnprice/toml-cli/releases/latest/download/toml-0.2.3-x86_64-linux.tar.gz && \
    tar -xvf toml-0.2.3-x86_64-linux.tar.gz && \
    mv toml-0.2.3-x86_64-linux/toml /usr/local/bin && \
    rm toml-0.2.3-x86_64-linux.tar.gz && \
    rm -rf toml-0.2.3-x86_64-linux

# Setup localnet environment
RUN sh scripts/localnet.sh

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090

CMD [ "sonrd", "start" ]


# ! ||--------------------------------------------------------------------------------||
# ! ||                               Sonr Operator Node                               ||
# ! ||--------------------------------------------------------------------------------||

FROM sonr-base AS operator

LABEL org.opencontainers.image.source https://github.com/sonrhq/core

# Download, extract, and install the toml-cli binary
RUN apk add --update curl
RUN curl -LO https://github.com/gnprice/toml-cli/releases/latest/download/toml-0.2.3-x86_64-linux.tar.gz && \
    tar -xvf toml-0.2.3-x86_64-linux.tar.gz && \
    mv toml-0.2.3-x86_64-linux/toml /usr/local/bin && \
    rm toml-0.2.3-x86_64-linux.tar.gz && \
    rm -rf toml-0.2.3-x86_64-linux

# Setup environment variables
ENV KEY="alice"
ENV CHAIN_ID=sonr-1
ENV MONIKER=florence
ENV KEYALGO=secp256k1
ENV KEYRING=test
ENV MNEMONIC="decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry"

# Initialize the node
RUN echo $MNEMONIC | sonrd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover
RUN sonrd init ${MONIKER} --chain-id ${CHAIN_ID} --home /root/.sonr
RUN sonrd add-genesis-account $KEY 100000000000000000000000000usnr,1000000000000000snr --keyring-backend $KEYRING
RUN sonrd gentx $KEY 1000000000000000000000usnr --keyring-backend $KEYRING --chain-id $CHAIN_ID
RUN sonrd collect-gentxs

# Update config.toml
RUN toml set $HOME/.sonr/config/config.toml rpc.laddr tcp://0.0.0.0:26657 > /tmp/config.toml && mv /tmp/config.toml $HOME/.sonr/config/config.toml
RUN toml set $HOME/.sonr/config/app.toml grpc.address 0.0.0.0:9000 > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
RUN toml set $HOME/.sonr/config/app.toml api.enable true > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
RUN toml set $HOME/.sonr/config/app.toml api.swagger true > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
RUN toml set $HOME/.sonr/config/app.toml api.address tcp://0.0.0.0:1317 > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml
RUN toml set $HOME/.sonr/config/app.toml minimum-gas-prices 0.0000snr > /tmp/app.toml && mv /tmp/app.toml $HOME/.sonr/config/app.toml


# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090

CMD [ "sonrd", "start" ]
