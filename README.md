# mangrove
gramedia product fetcher

## quick start
fetching product total `default=1000`
```sh
$ go build -o mangrove .
$ ./mangrove
```
## using flags
fetching products total (`-ti`) 5000
```sh
$ ./mangrove -ti 5000
```

this could trigger memory leak panic, tune `-ml` flag to lower than `1000` for better result (but slower)

> panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation.





	
