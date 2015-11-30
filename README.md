# Short ID

[![Build Status](https://travis-ci.org/neptulon/shortid.svg?branch=master)](https://travis-ci.org/neptulon/shortid)
[![GoDoc](https://godoc.org/github.com/neptulon/shortid?status.svg)](https://godoc.org/github.com/neptulon/shortid)

URL safe base64 encoded unique ID generator using "crypto/rand". Generated IDs are cryptographically secure pseudorandom numbers that are presented as an unpadded base64 encoded string.

## Example

```go
import (
	"github.com/neptulon/shortid"
)

func main() {
  id, err := ID(64)
  // id: XOga_Cq2h-8

  uuid, err := UUID()
  // uuid: gK_qBTwDH5i31evhbD-Wvw
}
```

## License

[MIT](LICENSE)
