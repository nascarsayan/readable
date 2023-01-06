### Fluuid

Fluuid is a simple go package for generating fluent UUIDs.

### Features

- Built on UUID v4
- Optionally pass your UUID to generate a unique sentence
- 128 Bit Crypto Secure
- Grammatically correct sentences
- Easy to remember
- Has a Shakespeare feeling
- Universally Unique Identifier
- Generate Low Entropy 32 Bit Tokens

### Installation

```bash
go get github.com/nascarsayan/fluuid
```

### Usage

#### Create readable UUIDs

```go
package main

import (
  "fmt"

  guid "github.com/google/uuid"
  "github.com/nascarsayan/fluuid"
)

sentence := fluuid.New()
fmt.Println(uid) 
// Marylee Cyril Pammi the plagioclase of Carter pushs Nathalie Curcio Glorianna and 20 noisy alligators

uuid := guid.New()
fmt.Println(uuid)
// a8b9d106-9ad5-4111-9693-bbe7c5a5ee86
sentence = fluuid.FromUUID(uuid)
fmt.Println(sentence)
// Kaye Fabio Groveman the lift of Allegre punched Blondie Dieter Lorilyn and 17 quiet snakes

hex, err := fluuid.ToUUID(sentence)
if err != nil {
  fmt.Errorf("error converting fluuid to uuid: %v", err)
}
fmt.Println(hex)
// a8b9d106-9ad5-4111-9693-bbe7c5a5ee86
```

#### Create smol IDs from UUIDs or fluuids

```go
package main

import (
  "fmt"

  "github.com/nascarsayan/fluuid"
)

sentence, err := fluuid.FromUUID(guid.New())
if err != nil {
  fmt.Errorf("error converting uuid to smol: %v", err)
}
fmt.Println(sentence)
// 43 weak rats judgementally felt
```
