# go-uuid
[![License][license-img]][license] [![Actions Status][action-img]][action] [![PkgGoDev][pkggodev-img]][pkggodev] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]
[![latest][latest-img]][latest-link]


Package uuid is a Go library that provides a very simple interface to generate and marshal UUIDs. UUIDs are universally unique identifiers that are based on RFC 4122 and are 128 bit (16 byte) values that can be generated as Version 4 random or pseudo-random id. This package supports both binary and text encoding of UUIDs, as well as parsing them. 
The main features of this package are:

- Easy to use: just import the package and call uuid.New() to get a new random UUID.
- Fast and secure: the package uses crypto/rand to generate random bytes for the UUIDs, ensuring high quality randomness and avoiding collisions.
- Flexible and interoperable: the package supports canonical string formats. 
- Compatible and consistent: the package follows the RFC 4122 specification and produces UUIDs that are compatible with other implementations and languages. The package also guarantees that the same UUID will always be encoded and decoded in the same way, regardless of the format.

To use this package, simply import it in your Go code:

import "go.melnyk.org/uuid"

Then you can create a new random UUID by calling uuid.New():

```
id := uuid.New()
```

To parse a UUID from a string or byte slice, use uuid.Parse():

```
id1, err := uuid.Parse("9ba7c648-9a77-4b49-91f7-4560b65fcf88") // parse from canonical string
id2, err := uuid.Parse("{9ba7c648-9a77-4b49-91f7-4560b65fcf88}") // parse from Microsoft formated string
```

To compare two UUIDs, use UUID.Equal():

```
if id1.Equal(id2) {
    // ...
}
```

For more information and examples, please refer to the package documentation at https://pkg.go.dev/go.melnyk.org/uuid.

## Contributing
We encourage and support an active, healthy community of contributors &mdash;
including you! Details are in the [contribution guide](CONTRIBUTING.md) and
the [code of conduct](CODE_OF_CONDUCT.md).

[![Contributor Covenant][covenant-img]](CODE_OF_CONDUCT.md)

## Alternatives
- https://github.com/google/uuid

[covenant-img]: https://img.shields.io/badge/contributor%20covenant-v1.4%20adopted-ff69b4.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: LICENSE
[action-img]: ../../workflows/Test/badge.svg
[action]: ../../actions
[goreport-img]: https://goreportcard.com/badge/go.melnyk.org/uuid
[goreport]: https://goreportcard.com/report/go.melnyk.org/uuid
[codecov-img]: https://codecov.io/gh/melnyk/go-uuid/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/melnyk/go-uuid
[pkggodev-img]: https://pkg.go.dev/badge/go.melnyk.org/uuid
[pkggodev]: https://pkg.go.dev/go.melnyk.org/uuid
[latest-img]: https://img.shields.io/github/v/release/melnyk/go-uuid
[latest-link]: https://github.com/melnyk/go-uuid/releases/latest
