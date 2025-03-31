# Packer

A `golang` app that will zip or unzip a folder to a location, allowing excluding path matches.

## Args

| Name        | Arg                | Description           |
|-------------|--------------------|-----------------------|
| Config Name | -c "<config_name>" | Path to config (default: `config.toml`)        |
| Pack Name   | -n "<pack_name>"   | Name of pack in config        |
| Unpack      | -u                 | Bool indicating unpack |

## Example `config.toml`

```toml
# %s will be replaced with the arg pack_name
ZipFmt = "~/my-zip-dir/%s.zip" 

[Packs.example]
Source = "~/my-dir/"
Excludes = []
```

### Pack

The `Source` directory will be zipped exluding filepaths containing an exlude value.

### Unpack

The zip file will be extracted and replace the `Source` directory.

# Personal TODO

- [x] Add args
- [x] nvim-data
- [x] nvim-config
- [x] publish.ps1
- [x] Pack all
- [x] Zip packer + toml
- [ ] Unpack
- [ ] Unpack backup
- [x] Remove ZipFile from config
- [x] Readme
- [ ] Add file count
 
