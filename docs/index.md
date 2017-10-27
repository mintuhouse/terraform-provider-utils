---
layout: default
---

# Terraform Provider Utils
Allow custom functions in terraform

# Requirements
-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

# Usage

```
# For example, restrict utils version in 0.1.x
provider "utils" {
  version = "~> 0.1"
}
```

### See individual data source / resource docs for their usage

[d/root_domain](d/root_domain) - Get root domain name from full domain