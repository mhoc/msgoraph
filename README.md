# msgoraph

[![Documentation](https://godoc.org/github.com/mhoc/msgoraph?status.svg)](http://godoc.org/github.com/mhoc/msgoraph)

A zero dependency Go Client for [Microsoft's Graph API](https://developer.microsoft.com/en-us/graph/docs/concepts/overview). This is built and distributed under all of the philosophies of [vgo](https://research.swtch.com/vgo) for future compatibility, but should work with a simple `go get`, `dep`, or your package management tool of choice until `vgo` is stable. 

## Disclaimers

This library is completely unaffiliated with Microsoft.

This library is in pre-release, under active development, and has no tests. We will do our best to ensure that tagged releases are stable enough to use the functionality they export, but bugs could happen. 

Because it is in pre-release, the Go Import Compatibility Rule does not apply. Backward-incompatible changes should be expected between all tagged versions and commits. 

## Example Usage

We'll get more examples on how to use this library online as it matures.

For the time being, check out [msgraph-cli](https://github.com/mhoc/msgraph-cli), which uses msgoraph to power most of its internals. That should give a sense of how using this library works, at least in terms of the version that the cli pins against.

## Supported Features

- Authorization on behalf of a user
- Users :: Create
- Users :: Delete
- Users :: Get
- Users :: List
- Users :: Update
