variable "tags" {
  type = map(string)
  validation {
    condition     = provider::assert::value(var.tags, "value1")
    error_message = "The tags map must contain the value 'value1'"
  }
}
