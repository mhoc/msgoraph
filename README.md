# msgoraph

A zero dependency Go Client for [Microsoft's Graph API](https://developer.microsoft.com/en-us/graph/docs/concepts/overview). This is built and distributed under all of the philosophies of [vgo](https://research.swtch.com/vgo) for future compatibility, but should work with a simple `go get`, `dep`, or your package management tool of choice until `vgo` is stable. 

## Stability Warning

This library is pre-release, under active development, and has no tests. We will do our best to ensure that tagged releases are stable enough to use the functionality they export, but bugs could happen. Becuase this is pre-release, the Go Import Compatibility Rule does not apply and backward-incompatible changes should be expected between minor pre-release versions. Make sure to pin your version. 

## Getting Started

```go
package main

import (
  "fmt"
  "github.com/mhoc/msgoraph"
  "github.com/mhoc/msgoraph/user"
)

func main() {
  client := msgoraph.NewClient(clientID, clientSecret)
  u, _ := client.Tenant(tenantID).User(emailAddress, user.DefaultFields)
  fmt.Printf("%v\n", u.PreferredName)
}
```

## Supported Operations

- `List Users`
- `Get User`
