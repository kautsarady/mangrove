# mangrove
gramedia product fetcher

# branch
- [full-concurrent](https://github.com/kautsarady/mangrove/tree/full-concurrent)

# quick start
```sh
$ go build -o mangrove .
$ ./mangrove
```
# using flags
```sh
$ ./mangrove -c "ebook" -ti 10000
```
Available option (29 Oct 2018).

| Category (-c)	|Total Item (-ti) |
| ------------- |:----------------:|
|     "buku"    |       56074	    |
|     "ebook"    |       33596	    |
|     "non-fiksi"    |     32189	    |
|     "buku-anak"    |      15680	    |
|     "remaja"    |      8204	    |
|     "adult-novels-ebook"    |      7754	    |
|     "fiksi-sastra"    |      6908	    |
|     "bayi-balita"    |      5634	    |
|     "agama"    |       4909	    |
<br>

this could trigger memory leak panic, tune `-ml` flag to lower than `500` for better result (but slower)

> panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation.





	
