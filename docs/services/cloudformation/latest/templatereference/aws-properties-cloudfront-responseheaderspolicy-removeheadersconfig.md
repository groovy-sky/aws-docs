This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy RemoveHeadersConfig

A list of HTTP header names that CloudFront removes from HTTP responses to requests that match the
cache behavior that this response headers policy is attached to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ RemoveHeader, ... ]
}

```

### YAML

```yaml

  Items:
    - RemoveHeader

```

## Properties

`Items`

The list of HTTP header names.

_Required_: Yes

_Type_: Array of [RemoveHeader](aws-properties-cloudfront-responseheaderspolicy-removeheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveHeader

ResponseHeadersPolicyConfig

All content copied from https://docs.aws.amazon.com/.
