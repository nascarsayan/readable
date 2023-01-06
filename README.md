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
// Dari Findlay Joaquin the pathos of Bandytown 
// preserves Daisy Cletis Alarice and 29 salty hogs

// + Interconvert between UUIDv4 and fluuid.

uuid := guid.New()
fmt.Println(uuid)
// 92e39d12-d0a6-4953-8999-edbf85f7ad66
sentence = fluuid.FromUUID(uuid)
fmt.Println(sentence)
// Janot Boniface Harriet the contestant of Broseley 
// preserved Carrissa Gayler Hahnert and 31 windy gnus

hex, err := fluuid.ToUUID(sentence)
if err != nil {
  fmt.Errorf("error converting fluuid to uuid: %v", err)
}
fmt.Println(hex)
// 92e39d12-d0a6-4953-8999-edbf85f7ad66
```

#### Create smol IDs from UUIDs or fluuids

```go
package main

import (
  "fmt"

  "github.com/nascarsayan/fluuid"
)

long := fluuid.New()
fmt.Println(long)
// Allix Ernestus Matthieu the millwright of Cementon 
// dislikes Nerissa Fairfax Drucilla and 13 old aardvarks
short, err := fluuid.Smol(long)
if err != nil {
  fmt.Errorf("error converting uuid to smol: %v", err)
}
fmt.Println(short)
// 64 light doves lively ran
```
