{
  lib,
  dockerTools,

  localPackages,
}:

dockerTools.buildLayeredImage {
  name = localPackages.calculator.pname;
  tag = localPackages.calculator.version;

  contents = [
    localPackages.calculator

    dockerTools.caCertificates
  ];

  config = {
    Entrypoint = [ (lib.getExe localPackages.calculator) ];

    Labels = {
      "org.opencontainers.image.title" = localPackages.calculator.pname;
      "org.opencontainers.image.description" = localPackages.calculator.meta.description;
      "org.opencontainers.image.url" = localPackages.calculator.meta.homepage;
      "org.opencontainers.image.source" = localPackages.calculator.meta.homepage;
      "org.opencontainers.image.version" = localPackages.calculator.version;
      "org.opencontainers.image.licenses" = localPackages.calculator.meta.license.spdxId;
    };
  };
}
