This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy AccessControlAllowOrigins

A list of origins (domain names) that CloudFront can use as the value for the
`Access-Control-Allow-Origin` HTTP response header.

For more information about the `Access-Control-Allow-Origin` HTTP response
header, see [Access-Control-Allow-Origin](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ String, ... ]
}

```

### YAML

```yaml

  Items:
    - String

```

## Properties

`Items`

The list of origins (domain names). You can specify `*` to allow all
origins.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessControlAllowMethods

AccessControlExposeHeaders

All content copied from https://docs.aws.amazon.com/.
