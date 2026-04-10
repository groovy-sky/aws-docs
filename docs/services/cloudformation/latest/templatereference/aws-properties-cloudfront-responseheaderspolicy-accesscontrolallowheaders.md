This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy AccessControlAllowHeaders

A list of HTTP header names that CloudFront includes as values for the
`Access-Control-Allow-Headers` HTTP response header.

For more information about the `Access-Control-Allow-Headers` HTTP response
header, see [Access-Control-Allow-Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.

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

The list of HTTP header names. You can specify `*` to allow all
headers.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::ResponseHeadersPolicy

AccessControlAllowMethods

All content copied from https://docs.aws.amazon.com/.
