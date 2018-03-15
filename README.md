# msgoraph

[![Documentation](https://godoc.org/github.com/mhoc/msgoraph?status.svg)](http://godoc.org/github.com/mhoc/msgoraph)

A zero dependency Go Client for [Microsoft's Graph API](https://developer.microsoft.com/en-us/graph/docs/concepts/overview). This is built and distributed under all of the philosophies of [vgo](https://research.swtch.com/vgo) for future compatibility, but should work with a simple `go get`, `dep`, or your package management tool of choice until `vgo` is stable. 

## Stability Warning

This library is pre-release, under active development, and has no tests. We will do our best to ensure that tagged releases are stable enough to use the functionality they export, but bugs could happen. Becuase this is pre-release, the Go Import Compatibility Rule does not apply and backward-incompatible changes should be expected between minor pre-release versions. Make sure to pin your version.

## Getting Started

```go
package main

import (
  "fmt"
  "github.com/mhoc/msgoraph"
)

func main() {
  client := msgoraph.NewClient(clientID, clientSecret)
  u, _ := client.Tenant(tenantID).UserWithFields(emailAddress, msgoraph.UserDefaultFields)
  fmt.Printf("%v\n", u.PreferredName)
}
```

## API Design

msgoraph's API is designed to be easy to use, yet scalable to meet accomodate the entire breadth of the Graph API (eventually).

In general, the API is designed around resources like a `User`. 
Most of these resources have at least one of Create, Read, List, Update, and Delete operations.
Additionally, many resources have "sub-resources" which each then have their own set of operations; like a User's mail messages.

All operations on a resource are methods on the resource type. For examples, users:

```go
msgoraph.User{ ID: "12345" }.Get()
```

While you can manually create resource types using, say, `User{}`, its best to stick with the pre-defined helpers
the package exports, and never manually instantiate the types the package exports.

```go
msgoraph.UserByID("12345").Get()
```

`List`-like operations are a little different, because they act on groups of resources, not individual resources. In these
cases, instead of the operation existing on the resource type it exists on a generally uninteresting intermediate type which
is generally the plural form of the individual resource:

```go
msgoraph.Users().List()
```

## Supported Operations

- `List Users`
- `Create User`
- `Get User`
- `Delete User`
