{
  lib,
  buildGoModule,

  antlr4,
  go-task,
}:

buildGoModule rec {
  pname = "calculator";
  version = "0.1.0";

  # See: https://nix.dev/guides/best-practices#reproducible-source-paths
  src = builtins.path {
    path = ./.;
    name = "calculator";
  };

  # The vendorHash is a SHA-256 hash of the dependencies. The hash makes sure
  # that each build gets the same dependencies.
  #
  # To get a new hash, do these steps:
  #   1. Run 'go mod vendor'.
  #   2. Run 'nix hash path --sri vendor/'.
  vendorHash = "sha256-4gn+GyQC79o70qLLV13CtIHXjrQCcPpXBvn6Jv48W8k=";

  # Make a binary that has no dynamic links. No dependency needs cgo. Thus
  # this flag keeps glibc out of the closure and out of the container image
  # that calculator-docker.nix makes. Taskfile.yaml sets CGO_ENABLED to the
  # same value.
  env.CGO_ENABLED = 0;

  nativeBuildInputs = [
    antlr4
    go-task
  ];

  # ANTLR makes the parser in internal/parser/ from the grammar in
  # grammar/Calculator.g4. The build makes it again, so that the binary
  # always matches the grammar and never an old copy in git.
  preBuild = ''
    task --force codegen:antlr4:go
  '';

  # The linker puts the version into the binary at build time.
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
