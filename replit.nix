{ pkgs }: {
  deps = [
    pkgs.iputils
    pkgs.unixtools.ping
    pkgs.tree
  ];
}