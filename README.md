# Evidence Tool


## Crawlhash
This tool crawls a directory tree and outputs hashes to a json log file. The hash,
absolute file path, and datetime of the recording are save as a list of json objects.

### Instalation

Opt 1. **Git Clone**


Opt 2. **Pre-built bianaries**


To make evitools easier to use, install the bianaries to folder in you your path

_make the directory_
```bash
mkdir -p .local/bin/evitools/
```

_build and move_
```bash
go build -o ~/local/bin/evitools/crawlhash [path]/main.go
```


### Usage


#### crawlhash

Crawlhash takes in one filepath argument. This is the path to the root directory scan.
A singular file, `log.json` will be written to the directory the directory that crawlhash was 
run from. This file contains the output from crawlhash.

```bash
crawlhash ~/path/to/dir
```
