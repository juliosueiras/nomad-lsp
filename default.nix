with import <nixpkgs> {};

buildGoModule rec {
  name = "nomad-lsp";
  version = "0.0.1";
  src = ./.;

  vendorSha256 = "J22ZWmqBTkMHnyrnEPCfFmme8f+x1JPvNBr6P28mNbc="; 

  subPackages = [ "." ];

  meta = with stdenv.lib; {
    description = "Language Server Protocol for Nomad";
    homepage = https://github.com/juliosueiras/nomad-lsp;
    license = licenses.mit;
    maintainers = with maintainers; [ juliosueiras ];
    platforms = platforms.all;
  };
}
