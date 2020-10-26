# go-tools

## protobuf-any-encoder

this CLI tool will take an input data with a specified type and wraps that within [Protobuf Any](https://developers.google.com/protocol-buffers/docs/proto3#any). It will output the `json` representation with the **encoded** value that you can use, in tools such as [BloomRPC](https://github.com/uw-labs/bloomrpc)

Currently supports `StringValue, Int32Value and Struct`

### Installation

```sh
go install github.com/edfoh/go-tools/protobuf-any-encoder
```

### Usage Examples

```sh
# json to struct 
data=$(cat <<EOF
{
    "id": "123",
    "inner": {
        "name": "test"
    }
}
EOF
)

protobuf-any-encoder -data $data -type struct

# string value
data="test"
protobuf-any-encoder -data $data -type string

# int32 value
data=456
protobuf-any-encoder -data $data -type int

```

Here's an example of the output

```json
{
    "type_url": "type.googleapis.com/google.protobuf.Struct",
    "value": "CgsKAmlkEgUaAzEyMwobCgVpbm5lchISKhAKDgoEbmFtZRIGGgR0ZXN0"
}
```
