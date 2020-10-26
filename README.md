# go-tools

## protobuf-any-encoder

this CLI tool will take an input data or file which has a json payload and wrapper that into a `Protobuf Struct` into a `Protobuf Any`. It will output the `json` representation that you can use, in tools such as [BloomRPC](https://github.com/uw-labs/bloomrpc)

### Installation

```sh
go install github.com/edfoh/go-tools/protobuf-any-encoder
```

### Usage Examples

```sh
data=$(cat <<EOF
{
    "id": "123",
    "inner": {
        "name": "test"
    }
}
EOF
)

protobuf-any-encoder -data-input $data
```


Reading from a file

```sh
protobuf-any-encoder -data-file ./data.json
```

Here's an example of the output

```json
{
    "type_url": "type.googleapis.com/google.protobuf.Struct",
    "value": "CgsKAmlkEgUaAzEyMwobCgVpbm5lchISKhAKDgoEbmFtZRIGGgR0ZXN0"
}
```
