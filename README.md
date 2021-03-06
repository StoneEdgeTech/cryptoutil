# Basic cryptographic helper functions
This repository contains two functions to encrypt and decrypt text using the AES algorithm.
```go
// encrypt your plaintext string using your plaintext key. An initialization vector (iv) will
// be generated and prepended to your byte-slice. Expect that the first int(aes.BlockSize)
// bytes in the cypherTextAsByteSlice is your iv. The iv must be unique but not secure.
cypherTextAsByteSlice, err := cryptoutil.AESEncrypt([]byte(key), []byte(plaintext))
```

```go
// decrypt the cyphertext byte slice, assuming the first int(aes.BlockSize) bytes are an
// initialization vector (iv). The iv must be unique but not secure
decryptedPlainTextAsString, err := cryptoutil.AESDecrypt([]byte(key), encrypted)
```

```go
// encrypt your plaintext as above, but wrap the []byte as a base64-encoded string
base64CypherText, err := B64AESEncrypt([]byte(key), []byte(plaintext))
```

```go
// decrypt a base64-encoded cyphertext string as above
decryptedPlainTextAsString, err := B64AESDecrypt([]byte(key), base64CypherText)
```
