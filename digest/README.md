Implement a command to calculate the checksum of data read from stdin.

Some of the popular checksum algorithms are: md5, sha256, sha512.

```
# Linux
$ echo a | md5sum
60b725f10c9c85c70d97880dfe8191b3  -

$ echo a | sha256sum
87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7  -

$ echo a | sha512sum
162b0b32f02482d5aca0a7c93dd03ceac3acd7e410a5f18f3fb990fc958ae0df6f32233b91831eaf99ca581a8c4ddf9c8ba315ac482db6d4ea01cc7884a635be  -
```

```
# Linux
$ echo a | md5sum
60b725f10c9c85c70d97880dfe8191b3  -

$ echo a | shasum -a 256
87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7  -

$ echo a | shasum -a 512
162b0b32f02482d5aca0a7c93dd03ceac3acd7e410a5f18f3fb990fc958ae0df6f32233b91831eaf99ca581a8c4ddf9c8ba315ac482db6d4ea01cc7884a635be  -
```

## Introduction

* [sha256.New](https://golang.org/pkg/crypto/sha256/#New)
* [sha512.New](https://golang.org/pkg/crypto/sha512/#New)
* [sha256.New](https://golang.org/pkg/crypto/md5/#New)

All the above functions return the same interface:

* [hash.Hash](https://golang.org/pkg/hash/#Hash)
  * hash.Hash is an io.Writer

To calculate the checksum:

```
// `hash` is a variable, of type hash.Hash
// write data to `hash`...

var digest []byte
digest = hash.Sum(digest)
```

Once you have the digest in bytes, encode it into hexadecimal.

* [hex.EncodeToString](https://golang.org/pkg/encoding/hex/#EncodeToString)

## Examples

```
$ echo a | go run digest.go md5
60b725f10c9c85c70d97880dfe8191b3

$ echo a | go run digest.go sha256
87428fc522803d31065e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7

$ echo a | go run digest.go sha512
162b0b32f02482d5aca0a7c93dd03ceac3acd7e410a5f18f3fb990fc958ae0df6f32233b91831eaf99ca581a8c4ddf9c8ba315ac482db6d4ea01cc7884a635be
```

If input data is different, output should be different:

```
$ echo ab | go run digest.go md5
daa8075d6ac5ff8d0c6d4650adb4ef29
```

You can use this command to calculate the digest of a HTTP response:

```
$ curl http://google.com | go run digest.go md5
```
