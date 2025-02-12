{
  description = "Breakfast productivity app suite";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
        inherit system;
        config.allowUnfree = true;
      };
    in {
      devShell = pkgs.mkShell {
        buildInputs = with pkgs; [
          mysql
          postman
          nodejs
          k6
          go
        ];
      };
      shellHook = ''
        export LD_LIBRARY_PATH=${pkgs.lib.makeLibraryPath [ pkgs.libglvnd ]}:$LD_LIBRARY_PATH
        echo "Welcome to the development environment!"
        USER_SHELL=$(getent passwd $USER | cut -d: -f7)
        exec $USER_SHELL
      '';
    }
  );
}

