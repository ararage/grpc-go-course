protoc -Igreet/proto --go_out=. --go_opt=module=github.com/ararage/grpc-go-course --go-grpc_out=. --go-grpc_opt=module=github.com/ararage/grpc-go-course greet/proto/dummy.proto

## Greet Project

### Build project

```console
$ make greet
```

### Run greet client

```console
$ ./bin/greet/client
```

### Run greet server

```console
$ ./bin/greet/server
```


## Calculator Project

### Build project

```console
$ make calculator
```

### Run calculator client

```console
$ ./bin/calculator/client
```

### Run calculator server

```console
$ ./bin/calculator/server
```