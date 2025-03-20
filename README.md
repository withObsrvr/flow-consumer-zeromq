# Flow Consumer: ZeroMQ

This plugin provides a ZeroMQ consumer for Obsrvr Flow. It allows the Flow system to consume data from ZeroMQ sockets.

## Building with Nix

This project uses Nix for reproducible builds.

### Prerequisites

- [Nix package manager](https://nixos.org/download.html) with flakes enabled

### Building

1. Clone the repository:
```
git clone https://github.com/withObsrvr/flow-consumer-zeromq.git
cd flow-consumer-zeromq
```

2. Build with Nix:
```
nix build
```

The built plugin will be available at `./result/lib/flow-consumer-zeromq.so`.

### Development

To enter a development shell with all dependencies:

```
nix develop
```

This will provide a shell with all the necessary dependencies, including:
- Go 1.23
- ZeroMQ, CZMQ, and libsodium libraries
- Development tools (gopls, delve)

## Dependencies

This plugin requires the following dependencies, which are automatically handled by the Nix configuration:

- ZeroMQ (libzmq)
- CZMQ (libczmq)
- libsodium

All these dependencies are managed through the `flake.nix` file, so you don't need to install them separately when using Nix.

## First Time Setup Notes

When setting up this project with Nix for the first time, you might encounter the following:

1. Initial vendorHash error - This is expected and has been resolved by adding the correct hash to the flake.nix file.
2. CGO is enabled by default for ZeroMQ integration.
3. The build uses Go's module system for dependencies.

## Plugin Details

This ZeroMQ consumer plugin connects to ZeroMQ sockets to consume data for processing in the Flow system. It requires CGO enabled due to the ZeroMQ C library dependencies. 