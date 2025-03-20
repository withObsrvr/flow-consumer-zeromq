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
- ZeroMQ and CZMQ libraries
- Development tools (gopls, delve)

## First Time Build Notes

When building for the first time, the initial build will likely fail with a vendorHash error. Update the vendorHash in flake.nix with the hash from the error message, and then run the build again.

## Plugin Details

This ZeroMQ consumer plugin connects to ZeroMQ sockets to consume data for processing in the Flow system. It requires CGO enabled due to the ZeroMQ C library dependencies. 