# Development Environment Setup

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.8+ (to run acceptance tests)
- [Go](https://golang.org/doc/install) >=1.21 (to build the provider plugin)

### Building the Provider

To compile the provider, run `make build`.

```console
make build
```

This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```console
ls -la ./$GOPATH/bin/terraform-provider-assert
```

### Testing the Provider

In order to test the provider, you can run `make test`.

```console
make test
```

### Using the Provider

With Terraform 1.8 and later, [development overrides for provider developers](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory when using Windows) with at least the following options:

```terraform
provider_installation {
  dev_overrides {
    "bschaatsbergen/assert" = "[REPLACE WITH GOPATH]/bin"
  }
  direct {}
}
```
