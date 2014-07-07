How to test:

Build...

```bash
$ go build .
```

Run...

```bash
$ ./worker
```

Create several jobs!
```bash
$ for i in {1..30}; do curl localhost:8000/work -d name=$USER -d delay=$(expr $i % 11)s; done
```
