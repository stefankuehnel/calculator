{
  description = "A Nix-flake-based development environment";

  inputs = {
    nixpkgs = {
      url = "https://nixos.org/channels/nixpkgs-unstable/nixexprs.tar.xz";
    };
  };

  outputs =
    {
      nixpkgs,
      ...
    }:
    let
      supportedSystems = nixpkgs.lib.systems.flakeExposed;

      forAllSystems =
        function:
        nixpkgs.lib.genAttrs supportedSystems (
          system:
          function {
            pkgs = import nixpkgs {
              inherit system;

              config = {
                allowUnfree = true;
              };

              overlays = [
                (final: previous: {
                  localPackages = {
                    calculator = final.callPackage ./calculator.nix { };
                    calculator_docker = final.callPackage ./calculator-docker.nix { };
                  };
                })
              ];
            };
            inherit system;
          }
        );

      dependencies = pkgs: [
        pkgs.antlr4

        pkgs.git

        pkgs.go
        pkgs.go-task
        pkgs.golangci-lint

        pkgs.nix
      ];

      devDependencies = pkgs: [
        pkgs.opencode
        pkgs.claude-code

        pkgs.docker

        pkgs.coreutils

        pkgs.nixd
        pkgs.nixfmt
      ];

      ciDependencies = pkgs: [
      ];

      cdDependencies = pkgs: [
        pkgs.semver-tool

        pkgs.skopeo
      ];

      fmtDependencies = pkgs: pkgs.nixfmt-tree;
    in
    {
      formatter = forAllSystems ({ pkgs, ... }: fmtDependencies pkgs);

      devShells = forAllSystems (
        { pkgs, ... }:
        {
          ciEnvironment = pkgs.mkShellNoCC {
            packages = (dependencies pkgs) ++ (ciDependencies pkgs);
          };

          cdEnvironment = pkgs.mkShellNoCC {
            packages = (dependencies pkgs) ++ (cdDependencies pkgs);
          };

          devEnvironment = pkgs.mkShellNoCC {
            packages =
              (dependencies pkgs) ++ (ciDependencies pkgs) ++ (cdDependencies pkgs) ++ (devDependencies pkgs);
          };
        }
      );

      packages = forAllSystems (
        { pkgs, ... }:
        {
          default = pkgs.localPackages.calculator;

          calculator = pkgs.localPackages.calculator;
          calculator-docker = pkgs.localPackages.calculator_docker;

          devEnvironment = pkgs.buildEnv {
            name = "development environment";
            paths =
              (dependencies pkgs) ++ (ciDependencies pkgs) ++ (cdDependencies pkgs) ++ (devDependencies pkgs);
          };
        }
      );
    };
}
