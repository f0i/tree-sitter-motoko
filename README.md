# Tree-sitter parser for Motoko

Tree sitter parser for the [Motoko programming language](https://github.com/dfinity/motoko).

Motoko is the language of the [Internet Computer](https://internetcomputer.org/), a web3 platform that serves HTTP requests through distributed, containerized applications called canisters.

## Setup

In addition to the normal nvim.treesitter config, add the following:

```lua
local parser_config = require('nvim-treesitter.parsers').get_parser_configs()
parser_config.motoko = {
  install_info = {
    url = 'https://github.com/f0i/tree-sitter-motoko', -- local path or git repo
    files = { 'src/parser.c', 'src/scanner.c' }, -- note that some parsers also require src/scanner.c or src/scanner.cc
    -- optional entries:
    branch = 'main', -- default branch in case of git repo if different from master
    generate_requires_npm = false, -- if stand-alone parser without npm dependencies
    requires_generate_from_grammar = false, -- if folder contains pre-generated src/parser.c
  },
}

vim.filetype.add { extension = { mo = 'motoko' } }
```

And copy the queries to the rtp.

```bash
mkdir -p ~/.config/nvim/queries/motoko/
cp queries/motoko/* ~/.config/nvim/queries/motoko/
```

## References

The grammar is based on [polychromatist/tree-sitter-motoko](https://github.com/polychromatist/tree-sitter-motoko).

Current Motoko Grammer: [Motoko Grammar](https://github.com/dfinity/motoko/blob/b8f76986b4f6b1556569c72decab3f7381bc6cdc/doc/md/examples/grammar.txt).
A copy of this file is located in specs/grammar.txt

