with import <nixpkgs> {};

buildGoModule rec {
  name = "nomad-lsp";
  version = "0.0.1";
  src = ./.;

  modSha256 = "0gfbl5lw0x3n6lw6i3cjjikyxgfxkjgqf9f6zliglxlnsbvmf3jw"; 

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
