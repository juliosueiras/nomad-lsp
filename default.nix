with import <nixpkgs> {};

buildGoModule rec {
  name = "nomad-lsp";
  version = "0.0.1";
  src = ./.;

  modSha256 = "0ck84vwznlc1lsfpzcr7qyrk1nx524syc6cci8xkcs5ddqf5s22g"; 

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
