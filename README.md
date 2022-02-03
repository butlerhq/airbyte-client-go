# Go API client for airbyte

Airbyte Configuration API
[https://airbyte.io](https://airbyte.io).

Generated using https://github.com/deepmap/oapi-codegen

## Installation

Install the following dependencies:

```shell
go get github.com/butlerhq/airbyte-api-go
```

Put the package under your project folder and add the following in import:

```golang
import "github.com/butlerhq/airbyte-api-go/airbyte"
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
* `PtrEmail`

## Author
matthieu@heybutler.io

