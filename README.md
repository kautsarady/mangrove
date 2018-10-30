# mangrove
gramedia product fetcher

## quick start
fetching `1000` products to `data.json`
```sh
$ go build -o mangrove .
$ ./mangrove
```
## using flags
- `-ti` total products to fetch (`default=1000`)
- `-o` output name file, MUST be a json file (`default=data.json`)
