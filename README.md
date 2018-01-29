go-cryptorz - encoder / decoder
---

Encode and Decode some data with cypher, base64 and AES.

Usage:
```golang
	key := []byte("thiagozs-poc-of-concept;1234567@")

	cz := NewCryptorz(key)

	encoder := cz.ZEncrypt("Teste de encode/decoder")

	fmt.Printf("Encoder String: %s\n", encoder)

	decoder := cz.ZDecrypt(encoder)

	fmt.Printf("Decoder String: %s\n", decoder)

```

For more details, you can see the source code.

### Author
- @thiagozs