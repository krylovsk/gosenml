# GoSenML

Go (golang) library for [SenML](http://www.ietf.org/archive/id/draft-jennings-senml-10.txt)

# TODO
- [ ] XML encoder
- [ ] Binary (exi/cbor) encoder
- [ ] Message 'compaction'
- [ ] Proper message validation

# Example

```go
package main

import (
	"fmt"
	"github.com/krylovsk/gosenml"
)

func main() {
	v := gosenml.Entry{
		Name:  "sensor1",
		Value: 42,
	}

	m1 := gosenml.NewMessage(v)
	m1.BaseName = "http://example.com/"
	m1.BaseUnits = "degC"

	encoder := gosenml.NewJsonEncoder()
	b, _ := encoder.EncodeMessage(m1)
	fmt.Println(string(b))

	m2, _ := encoder.DecodeMessage(b)

	m3 := m2.Expand()
	b, _ = encoder.EncodeMessage(&m3)
	fmt.Println(string(b))
}
```
