# autosdk

## auto generate sdk-go

```shell
make gensdk.debug.fmt
```

## sdk directory introduction

```shell

```autosdk
├── clientset.go  # generated clientset
├── fake
│   └── fake_clientset.go  # generated fake clientset
├── go.mod
├── go.sum
├── pb # generated pb.go by proto files
│   └── corev1
│       ├── core.pb.go
│       ├── credential.pb.go
│       └── machine.pb.go
├── rest # sdk request frame
│   ├── client.go
│   ├── option.go
│   └── request.go
└── typed # generated sdk file by your proto's GRPC-gateway method
    ├── corev1
    │   ├── corev1_client.go
    │   ├── credential.go
    │   ├── credential_expansion.go
    │   ├── fake # fake GRPC-gateway method
    │   │   ├── fake_corev1_client.go
    │   │   ├── fake_credential.go
    │   │   └── fake_machine.go
    │   ├── machine.go
    │   └── machine_expansion.go
    ├── direct_client.go
    └── fake
        └── fake_direct_client.go
