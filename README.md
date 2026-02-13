# esmi-go

`esmi-go` is a Go binding for AMD E-SMI (EPYC™ System Management Interface), which wraps `/dev/hsmp` character device for AMD EPYC processors.

## Requirements

- Linux
- Go 1.21+
- CGO enabled
- C toolchain (`gcc` or `clang`)
- [AMD HSMP kernel module](https://github.com/amd/amd_hsmp)

## Quick Start

```go
cli, err := esmi.NewClient()
if err != nil {
    log.Fatal(err)
}
defer cli.Close()

sockets, err := cli.NumberOfSockets()
if err != nil {
    log.Fatal(err)
}

proto, err := cli.HSMPProtoVersion()
if err != nil {
    log.Fatal(err)
}

log.Printf("sockets=%d hsmp_proto=%d", sockets, proto)
```

