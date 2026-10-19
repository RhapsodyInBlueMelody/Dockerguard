{
  description = "Go Programming environtment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          gotools
          go-tools
        ];

        shellHook = ''
          echo "🦫 Go Programming Language Environtment"
          echo "Go version: $(go version)"
        '';
      };
    };
}
