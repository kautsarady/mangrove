# mangrove
gramedia product fetcher

## branch
- [full-concurrent](https://github.com/kautsarady/mangrove/tree/full-concurrent)

## quick start
fetching `1000` products to `data.json`
```sh
$ go build -o mangrove .
$ ./mangrove
```
this could trigger memory leak panic, tune `-ml` flag to lower than `1000` for better result (but slower)

> panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation.
## using flags
- `-ti` total products to fetch (`default=1000`)
- `-o` output name file, MUST be a json file (`default=data.json`)
- `-ml` max load per fetch, depend on host machine, higher is faster (`default=1000`)
