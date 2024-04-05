---
page_title: "ip function - terraform-provider-assert"
subcategory: "IP Address Functions"
description: |-
  Checks whether a string is a valid IP address (IPv4 or IPv6)
---

# function: ip



## Terraform Test Example

```terraform
run "check_valid_ip_google_compute_address" {

  command = plan

  assert {
    condition     = provider::assert::ip(google_compute_address.example.address)
    error_message = "Address is not a valid IP address"
  }
}
```

## Variable Validation Example

```terraform
variable "ip_address" {
  type = string
  validation {
    condition     = provider::assert::ip(var.ip_address)
    error_message = "Invalid IP address"
  }
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
ip(ip_address string) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `ip_address` (String) The string to check


## Return Type

The return type of `ip` is a boolean.