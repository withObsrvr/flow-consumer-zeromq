{
  description = "Obsrvr Flow Plugin: ZeroMQ Consumer";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages = {
          default = pkgs.buildGoModule {
            pname = "flow-consumer-zeromq";
            version = "0.1.0";
            src = ./.;
            
            # This is the calculated hash for the dependencies
            vendorHash = "sha256-sX8nlfZWzH/Wovn3J2hcQR/GJ3BXBpbCL/1bNWUs30Q=";
            
            # ZeroMQ requires CGO
            env = {
              CGO_ENABLED = "1";
              # Setting proxy to direct allows module downloads
              GOPROXY = "direct";
            };
            
            # Build as a shared library/plugin
            buildPhase = ''
              runHook preBuild
              go build -buildmode=plugin -o flow-consumer-zeromq.so .
              runHook postBuild
            '';

            # Custom install phase for the plugin
            installPhase = ''
              runHook preInstall
              mkdir -p $out/lib
              cp flow-consumer-zeromq.so $out/lib/
              # Also install a copy of go.mod for future reference
              mkdir -p $out/share
              cp go.mod $out/share/
              # Copy go.sum if it exists
              if [ -f go.sum ]; then
                cp go.sum $out/share/
              fi
              runHook postInstall
            '';
            
            # Add ZeroMQ dependencies
            nativeBuildInputs = [ pkgs.pkg-config ];
            buildInputs = [ 
              pkgs.zeromq
              pkgs.czmq
              pkgs.libsodium
            ];
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ 
            # Match Go version from go.mod (1.23.4)
            go_1_23
            # ZeroMQ dependencies
            pkg-config
            zeromq
            czmq
            libsodium
            # Development tools
            gopls
            delve
          ];
          
          # Enable CGO for ZeroMQ
          env = {
            CGO_ENABLED = "1";
          };
        };
      }
    );
} 