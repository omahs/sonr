

<div style="text-align: center;">

[![Sonr Logo Banner](docs/static/images/gh-banner.png)](https://snr.la/h)

[![Go Reference](https://pkg.go.dev/badge/github.com/sonr-hq/sonr.svg)](https://pkg.go.dev/github.com/sonr-hq/sonr)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonr-hq/sonr)](https://goreportcard.com/report/github.com/sonr-hq/sonr)
[![License](https://img.shields.io/github/license/sonr-hq/sonr)](https://github.com/sonr-hq/sonr)

</div>


<p align="center"> Sonr is a <strong>peer-to-peer identity</strong> and <strong>asset management system</strong> that leverages <italic>DID Documents, WebAuthn, and IPFS</italic> - to provide users with a <strong>secure, user-friendly</strong> way to manage their <strong>digital identity and assets.</strong>
    <br>
</p>


</br>

## Getting Started

#### Prerequisites
- Cosmos SDK: v0.46.7
- Ignite CLI: v0.25.2
- Golang: 1.18.10 darwin/arm64
- Taskfile v3.20.0

Setting up your local development environment:

```sh
# Clone the repository
git clone https://github.com/sonr-hq/sonr.git

# Install dependencies
sh scripts/install.sh

# Display all available tasks
task
```

## Documentation

Sonr utilizes Mintlify to generate documentation from the source code. To run the documentation server, execute `task docs` from the root directory. Or, visit the [documentation site](https://snr.la/docs).

<details>
<summary><strong>Architecture Diagram</strong></summary>

![Architecture](./docs/static/images/diagrams/architecture-light.svg)

</details>

<details>
<summary><strong>Repository Structure Map</strong></summary>

![Repository structure](./docs/static/images/diagrams/repo-structure.svg)

</details>

## Usage

### Project Commands
`task serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

<details>
<summary><strong>Command Reference</strong></summary>

```sh
* docs:                 Serve the docs locally
* chain:build:          Build the blockchain                  (aliases: build)
* chain:generate:       Generate the protobuf files           (aliases: gen)
* chain:serve:          Serve the blockchain locally          (aliases: serve)
* motor:android:        Bind the Motor Android Framework      (aliases: android)
* motor:ios:            Bind the Motor iOS Framework          (aliases: ios)
* motor:web:            Build the Motor WASM Framework        (aliases: wasm)
* web:dev:              Run the web app in dev mode           (aliases: web)
```
</details>
<details>
<summary><strong>Publish New Release</strong></summary>

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```sh
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

</details>

### Installation
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```sh
curl https://get.ignite.com/sonr-hq/sonr! | sudo bash
```
Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Community

For help and discussion around DID Documents, best practices, and more, join us on [Discord](https://snr.la/dcrd).

For updates on the Sonr Daemon, [follow this repo on GitHub](https://github.com/sonr-hq/sonr).

For feature requests, bugs, or technical questions, [submit feedback here](https://roadmap.sonr.io). For general inquiries or inclusion in our upcoming feature betas, reach out on [Twitter](https://snr.la/tw).
