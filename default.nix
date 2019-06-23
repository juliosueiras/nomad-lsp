with import <nixpkgs> {};

buildGoModule rec {
  name = "nomad-lsp";
  version = "0.0.1";
  src = ./.;

  modSha256 = "031qds0wqv5sx2j04gdr3wkp8dcvgkk3pxczy5vqq90n92vk57v2"; 

  goPackagePath = "github.com/juliosueiras/nomad-lsp";
  subPackages = [ "." ];

  meta = with stdenv.lib; {
    description = "Language Server Protocol for Nomad";
    homepage = https://github.com/juliosueiras/nomad-lsp;
    license = licenses.mit;
    maintainers = with maintainers; [ juliosueiras ];
    platforms = platforms.all;
  };
}
