# Short ID

[![Build Status](https://travis-ci.org/neptulon/shortid.svg?branch=master)](https://travis-ci.org/neptulon/shortid)
[![GoDoc](https://godoc.org/github.com/neptulon/shortid?status.svg)](https://godoc.org/github.com/neptulon/shortid)

URL-Safe Base64 encoded (unpadded) unique ID generator using "crypto/rand".

You can generate:
* Crypto-safe IDs of custom length.
* UUID v4 (128 bit) in short ID form.

## Example

```go
import (
	"github.com/neptulon/shortid"
)

func main() {
	// short ID with a custom (64 bits) length
	id, err := shortid.ID(64)
	// id: XOga_Cq2h-8

	// standard UUID (128 bits) in short ID form
	uuid, err := shortid.UUID()
	// uuid: gK_qBTwDH5i31evhbD-Wvw
}
```

## License

[MIT](LICENSE)
