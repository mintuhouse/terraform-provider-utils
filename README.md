# Terraform Provider Utils
Allow custom functions in terraform

[![Build Status](https://travis-ci.org/mintuhouse/terraform-provider-utils.svg?branch=master)](https://travis-ci.org/mintuhouse/terraform-provider-utils)

# Requirements
-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

# Usage

```sh
$ go get github.com/mintuhouse/terraform-provider-utils
$ ln -s $GOPATH/bin/terraform-provider-utils ~/.terraform.d/plugins
```

```
# For example, restrict utils version in 0.1.x
provider "utils" {
  version = "~> 0.1"
}
```

# Why? The Story
* Allowing custom functions inside terraform interpolations is a [long standing request](https://github.com/hashicorp/terraform/issues/2771). Once it is implemented, this repository will then be converted to such a plugin.
* Until then, as per [this comment](https://github.com/hashicorp/terraform/issues/15603#issuecomment-316784100), only (and preferred) way to achieve that today is to use a custom provider data source.

# Development
## Guidelines
* Don't define any of the functions inside this repository. Instead import an appropriate package.
* In some cases you may want to create an [external data source](https://www.terraform.io/docs/providers/external/data_source.html). Any scripts part of such data sources should be part of `external/` folder

## Testing
```sh
$ cd $GOPATH/github.com/mintuhouse/terraform-provider-utils
$ make test
```

# Author
[Hasan Kumar](https://github.com/mintuhouse)