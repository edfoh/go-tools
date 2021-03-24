# go-tools

## HMAC Generator

This tool is to support HMAC generation of webhooks. It uses SHA-256, and generates a signature based on input secret, timestamp and request body.

### Installation 

Git clone this repository first, then run the following command

```sh
go install github.com/edfoh/go-tools/hmac-gen
```


### Usage Example 

```sh
# assuming you have a file called payload.txt that has your request body
payload=$(cat ./payload.txt)

hmac-gen -secret -gq63hsjecrEcDQySBxRTX1pMB2q3vjvr2uX09h9eCc= \
    -ts 2021-03-24T00:45:29Z \
    -payload $payload

# Result: DGM3pfJyjidZXNyJrFiQR5L6BPZgi1JxuJmJ4ycSRi4=

```

## protobuf-any-encoder

this CLI tool will take an input data with a specified type and wraps that within [Protobuf Any](https://developers.google.com/protocol-buffers/docs/proto3#any). It will output the `json` representation with the **encoded** value that you can use, in tools such as [BloomRPC](https://github.com/uw-labs/bloomrpc)

Currently supports `StringValue, Int32Value and Struct`

### Installation

Git clone this repository first, then run the following command

```sh
go install github.com/edfoh/go-tools/protobuf-any-encoder
```

Note: Please make sure your `GOPATH` is set into your PATH.

```
GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"
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
