This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy CustomHeadersConfig

A list of HTTP response header names and their values. CloudFront includes these headers in
HTTP responses that it sends for requests that match a cache behavior that's associated
with this response headers policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ CustomHeader, ... ]
}

```

### YAML

```yaml

  Items:
    - CustomHeader

```

## Properties

`Items`

The list of HTTP response headers and their values.

_Required_: Yes

_Type_: Array of [CustomHeader](aws-properties-cloudfront-responseheaderspolicy-customheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomHeader

FrameOptions

All content copied from https://docs.aws.amazon.com/.
