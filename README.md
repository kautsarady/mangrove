# mangrove
gramedia product fetcher

# quick start
```sh
$ go build -o mangrove .
$ ./mangrove
```
# using flags
```sh
$ ./mangrove -c "ebook" -pp 100 -ti 10000
```
Available option (29 Oct 2018).

| Category (-c)	| Per Page (-pp) | Total Item (-ti) |
| ------------- |:--------------:| ----------------:|
|     "buku"    |      100	 |       56074	    |
|     "ebook"    |      100	 |       33596	    |
|     "non-fiksi"    |      100	 |       32189	    |
|     "buku-anak"    |      100	 |       15680	    |
|     "remaja"    |      100	 |       8204	    |
|     "adult-novels-ebook"    |      100	 |       7754	    |
|     "fiksi-sastra"    |      100	 |       6908	    |
|     "bayi-balita"    |      100	 |       5634	    |
|     "agama"    |      100	 |       4909	    |
<br>
this could triggering memory leak panic, tune `-ml` flag to lower than `500` for better result (but slower)

> panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation.





	
