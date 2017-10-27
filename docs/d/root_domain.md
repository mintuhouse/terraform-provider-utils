---
layout: default
---

# utils\_root\_domain

Provides a utils data source to get the root domain from a full domain name

## Example Usage

```hcl
data "utils_root_domain" "default" {
  domain          = "test.example.com"
}

# Available as data.utils_root_domain.default.name
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The full domain

## Attributes Reference

The following attributes are exported:

* `name` - The root domain name of the full domain passed.
