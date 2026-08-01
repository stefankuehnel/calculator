{
  lib,
  buildGoModule,
}:

buildGoModule rec {
  pname = "calculator";
  version = "0.1.0";

  # See: https://nix.dev/guides/best-practices#reproducible-source-paths
  src = builtins.path {
    path = ./.;
    name = "calculator";
  };

  # The vendorHash is a SHA-256 hash of the vendored dependencies, used for reproducible builds.
  #
  # How to update:
  #   1. Run 'go mod vendor'
  #   2. Run 'nix hash path --sri vendor/'
  vendorHash = "sha256-9jK3jKbFp+5WSQfMbNzwIB55bC5KScZOaFHItffTF00=";

  # Inject version at build time via ldflags
  ldflags = [
    "-s"
    "-w"
    "-X=github.com/stefankuehnel/calculator/cmd.version=${version}"
  ];

  meta = with lib; {
    description = "A Command-Line Interface (CLI) Calculator Written in Go";
    homepage = "https://github.com/stefankuehnel/calculator";
    license = licenses.gpl3;
    mainProgram = "calculator";
    maintainers = [
      {
        name = "Stefan Kühnel";
        email = "git@stefankuehnel.com";
      }
    ];
  };
}
